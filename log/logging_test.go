package log

import (
	"io"
	"io/ioutil"
	"os"
	"testing"

	"github.com/mattn/go-colorable"
	"github.com/mattn/go-isatty"
	"github.com/stretchr/testify/assert"
)

var (
	ostream Handler
	glogger *GlogHandler
)

func TestMain(m *testing.M) {
	PrintOrigins(true)

	usecolor := (isatty.IsTerminal(os.Stderr.Fd()) || isatty.IsCygwinTerminal(os.Stderr.Fd())) && os.Getenv("TERM") != "dumb"
	output := io.Writer(os.Stderr)
	if usecolor {
		output = colorable.NewColorableStderr()
	}
	ostream = StreamHandler(output, TerminalFormat(usecolor))

	glogger = NewGlogHandler(ostream)
	glogger.Verbosity(LvlInfo)

	Root().SetHandler(glogger)

	os.Exit(m.Run())
}

func TestUsage(t *testing.T) {
	//simple usage, lv ignore
	Debug("simple usage")

	//new logger context
	newLogger := New("contextKey", "contextValue")
	newLogger.Info("msg", "key", "value")
}

func TestVmodule(t *testing.T) {
	//new logger
	newLogger := New()

	//set verbosity to 3
	glogger.Verbosity(LvlInfo)

	//specifically set logging/* package to 5(less Severity)
	err := glogger.Vmodule("logging/*=5")
	if err != nil {
		t.Fatal(err)
	}

	//try to print debug level
	newLogger.Debug("output", "should seen", "logging/*=5")
}

func TestFileRotate(t *testing.T) {
	config := defaultConfig
	config.MaxSize = 1

	//clean
	_ = os.RemoveAll(config.LogDir)

	//setup file
	rfh := NewFileRotateHandler(defaultConfig, TerminalFormat(false))

	glogger.SetHandler(MultiHandler(DiscardHandler(), rfh))

	//simple usage
	newLogger := New("contextKey", "contextvalue")

	buf := make([]byte, 0x1000)
	for i := 0; i < 200; i++ {
		newLogger.Info("context info", "value", buf)
	}

	files, _ := ioutil.ReadDir(config.LogDir)

	//two file: newfile, oldfile compress
	assert.Equal(t, len(files), 2)

	//clean
	_ = os.RemoveAll(config.LogDir)
}
