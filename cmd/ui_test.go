package cmd_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "github.com/onsi/gomega/gbytes"

	. "github.com/goodmustache/appli.ed/cmd"
)

var _ = Describe("UI", func() {
	var (
		ui *TerminalUI
	)

	BeforeEach(func() {
		ui = NewTestTerminalUI(NewBuffer(), NewBuffer())
	})

	Describe("DisplayText", func() {
		It("displays text to the user", func() {
			passedTypes := map[string]interface{}{
				"SomeText":      "foo",
				"SomeOtherText": 2,
			}
			ui.DisplayText("A {{.SomeText}}, {{.SomeOtherText}}", passedTypes)

			Expect(ui.Out).To(Say("A foo, 2\n"))
		})
	})
})
