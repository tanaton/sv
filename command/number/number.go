package number

import (
	"bufio"
	"flag"
	"fmt"

	"../../command"
)

type NumberArg struct {
	Delim string
}

func Number(c *command.CLI, args []string) int {
	cmd := NumberArg{}
	fs := flag.NewFlagSet("sv-number", flag.ExitOnError)
	fs.SetOutput(c.Err)
	fs.StringVar(&cmd.Delim, "d", ",", "delim")
	fs.Parse(args)

	scanner := bufio.NewScanner(c.In)
	// ヘッダー
	if scanner.Scan() == false {
		fmt.Fprintln(c.Err, scanner.Err())
		return 1
	}
	fmt.Fprintf(c.Out, "No%s%s\n", cmd.Delim, scanner.Text())
	// データ
	num := 0
	for scanner.Scan() {
		fmt.Fprintf(c.Out, "%d%s%s\n", num, cmd.Delim, scanner.Text())
		num++
	}
	return 0
}
