package gui

import (
	"fmt"
	log "github.com/Sirupsen/logrus"
	"github.com/andlabs/ui"
	"ql/ast/visitor"
	"ql/interfaces"
	"strconv"
)

type GUI struct {
	visitor.BaseVisitor
	GUIForm                   *GUIForm
	Symbols                   interfaces.VarIDValueSymbols
	RegisteredOnShowCallbacks []func()
	SaveDataCallback          func() (interface{}, error)
	Window                    *ui.Window
}

// NewGUI is a method returning a new GUI
func NewGUI() *GUI {
	return new(GUI)
}

func (this *GUI) InitializeGUIForm(form interfaces.Form, symbols interfaces.VarIDValueSymbols) {
	this.setSymbols(symbols)
	this.createAndSetGUIFormFromForm(form)

	form.Accept(this, this.Symbols)
}

func (this *GUI) setSymbols(symbols interfaces.VarIDValueSymbols) {
	if symbols == nil {
		panic("Passing nil symbols to GUI setSymbols")
	}

	this.Symbols = symbols
}

// createAndSetGUIFormFromForm uses Form to create GUIForm
func (this *GUI) createAndSetGUIFormFromForm(form interfaces.Form) {
	if form == nil {
		panic("Passing nil form to GUI createAndSetGUIFormFromForm")
	}

	this.GUIForm = newGUIForm(form)
}

// registerOnShowCallback registers callback functions that are called once the UI is presented to the user
func (this *GUI) registerOnShowCallback(callback func()) {
	this.RegisteredOnShowCallbacks = append(this.RegisteredOnShowCallbacks, callback)
}

// ShowWindow shows the UI after invocation
func (this *GUI) ShowWindow(errorsToDisplay, warningsToDisplay []error) {
	guiError := ui.Main(func() {
		log.Info("Showing GUI")

		contentBox := ui.NewVerticalBox()
		contentBox.SetPadded(true)
		this.Window = ui.NewWindow("QL", 800, 600, false)

		this.Window.OnClosing(func(w *ui.Window) bool {
			log.Info("Destroy of window initiated")
			this.Symbols.SaveToDisk()
			ui.Quit()
			return true
		})

		this.Window.SetChild(contentBox)
		this.Window.SetMargined(true)
		this.Window.Show()

		// if there are any errors/warnings, show a dialog
		if errorsPresent := this.showErrorDialogIfNecessary(errorsToDisplay); !errorsPresent {
			this.showWarningDialogIfNecessary(warningsToDisplay)
			this.ShowForm()
		}
	})

	if guiError != nil {
		log.WithFields(log.Fields{"guiError": guiError}).Panic("Encountered GUI error")
	}
}

// ShowForm visually shows a GUIForm representation of a form
func (this *GUI) ShowForm() {
	this.GUIForm.show(this.Window)

	for _, registeredOnShowCallback := range this.RegisteredOnShowCallbacks {
		registeredOnShowCallback()
	}

	this.addSubmitButton()
}

// VisitForm creates the top level questions in the form's inner body
func (this *GUI) VisitForm(f interfaces.Form, context interface{}) interface{} {
	guiQuestions := handleQuestions(this, f.Questions())

	this.registerOnShowCallback(func() {
		this.GUIForm.addQuestionContainer(this.GUIForm.createQuestionTable(guiQuestions))
	})

	return nil
}

// VisitIf creates questions embedded in an its body and registers show/hide callbacks
func (this *GUI) VisitIf(ifStmt interfaces.If, context interface{}) interface{} {
	guiQuestions := handleQuestions(this, ifStmt.Questions())
	questionsEncompassingContainer := this.GUIForm.createQuestionTable(guiQuestions)

	this.registerOnShowCallback(func() {
		this.GUIForm.addQuestionContainer(questionsEncompassingContainer)
		this.showOrHideContainerDependingOnIfEval(ifStmt, questionsEncompassingContainer)
	})

	this.Symbols.RegisterCallback(func() {
		log.Debug("Received symbols update callback in If")
		this.showOrHideContainerDependingOnIfEval(ifStmt, questionsEncompassingContainer)
	})

	return nil
}

// VisitIfElse creates questions embedded in an its bodies and registers show/hide callbacks
func (this *GUI) VisitIfElse(ifElse interfaces.IfElse, context interface{}) interface{} {
	guiQuestionsIfBody := handleQuestions(this, ifElse.IfBodyQuestions())
	guiQuestionsElseBody := handleQuestions(this, ifElse.ElseBodyQuestions())

	ifQuestionsEncompassingContainer := this.GUIForm.createQuestionTable(guiQuestionsIfBody)
	elseQuestionsEncompassingContainer := this.GUIForm.createQuestionTable(guiQuestionsElseBody)

	this.registerOnShowCallback(func() {
		this.GUIForm.addQuestionContainer(ifQuestionsEncompassingContainer)
		this.GUIForm.addQuestionContainer(elseQuestionsEncompassingContainer)

		this.showOrHideContainerDependingOnIfElseEval(ifElse, ifQuestionsEncompassingContainer, elseQuestionsEncompassingContainer)
	})

	this.Symbols.RegisterCallback(func() {
		log.Debug("Received symbols update callback in IfElse")
		this.showOrHideContainerDependingOnIfElseEval(ifElse, ifQuestionsEncompassingContainer, elseQuestionsEncompassingContainer)
	})

	return nil
}

func (this *GUI) showOrHideContainerDependingOnIfEval(ifStmt interfaces.If, conditionalContainer *ui.Box) {
	conditionValue := ifStmt.EvalConditionAsBool(this.Symbols)

	// if condition evals to true
	if conditionValue {
		conditionalContainer.Show()
		return
	}

	// condition evals to false
	conditionalContainer.Hide()
}

func (this *GUI) showOrHideContainerDependingOnIfElseEval(ifElseStmt interfaces.IfElse, conditionalContainerIfBody, conditionalContainerElseBody *ui.Box) {
	conditionValue := ifElseStmt.EvalConditionAsBool(this.Symbols)

	// ifElse if block condition evals to true
	if conditionValue {
		conditionalContainerIfBody.Show()
		conditionalContainerElseBody.Hide()
		return
	}

	// condition evals to false, show else body
	conditionalContainerIfBody.Hide()
	conditionalContainerElseBody.Show()
}

func handleQuestions(this *GUI, questions []interfaces.Question) []*GUIQuestion {
	var guiQuestions []*GUIQuestion

	for _, question := range questions {
		var guiQuestion *GUIQuestion

		switch question := question.(type) {
		case interfaces.ComputedQuestion:
			guiQuestion = this.handleComputedQuestion(question).GUIQuestion
		case interfaces.InputQuestion:
			guiQuestion = this.handleInputQuestion(question).GUIQuestion
		}

		guiQuestions = append(guiQuestions, guiQuestion)
	}

	return guiQuestions
}

func (this *GUI) handleInputQuestion(question interfaces.InputQuestion) *GUIInputQuestion {
	var guiQuestion *GUIInputQuestion

	questionCallback := func(inputExpr interfaces.Expr, err error) {
		guiQuestion.resetErrorLabelText()
		if err != nil {
			if numError, isNumError := err.(*strconv.NumError); isNumError && numError.Err.Error() == "invalid syntax" {
				guiQuestion.changeErrorLabelText("Not a valid number!")
				log.Error("Presenting invalid number error to user")
			}

			return
		}

		questionIdentifier := question.VarDeclVariableIdentifier()
		log.WithFields(log.Fields{"input": inputExpr, "identifier": questionIdentifier}).Debug("Question input received")
		this.Symbols.SetExprForVarID(inputExpr, questionIdentifier)

		this.updateComputedQuestions()
	}

	guiQuestion = newGUIInputQuestion(question, questionCallback)

	return guiQuestion
}

func (this *GUI) handleComputedQuestion(question interfaces.ComputedQuestion) *GUIComputedQuestion {
	computation := question.Computation()
	guiQuestion := newGUIComputedQuestion(question, computation, question.VarDeclVariableIdentifier())

	this.GUIForm.addComputedQuestion(guiQuestion)

	return guiQuestion
}

// updateComputedQuestions is called after the input of a input question changes and updates computed questions' values
func (this *GUI) updateComputedQuestions() {
	for _, computedQuestion := range this.GUIForm.ComputedQuestions {
		computedQuestionEvalValue := computedQuestion.Expr.Eval(this.Symbols).PrimitiveValue()

		switch computedQuestionEvalValueAsserted := computedQuestionEvalValue.(type) {
		case bool:
			computedQuestion.changeCheckboxValue(computedQuestionEvalValueAsserted)
		case int, string:
			computedQuestion.changeFieldValueText(fmt.Sprintf("%v", computedQuestionEvalValueAsserted))
		}

		// save the computed value to the symbol table
		this.Symbols.SetExprForVarID(computedQuestion.Expr, computedQuestion.VarID)

		log.WithFields(log.Fields{"evaluatesTo": computedQuestionEvalValue}).Info("Computed question value changed")
	}
}

// addSubmitButton adds a submit button to the form
func (this *GUI) addSubmitButton() {
	log.Info("Adding submit button to GUI")

	button := createButton("Submit", func(b *ui.Button) {
		log.Debug("Submit button clicked, save of data initiated")
		this.Symbols.SaveToDisk()
		showMessageBoxForNotification("Data saved to file", "The information entered in the form has been saved to a file.", this.Window)
	})

	this.GUIForm.FormContainer.Append(button, false)
}

// showErrorDialogIfNecessary displays an error dialog only if there are errors
func (this *GUI) showErrorDialogIfNecessary(errorsToDisplay []error) bool {
	if len(errorsToDisplay) == 0 {
		return false
	}

	showMessageBoxForErrors("Errors", errorsToDisplay, this.Window)

	return true
}

// showWarningDialogIfNecessary displays a warning dialog only if there are warnings
func (this *GUI) showWarningDialogIfNecessary(warningsToDisplay []error) bool {
	if len(warningsToDisplay) == 0 {
		return false
	}

	showMessageBoxForErrors("Warnings", warningsToDisplay, this.Window)

	return true
}
