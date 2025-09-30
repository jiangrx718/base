package commands

import (
	"base/commands/generate"
	"base/commands/gorm"
	"base/commands/migrate"

	"github.com/urfave/cli/v2"
)

func All() []*cli.Command {
	commands := []*cli.Command{
		migrate.Command(),
		generate.Command(),
		gorm.Command(),
	}
	return commands
}
