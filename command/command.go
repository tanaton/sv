package command

import (
	"io"
)

const Delim = ","

type CLI struct {
	In  io.Reader
	Out io.Writer
	Err io.Writer
}

var Cell map[string]int

func init() {
	Cell = make(map[string]int)
	index := 0
	for code := 'a'; code <= 'z'; code++ {
		Cell[string([]rune{code})] = index
		index++
	}
	for base := 'a'; base <= 'z'; base++ {
		for code := 'a'; code <= 'z'; code++ {
			Cell[string([]rune{base, code})] = index
			index++
		}
	}
}
