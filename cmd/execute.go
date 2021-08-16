package cmd

type PositionalArgs struct {
	Path string `positional-arg-name:"PATH_TO_SEARCH" required:"true" description:"The path of the directory containing the *.ed directories"`
}

type Execute struct {
	Positional PositionalArgs `positional-args:"yes"`

	UI UI
}

func (e Execute) Execute(args []string) error {
	return nil
}
