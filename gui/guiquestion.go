package gui

import (
	log "github.com/Sirupsen/logrus"
	"github.com/andlabs/ui"
	"ql/ast/expr"
	"ql/ast/visitor"
	"ql/interfaces"
	"strconv"
)

// GUIQuestion represents a question's elements in term of interface objects
type GUIQuestion struct {
	Label      *ui.Label
	Element    ui.Control
	ErrorLabel *ui.Label
	Question   interfaces.Question
}

// createGUIQuestion creates a GUIQuestion, the last argument indicates if the question should be disabled (no entry allowed)
func createGUIQuestion(question interfaces.Question, callback func(interfaces.Expr, error), disabled bool) *GUIQuestion {
	questionLabel := createLabel(question.LabelAsString())
	guiElementVisitor := newQuestionTypeToGUIElementVisitor(question, callback, disabled)
	questionElement := question.VarDeclValueType().Accept(guiElementVisitor, nil).(ui.Control)
	errorLabel := createLabel("")

	return &GUIQuestion{questionLabel, questionElement, errorLabel, question}
}

// createEnabledGUIQuestion is a convenience method for creating a GUIQuestion that is enabled
func createEnabledGUIQuestion(question interfaces.Question, callback func(interfaces.Expr, error)) *GUIQuestion {
	return createGUIQuestion(question, callback, false)
}

// createDisabledGUIQuestion is a convenience method for creating a GUIQuestion that is disabled
func createDisabledGUIQuestion(question interfaces.Question) *GUIQuestion {
	return createGUIQuestion(question, nil, true)
}

// changeFieldValueText changes the displayed value of an input field
func (this *GUIQuestion) changeFieldValueText(newText string) {
	log.WithFields(log.Fields{"newLabelText": newText}).Debug("Changing text of element")
	this.Element.(*ui.Entry).SetText(newText)
}

// ChangeErrorLabelText changes the error text feedback presented when error occurs
func (this *GUIQuestion) changeErrorLabelText(newText string) {
	this.ErrorLabel.SetText(newText)
}

// changeCheckboxValue changes the displayed value of a checkbox
func (this *GUIQuestion) changeCheckboxValue(newValue bool) {
	log.WithFields(log.Fields{"newValue": newValue}).Debug("Changing checkbox value")
	this.Element.(*ui.Checkbox).SetChecked(newValue)
}

// ResetErrorLabelText removes the error text presented to the user
func (this *GUIQuestion) resetErrorLabelText() {
	this.changeErrorLabelText("")
}

type QuestionTypeToGUIElementVisitor struct {
	callback func(interfaces.Expr, error)
	disabled bool
	interfaces.Question
	*visitor.BaseVisitor
}

// newQuestionTypeToGUIElementVisitor returns a new QuestionTypeToGUIElementVisitor
func newQuestionTypeToGUIElementVisitor(question interfaces.Question, callback func(interfaces.Expr, error), disabled bool) *QuestionTypeToGUIElementVisitor {
	return &QuestionTypeToGUIElementVisitor{callback, disabled, question, visitor.NewBaseVisitor()}
}

// VisitIntegerType returns a GUI element for integers when question type is integer
func (this *QuestionTypeToGUIElementVisitor) VisitIntegerType(i interfaces.IntegerType, context interface{}) interface{} {
	var UIEntity ui.Control
	inputField := createInputTextField("", this.disabled)
	inputField.OnChanged(func(*ui.Entry) {
		inputText := inputField.Text()

		log.WithFields(log.Fields{"value": inputText}).Debug("Input text value changed (integer field)")

		inputTextAsInt, err := strconv.Atoi(inputText)
		if inputText == "" {
			if this.callback != nil {
				this.callback(expr.NewIntegerLiteral(inputTextAsInt), nil)
			}
			return
		} else if err != nil {
			log.Warn("Could not convert input text string to int")

			if this.callback != nil {
				this.callback(nil, err)
			}

			return
		}

		if this.callback != nil {
			this.callback(expr.NewIntegerLiteral(inputTextAsInt), nil)
		}
	})

	UIEntity = inputField

	return UIEntity
}

// VisitBoolType creates a GUI element for booleans when question type is boolean
func (this *QuestionTypeToGUIElementVisitor) VisitBoolType(bo interfaces.BoolType, context interface{}) interface{} {
	var UIEntity ui.Control
	checkbox := createCheckboxConditional(this.disabled)
	checkbox.OnToggled(func(*ui.Checkbox) {
		log.WithFields(log.Fields{"value": checkbox.Checked()}).Debug("Checkbox value changed")

		if this.callback != nil {
			this.callback(expr.NewBoolLiteral(checkbox.Checked()), nil)
		}
	})

	UIEntity = checkbox

	return UIEntity
}

// VisitStringType creates a GUI element for strings when question type is string
func (this *QuestionTypeToGUIElementVisitor) VisitStringType(st interfaces.StringType, context interface{}) interface{} {
	var UIEntity ui.Control
	inputField := createInputTextField("", this.disabled)
	inputField.OnChanged(func(*ui.Entry) {
		inputText := inputField.Text()

		log.WithFields(log.Fields{"value": inputText}).Debug("Input text value changed (string field)")

		if this.callback != nil {
			this.callback(expr.NewStringLiteral(inputText), nil)
		}
	})

	UIEntity = inputField

	return UIEntity
}
