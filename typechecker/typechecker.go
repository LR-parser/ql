package typechecker

import (
	log "github.com/Sirupsen/logrus"
	"ql/interfaces"
)

type TypeChecker struct {
	interfaces.TypeChecker
	errorsEncountered    []error
	warningsEncountered  []error
	knownErrors          map[string]bool
	usedLabels           map[interfaces.StringLiteral]interfaces.VarID
	knownIdentifiers     map[interfaces.VarID]bool
	dependenciesForVarID map[interfaces.VarID][]interfaces.VarID
}

func NewTypeChecker() *TypeChecker {
	return &TypeChecker{errorsEncountered: make([]error, 0), warningsEncountered: make([]error, 0), usedLabels: make(map[interfaces.StringLiteral]interfaces.VarID), knownErrors: make(map[string]bool), knownIdentifiers: make(map[interfaces.VarID]bool), dependenciesForVarID: make(map[interfaces.VarID][]interfaces.VarID)}
}

func (this *TypeChecker) dependencyListForVarDecl(varDecl interfaces.VarDecl) []interfaces.VarID {
	varID := varDecl.VariableIdentifier()

	if varID == nil {
		panic("Attempting to get dependencies of nil varID")
	}

	var dependencies []interfaces.VarID

	for _, directDependentVarID := range this.dependenciesForVarID[varID] {
		// first add the direct dependencies as dependencies
		dependencies = append(dependencies, directDependentVarID)

		// then add as dependencies: the dependencies of the directDependency (i.e. the indirect dependencies of the VarID passed)
		dependencies = append(dependencies, this.recursivelyObtainDependenciesForVarID(directDependentVarID, varID)...)
	}

	return dependencies
}

func (this *TypeChecker) recursivelyObtainDependenciesForVarID(varIDToObtainDependenciesFor, varIDNeededToCompleteCycle interfaces.VarID) []interfaces.VarID {
	if varIDToObtainDependenciesFor == nil {
		panic("Attempting to recursively obtain dependencies for nil varIDDependentOn")
	} else if varIDNeededToCompleteCycle == nil {
		panic("Can't recursively obtain dependencies when varIDNeededToCompleteCycle is nil (no stop condition)")
	}

	// if we try to get dependencies of the currently visited VarID we've found a cycle, thus return to prevent infinite loop :)
	if varIDToObtainDependenciesFor == varIDNeededToCompleteCycle {
		return nil
	}

	var dependencies []interfaces.VarID
	for _, dependentVarID := range this.dependenciesForVarID[varIDToObtainDependenciesFor] {
		// if we encounter ourselves, we found a cyclic dependency
		if dependentVarID == varIDToObtainDependenciesFor {
			return nil
		}

		dependencies = append(dependencies, dependentVarID)
		dependencies = append(dependencies, this.recursivelyObtainDependenciesForVarID(dependentVarID, varIDNeededToCompleteCycle)...)
	}

	return dependencies
}

func (this *TypeChecker) DependencyListForVarDeclContainsReferenceToSelf(varDecl interfaces.VarDecl) bool {
	depencyListForQuestionID := this.dependencyListForVarDecl(varDecl)

	for _, dependentVarID := range depencyListForQuestionID {
		if dependentVarID == varDecl.VariableIdentifier() {
			return true
		}
	}

	return false
}

func (this *TypeChecker) AddDependencyForVarDecl(varIDDependentOn interfaces.VarID, varDecl interfaces.VarDecl) {
	// if we are not in a ComputedQuestion context, don't add any dependencies
	if varDecl == nil {
		return
	}

	if varIDDependentOn == nil {
		panic("Attempting to add VarID dependency but passed dependency is nil")
	}

	for _, dependingVarIDAlreadyKnown := range this.dependenciesForVarID[varDecl.VariableIdentifier()] {
		// don't add it again if its already known to be a dependency
		if dependingVarIDAlreadyKnown == varIDDependentOn {
			return
		}
	}

	varID := varDecl.VariableIdentifier()
	this.dependenciesForVarID[varID] = append(this.dependenciesForVarID[varID], varIDDependentOn)

	log.WithFields(log.Fields{"visitedVarID": varID, "varIDDependentOn": varIDDependentOn, "resultingMap": this.dependenciesForVarID}).Debug("Added dependency for currently visited VarID for type checking")
}

func (this *TypeChecker) AddEncounteredError(encounteredError error) {
	if encounteredError == nil {
		panic("Trying to add nil encountered error")
	}

	log.WithFields(log.Fields{"errorEncountered": encounteredError}).Info("Added encountered type checking error")

	if !this.isErrorKnown(encounteredError) {
		this.errorsEncountered = append(this.errorsEncountered, encounteredError)
		this.knownErrors[encounteredError.Error()] = true
	}
}

func (this *TypeChecker) AddEncounteredWarning(encounteredWarning error) {
	if encounteredWarning == nil {
		panic("Trying to add nil encountered warning")
	}

	log.WithFields(log.Fields{"warningEncountered": encounteredWarning}).Info("Added encountered type checking warning")

	if !this.isErrorKnown(encounteredWarning) {
		this.warningsEncountered = append(this.warningsEncountered, encounteredWarning)
		this.knownErrors[encounteredWarning.Error()] = true
	}
}

func (this *TypeChecker) isErrorKnown(errorToCheck error) bool {
	_, errorAlreadyKnown := this.knownErrors[errorToCheck.Error()]
	return errorAlreadyKnown
}

func (this *TypeChecker) EncounteredErrors() []error {
	return this.errorsEncountered
}

func (this *TypeChecker) EncounteredWarnings() []error {
	return this.warningsEncountered
}

func (this *TypeChecker) IsLabelUsed(label interfaces.StringLiteral) bool {
	_, exists := this.usedLabels[label]
	return exists
}

func (this *TypeChecker) VarIDForLabel(label interfaces.StringLiteral) interfaces.VarID {
	return this.usedLabels[label]
}

func (this *TypeChecker) MarkLabelAsUsed(label interfaces.StringLiteral, varDecl interfaces.VarDecl) {
	log.WithFields(log.Fields{"label": label}).Debug("Marking label as used")
	this.usedLabels[label] = varDecl.VariableIdentifier()
}

func (this *TypeChecker) VarDeclIsKnown(varDecl interfaces.VarDecl) bool {
	if isKnown, exists := this.knownIdentifiers[varDecl.VariableIdentifier()]; exists {
		return isKnown
	}

	return false
}

func (this *TypeChecker) MarkVarIDAsKnown(varID interfaces.VarID) {
	log.WithFields(log.Fields{"VarDecl": varID}).Debug("Marking VarDecl as known")
	this.knownIdentifiers[varID] = true
}

// MarkVarIDAsUnknown stores that the VarID is currently unknown, if VarID remains unmarked it is a reference to a undefined question
func (this *TypeChecker) MarkVarIDAsUnknown(varID interfaces.VarID) {
	log.WithFields(log.Fields{"VarDecl": varID}).Debug("Marking VarDecl as unknown")
	this.knownIdentifiers[varID] = false
}

// KnownIdentifiers returns a map where the keys are identifiers that are currently known
func (this *TypeChecker) KnownIdentifiers() map[interfaces.VarID]bool {
	return this.knownIdentifiers
}
