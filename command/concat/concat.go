package concat

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"sort"
	"strings"

	"github.com/tanaton/sv/command"
)

type ConcatArg struct {
	Num   int
	Find  string
	Delim string
	count int
}

func Concat(c *command.CLI, args []string) int {
	cmd := ConcatArg{}
	fs := flag.NewFlagSet("sv-concat", flag.ExitOnError)
	fs.SetOutput(c.Err)
	fs.IntVar(&cmd.Num, "n", 1, "num")
	fs.StringVar(&cmd.Find, "f", "", "find")
	fs.StringVar(&cmd.Delim, "d", ",", "delim")
	fs.Parse(args)

	cmd.count = 0
	for _, arg := range fs.Args() {
		var err error
		if arg == "-" {
			err = cmd.cat(c, c.In)
		} else {
			err = cmd.glob(c, arg)
		}
		if err != nil {
			fmt.Fprintln(c.Err, err)
			return 1
		}
	}
	return 0
}

func (cmd *ConcatArg) cat(c *command.CLI, r io.Reader) error {
	var w io.Writer
	scanner := bufio.NewScanner(r)
	if cmd.count == 0 {
		w = c.Out
	} else {
		w = ioutil.Discard
	}
	if cmd.Find == "" {
		cmd.line(scanner, w)
	} else {
		cmd.find(scanner, w)
	}
	for scanner.Scan() {
		fmt.Fprintln(c.Out, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		return err
	}
	cmd.count++
	return nil
}

func (cmd *ConcatArg) glob(c *command.CLI, arg string) error {
	pl, err := filepath.Glob(arg)
	if err != nil {
		return err
	}
	sort.Strings(pl)
	for _, p := range pl {
		err := func() error {
			fp, err := os.Open(p)
			if err != nil {
				return err
			}
			defer fp.Close()
			return cmd.cat(c, fp)
		}()
		if err != nil {
			return err
		}
	}
	return nil
}

func (cmd *ConcatArg) line(scanner *bufio.Scanner, w io.Writer) {
	linecount := 0
	if cmd.Num > 0 {
		for scanner.Scan() {
			fmt.Fprintln(w, scanner.Text())
			linecount++
			if linecount >= cmd.Num {
				break
			}
		}
	}
}

func (cmd *ConcatArg) find(scanner *bufio.Scanner, w io.Writer) {
	for scanner.Scan() {
		line := scanner.Text()
		fmt.Fprintln(w, line)
		if strings.Index(line, cmd.Find) == 0 {
			break
		}
	}
}
