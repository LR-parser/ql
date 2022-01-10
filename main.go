package main

import (
	"fmt"
	log "github.com/Sirupsen/logrus"
	"github.com/codegangsta/cli"
	"io/ioutil"
	"os"
	"ql/gui"
	"ql/interfaces"
	"ql/lexer"
	"ql/parser"
	"ql/symbols"
	"ql/typechecker"
)

// main sets up the CLI command, defaults to trying to load, parse and present form.ql file in current working directory
func main() {
	app := cli.NewApp()
	app.Name = "QL"
	app.Usage = "generate a questionnaire form from a DSL"
	app.Version = "1.0.0"

	var filePathPassed string
	var debugModeEnabled bool

	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:        "filepath",
			Value:       "form.ql",
			Usage:       "path to a .ql file",
			Destination: &filePathPassed,
		},
		cli.BoolFlag{
			Name:        "debugmode",
			Usage:       "if passed, the minimum log level displayed is set to debug",
			Destination: &debugModeEnabled,
		},
	}

	app.Action = func(c *cli.Context) {
		filePath := filePathPassed
		if c.NArg() > 0 {
			filePath = c.Args()[0]
		}

		if debugModeEnabled {
			enableDebugLevelLogs()
		} else {
			enableFatalLevelLogs()
		}

		log.WithFields(log.Fields{"filePath": filePath}).Info("Passed file path as CLI argument/flag")

		// if no file at path exists, exit program
		exitIfFilePathIsInvalid(filePath)

		initQL(filePath)
	}

	app.Run(os.Args)
}

// exitIfFilePathIsInvalid exits program if no file exists at filePath or filePath refers to dir
func exitIfFilePathIsInvalid(filePath string) {
	// no file at path exists
	if src, err := os.Stat(filePath); os.IsNotExist(err) {
		fmt.Printf("Could not continue: no file exists at path: %s \n", filePath)
		os.Exit(1)
	} else if src.IsDir() {
		fmt.Printf("Could not continue: path to directory instead of file passed: %s\n", filePath)
		os.Exit(2)
	}

	log.WithFields(log.Fields{"filePath": filePath}).Info("Confirmed that path to file is valid")
}

// enableDebugLevelLogs sets the minimum level of logs displayed to debug
func enableDebugLevelLogs() {
	log.SetLevel(log.DebugLevel)
	log.Info("Enabled debug level logging")
}

// enableDebugLevelLogs sets the minimum level of logs displayed to fatal
func enableFatalLevelLogs() {
	log.SetLevel(log.FatalLevel)
}

// initQL loads the contents of the file at the passed file path, starts lexer/parser and type checking and presents the GUI
func initQL(filePath string) {
	gui := gui.NewGUI()

	fileContent, _ := ioutil.ReadFile(filePath)

	log.WithFields(log.Fields{"filePath": filePath}).Info("Loaded file")

	parsedForm, parseError := lexAndParse(fileContent)

	if parseError != nil {
		errors := []error{parseError}
		log.WithFields(log.Fields{"parseError": parseError}).Info("Adding parse error to errors list displayed in GUI")

		// only show the error dialog
		gui.ShowWindow(errors, nil)
		return
	}

	typeCheckErrors, typeCheckWarnings := conductTypeChecking(parsedForm)

	varIDValueSymbols := symbolsWithDefaultValuesForVarIDs(parsedForm)

	gui.InitializeGUIForm(parsedForm, varIDValueSymbols)
	gui.ShowWindow(typeCheckErrors, typeCheckWarnings)
}

// lexAndParse conducts lexical analysis and parses the past file content, returns a parsed form and a potential error
func lexAndParse(fileContent []byte) (interfaces.Form, error) {
	lexer := lexer.NewLexer(fileContent)

	parser := parser.NewParser()
	log.Info("Initiating parsing of file")
	parseResult, parseErr := parser.Parse(lexer)

	if parseErr != nil {
		log.WithFields(log.Fields{"parseErr": parseErr}).Error("Could not parse")
	}

	log.WithFields(log.Fields{"Result": parseResult}).Info("Form parsed")

	parsedForm, parseResultIsForm := parseResult.(interfaces.Form)
	if parseErr == nil && !parseResultIsForm {
		log.Panic("Parse result could not be casted to Form")
	}

	return parsedForm, parseErr
}

// conductTypeChecking starts type checking the passed form for errors and warnings
func conductTypeChecking(form interfaces.Form) ([]error, []error) {
	typeChecker := typechecker.NewTypeChecker()
	typeCheckerArgs := typechecker.NewTypeCheckArgs(typeChecker, symbols.NewTypeCheckSymbols())
	form.TypeCheck(&typeCheckerArgs)
	warnings := typeChecker.EncounteredWarnings()
	errors := typeChecker.EncounteredErrors()

	log.WithFields(log.Fields{"errors": errors, "warnings": warnings}).Error("Type checking finished")

	return errors, warnings
}

// symbolsWithDefaultValuesForVarIDs creates new symbols and starts process that sets default values for all VarIDs
func symbolsWithDefaultValuesForVarIDs(form interfaces.Form) interfaces.VarIDValueSymbols {
	varIDDefaultValueVisitor := symbols.NewDefaultVarIDValueVisitor()
	return varIDDefaultValueVisitor.StartSettingDefaultValuesForVarIDs(form)
}
