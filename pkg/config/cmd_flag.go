package config

type CmdFlag []string

func (i *CmdFlag) String() string {
	return "my string representation"
}

func (i *CmdFlag) Set(value string) error {
	*i = append(*i, value)
	return nil
}
