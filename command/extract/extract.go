package extract

import (
	"bufio"
	"flag"
	"fmt"
	"strings"

	"github.com/tanaton/sv/command"
)

type ExtractArg struct {
	Cell  string
	Delim string
	clist []string
}

func Extract(c *command.CLI, args []string) int {
	cmd := ExtractArg{}
	fs := flag.NewFlagSet("sv-extract", flag.ExitOnError)
	fs.SetOutput(c.Err)
	fs.StringVar(&cmd.Cell, "c", "", "cell")
	fs.StringVar(&cmd.Delim, "d", ",", "delim")
	fs.Parse(args)

	cmd.clist = strings.Split(strings.ToLower(cmd.Cell), ",")

	scanner := bufio.NewScanner(c.In)
	// header
	if scanner.Scan() == false {
		fmt.Fprintln(c.Err, scanner.Err())
		return 1
	}
	fmt.Fprintln(c.Out, cmd.extractLine(scanner.Text()))
	for scanner.Scan() {
		fmt.Fprintln(c.Out, cmd.extractLine(scanner.Text()))
	}
	return 0
}

func (cmd *ExtractArg) extractLine(line string) string {
	out := []string{}
	list := strings.Split(line, cmd.Delim)
	l := len(list)
	for _, it := range cmd.clist {
		index, ok := command.Cell[it]
		if ok && l > index {
			out = append(out, list[index])
		}
	}
	return strings.Join(out, cmd.Delim)
}
