package dbindex

import (
	"errors"
	"flag"
	"fmt"
	sutil "github.com/hidai620/go-utils/stringutil"
	"os"
	"strings"
)

// アウトプットタイプ
type OutputType int

const (
	_ OutputType = iota
	CONSOLE
	CSV
)

func GetOutputType(o string) (OutputType, error) {
	switch strings.ToUpper(o) {
	case "CONSOLE":
		return CONSOLE, nil
	case "CSV":
		return CSV, nil
	default:
		return 0, errors.New(fmt.Sprintf("Pselese select type following list. console, csv"))
	}
}

func (o OutputType) String() string {
	var ret string
	switch o {
	case CONSOLE:
		ret = "CONSOLE"
	case CSV:
		ret = "CSV"
	}
	return ret
}

// Parse parses command line flags.
func ParseCommandLineOption() (*Option, error) {
	var config, out, tableNames string
	flag.StringVar(&config, "config", "config.toml", "Absolute or relrative path of config file.")
	flag.StringVar(&out, "out", "console", `Output type of result. "console" or "csv"`)
	flag.StringVar(&tableNames, "table", "", `Analyze Target table name.`)
	flag.Parse()

	err := existsFile(config)
	if err != nil {
		return nil, err
	}

	outputType, err := GetOutputType(out)
	if err != nil {
		return nil, err
	}

	tableNamesArray := sutil.Split(tableNames, ",")
	return NewOption(outputType, config, tableNamesArray), nil
}

// CommandLineOption has command line arguments.
type Option struct {
	Out        OutputType
	ConfigPath string
	TableNames []string
}

// New returns CommandLineOption created with arguments
func NewOption(out OutputType, configPath string, tableNames []string) *Option {
	return &Option{
		Out:        out,
		ConfigPath: configPath,
		TableNames: tableNames,
	}
}

// 同じ値を持つ場合、trueを返す
func (c *Option) Equals(c2 *Option) bool {
	return c.Out == c2.Out &&
		c.ConfigPath == c2.ConfigPath
}

// existsFile returns error if argument filePath does not exist.
func existsFile(filePath string) error {
	_, err := os.Stat(string(filePath))
	if err != nil {
		return errors.New(fmt.Sprintf("Specified file does not exist. %s", filePath))
	}
	return nil
}