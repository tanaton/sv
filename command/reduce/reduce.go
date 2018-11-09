package reduce

import (
	"bufio"
	"flag"
	"fmt"
	"strconv"
	"strings"

	"github.com/tanaton/sv/command"
)

type ReduceArg struct {
	Num   int
	A     float64
	Cell  int
	Delim string
}

func Reduce(c *command.CLI, args []string) int {
	cmd := ReduceArg{}
	fs := flag.NewFlagSet("sv-reduce", flag.ExitOnError)
	fs.SetOutput(c.Err)
	fs.IntVar(&cmd.Num, "n", 1, "num")
	fs.Float64Var(&cmd.A, "a", 0, "a")
	fs.IntVar(&cmd.Cell, "c", 0, "cell")
	fs.StringVar(&cmd.Delim, "d", ",", "delim")
	fs.Parse(args)

	scanner := bufio.NewScanner(c.In)
	// header
	if scanner.Scan() == false {
		fmt.Fprintln(c.Err, scanner.Err())
		return 1
	}
	line := scanner.Text()
	fmt.Fprintln(c.Out, line)

	// data
	count := 0
	var nold float64
	for scanner.Scan() {
		line = scanner.Text()
		if cmd.Num > 1 {
			if count%cmd.Num == 0 {
				fmt.Fprintln(c.Out, line)
			}
		} else if cmd.A > 1e-8 {
			cel := strings.Split(line, cmd.Delim)
			no, _ := strconv.ParseFloat(cel[cmd.Cell], 64)
			if no >= nold+cmd.A {
				nold = no
				fmt.Fprintln(c.Out, line)
			}
		}
		count++
	}
	return 0
}
