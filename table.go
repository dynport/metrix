package main

import (
	"fmt"
	"strings"
)

type Table struct {
	Columns [][]string
	Lengths map[int]int
}

func NewTable() *Table {
	return &Table{
		Columns: [][]string {},
		Lengths: map[int]int {},
	}
}

func (self *Table) Lines() (lines []string) {
	for _, col := range self.Columns {
		cl := []string{}
		for i, v := range col {
			cl = append(cl, fmt.Sprintf("%-*s", self.Lengths[i], v))
		}
		lines = append(lines, strings.Join(cl, "\t"))
	}
	return
}

func (self *Table) Add(cols []string) {
	for i, v := range cols {
		if self.Lengths[i] < len(v) {
			self.Lengths[i] = len(v)
		}
	}
	self.Columns = append(self.Columns, cols)
}
