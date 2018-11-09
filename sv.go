package main

import (
	"os"

	"./command"
	"./command/concat"
	"./command/extract"
	"./command/number"
	"./command/reduce"
)

func main() {
	cli := command.CLI{
		In:  os.Stdin,
		Out: os.Stdout,
		Err: os.Stderr,
	}
	var code int
	if len(os.Args) >= 2 {
		switch os.Args[1] {
		case "reduce":
			code = reduce.Reduce(&cli, os.Args[2:])
		case "cat", "concat":
			code = concat.Concat(&cli, os.Args[2:])
		case "extract", "ext":
			code = extract.Extract(&cli, os.Args[2:])
		case "num", "number":
			code = number.Number(&cli, os.Args[2:])
		default:
			code = 1
		}
	} else {
		code = 1
	}
	os.Exit(code)
}
