package convert

type convert0609 struct {
	version     string
	commandPath string
}

func (c convert0609) Version() string {
	return c.version
}

func (c convert0609) Execute(params []string) (string, error) {
	return "", nil
}
