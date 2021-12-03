package cmd

import (
	"fmt"
	"io"
	"os"
	"strings"
	"text/template"
)

//counterfeiter:generate . UI

// UI is the interface to interact with the user.
type UI interface {
	// DisplayText renders text to the terminal.
	DisplayText(text string, values ...map[string]interface{})
}

// TerminalUI is an implementation of UI.
type TerminalUI struct {
	Out   io.Writer
	Error io.Writer
}

// NewTerminalUI returns an instance of TerminalUI that attaches to STDOUT and
// STDERR.
func NewTerminalUI() *TerminalUI {
	return &TerminalUI{
		Out:   os.Stdout,
		Error: os.Stderr,
	}
}

// NewTestTerminalUI returns a TerminalUI with the provided io.Writers for out and error.
func NewTestTerminalUI(out io.Writer, err io.Writer) *TerminalUI {
	return &TerminalUI{
		Out:   out,
		Error: err,
	}
}

func (ui *TerminalUI) DisplayText(text string, values ...map[string]interface{}) {
	fmt.Fprintln(ui.Out, ui.process(text, values...))
}

func (ui *TerminalUI) process(text string, values ...map[string]interface{}) string {
	mapped := ui.firstOrEmpty(values)

	t := template.Must(template.New("letter").Parse(text))
	output := new(strings.Builder)
	err := t.Execute(output, mapped)
	if err != nil {
		panic(
			fmt.Errorf(
				"error processing text:\nText: %s\nValues: %#v\nError: %w",
				text,
				mapped,
				err,
			),
		)
	}

	return output.String()
}

// firstOrEmpty returns the first map otherwise it returns the empty map.
func (*TerminalUI) firstOrEmpty(list []map[string]interface{}) map[string]interface{} {
	if len(list) == 0 {
		return map[string]interface{}{}
	}
	return list[0]
}
