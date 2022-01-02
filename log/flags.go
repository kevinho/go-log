package log

import "github.com/urfave/cli"

var (
	LevelFlag   = cli.IntFlag{Name: "log.level", Value: int(LvlInfo)}
	ModulesFlag = cli.StringFlag{Name: "log.module", Value: ""}
	PathFlag    = cli.StringFlag{Name: "log.path", Value: "logs"}
)
