package gui

import (
	log "github.com/Sirupsen/logrus"
	"github.com/andlabs/ui"
	"ql/interfaces"
)

type GUIForm struct {
	Form               interfaces.Form
	QuestionContainers []*ui.Box
	ComputedQuestions  []*GUIComputedQuestion
	FormContainer      *ui.Box
}

// newGUIForm is a constructor method returning a new GUIForm with initialized embedded Form
func newGUIForm(form interfaces.Form) *GUIForm {
	if form == nil {
		panic("Passing nil form to GUI newGUIForm")
	}

	return &GUIForm{Form: form, ComputedQuestions: make([]*GUIComputedQuestion, 0)}
}

// show displays the form box. It should only be called if no semantic errors are present
func (this *GUIForm) show(window *ui.Window) {
	log.WithFields(log.Fields{"identifier": this.Form.Identifier()}).Info("Showing form")

	this.FormContainer = ui.NewVerticalBox()
	window.SetChild(this.FormContainer)
}

// addQuestionContainer appends the question container box to the form container
func (this *GUIForm) addQuestionContainer(questionContainer *ui.Box) {
	this.FormContainer.Append(questionContainer, false)
	log.Info("Adding question container to form")
	this.QuestionContainers = append(this.QuestionContainers, questionContainer)
}

// addComputedQuestion adds a pointer to the GUIComputedQuestion to the GUIForm
func (this *GUIForm) addComputedQuestion(question *GUIComputedQuestion) {
	if question == nil {
		panic("Passing nil question to GUI AddComputedQuestion")
	}

	log.Info("Adding computed question to form")
	this.ComputedQuestions = append(this.ComputedQuestions, question)
}

// createQuestionTable creates a table box containing the passed GUIQuestions
func (this *GUIForm) createQuestionTable(questions []*GUIQuestion) *ui.Box {
	table := ui.NewVerticalBox()

	for _, question := range questions {
		attachQuestionToTable(table, question)
	}

	log.WithFields(log.Fields{"NumOfQuestions": len(questions)}).Info("Created question table")

	return table
}

// attachQuestionToTable is a helper method that attaches a GUIQuestion to the supplied box
func attachQuestionToTable(table *ui.Box, question *GUIQuestion) {
	table.Append(question.Label, false)
	table.Append(question.Element, false)
	table.Append(question.ErrorLabel, false)
}
