package cmd

import (
	"io"
	"os"
)

//counterfeiter:generate . UI

type UI interface {
	DisplayText(text string, values ...map[string]interface{})
}

type TerminalUI struct {
	Out   io.Writer
	Error io.Writer
}

func NewTerminalUI() *TerminalUI {
	return &TerminalUI{
		Out:   os.Stdout,
		Error: os.Stderr,
	}
}

func NewTestTerminalUI(out io.Writer, err io.Writer) *TerminalUI {
	return &TerminalUI{
		Out:   out,
		Error: err,
	}
}

func (ui *TerminalUI) DisplayText(text string, values ...map[string]interface{}) {
	// will do something
}
