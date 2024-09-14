package cmd

type Flag struct {
	Name       string
	Usage      string
	IsRequired bool
}

const (
	FlagNameDriver   = "driver"
	FlagNameHost     = "host"
	FlagNameUsername = "username"
	FlagNamePassword = "password"
	FlagNameDatabase = "database"
)
