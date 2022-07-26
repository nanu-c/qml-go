package main

import (
	"fmt"
	"os"

	"github.com/nanu-c/qml-go"
)

func main() {
	if err := qml.Run(run); err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
}

func run() error {
	qml.LoadTranslator("fr_FR", "qml", "_", "i18n", ".qm")
	engine := qml.NewEngine()
	component, err := engine.LoadFile("main.qml")
	if err != nil {
		return err
	}
	window := component.CreateWindow(nil)
	window.Show()
	window.Wait()
	return nil
}
