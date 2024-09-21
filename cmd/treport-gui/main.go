package main

import (
	"fmt"
	"os"

	"github.com/leonzag/treport/internal"
)

func main() {
	app, err := internal.NewAppGUI()
	if err != nil {
		fatal(err)
	}
	if err := app.ShowAndRun(); err != nil {
		fatal(err)
	}
}

func fatal(err error) {
	fmt.Fprintln(os.Stderr, err)
	os.Exit(1)
}
