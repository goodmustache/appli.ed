package cmd

type PositionalArgs struct {
	ConfigPath string `positional-arg-name:"PATH_TO_CONFIG" required:"true" description:"Path to the appli.ed configuration file."`
}

type Execute struct {
	Positional PositionalArgs `positional-args:"yes"`

	UI UI
}

func (e Execute) Execute(args []string) error {
	e.UI.DisplayText(
		"Reading config '{{.Path}}'",
		map[string]interface{}{
			"Path": e.Positional.ConfigPath,
		},
	)
	return nil
}
