package main

import (
	"fmt"
	"github.com/kevinho/go-log/log"
	"io"
	"os"

	"github.com/mattn/go-colorable"
	"github.com/mattn/go-isatty"
)

func SetupLogger(printOrigin bool, level int, vmodule string, cfg *log.RotateConfig) error {
	log.PrintOrigins(printOrigin)
	fmt.Println("printOrigin", printOrigin)

	verbosity := log.Lvl(level)
	log.Verbosity(verbosity)
	err := log.Vmodule(vmodule)
	if err != nil {
		return err
	}

	if cfg != nil {
		//terminal handler
		usecolor := (isatty.IsTerminal(os.Stdout.Fd()) || isatty.IsCygwinTerminal(os.Stdout.Fd())) && os.Getenv("TERM") != "dumb"
		output := io.Writer(os.Stdout)
		if usecolor {
			output = colorable.NewColorableStderr()
		}

		rfh := log.NewFileRotateHandler(cfg, log.TerminalFormat(usecolor))
		ostream := log.StreamHandler(output, log.TerminalFormat(usecolor))
		log.GRoot().SetHandler(log.MultiHandler(ostream, rfh))
	}

	return nil
}

func PrintOrigin(yesOrNot bool) {
	log.PrintOrigins(yesOrNot)
}

func Verbosity(level int) {
	log.Verbosity(log.Lvl(level))
}

func Vmodule(pattern string) error {
	return log.Vmodule(pattern)
}
