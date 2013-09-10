package main

import (
	"errors"
	"fmt"
	"os"
	"regexp"
	"strings"
)

var optParseRegex = regexp.MustCompile("^[\\-]{1,2}(.*)")

const (
	FINISHED = "finished"
)

type OptParser struct {
	CurrentKey, CurrentValue string
	KnownKeys                []string
	Defaults                 map[string]string
	Values                   map[string]string
	Descriptions             map[string]string
}

func (self *OptParser) AddDescription(key, description string) {
	if self.Descriptions == nil {
		self.Descriptions = map[string]string{key: description}
	} else {
		self.Descriptions[key] = description
	}
}

func (self *OptParser) PrintDefaults() {
	table := NewTable()
	table.Add([]string{"USAGE: " + os.Args[0]})
	for _, key := range self.KnownKeys {
		value := self.Defaults[key]
		line := []string{"  --" + key}
		ds := strings.Split(self.Descriptions[key], "\n")
		line = append(line, ds[0])
		table.Add(line)

		for _, d := range ds[1:] {
			table.Add([]string{"", d})
		}

		if value != "" && value != "true" {
			table.Add([]string{"", "DEFAULT: " + value})
		}

		if len(ds) > 1 || (value != "" && value != "true") {
			table.Add([]string{})
		}
	}
	lines := table.Lines()
	fmt.Println(strings.Join(lines, "\n"))
}

func (self *OptParser) AddKey(key, usage string) {
	self.KnownKeys = append(self.KnownKeys, key)
	self.AddDescription(key, usage)
}

func (self *OptParser) Add(key, defaultValue, usage string) {
	if self.Defaults == nil {
		self.Defaults = map[string]string{}
	}
	self.Defaults[key] = defaultValue
	self.AddKey(key, usage)
}

func NewOptParser() *OptParser {
	return &OptParser{
		Defaults: map[string]string{},
		Values:   map[string]string{},
	}
}

func (self *OptParser) ProcessAll(args []string) (e error) {
	self.Values = map[string]string{}
	for _, arg := range args {
		if e = self.Process(arg); e != nil {
			return
		}
	}
	self.ProcessKey(FINISHED)
	return
}

func (self *OptParser) Get(key string) string {
	return self.Values[key]
}

func (self *OptParser) Process(arg string) (e error) {
	matches := optParseRegex.FindStringSubmatch(arg)
	if len(matches) == 2 {
		e = self.ProcessKey(matches[1])
	} else {
		e = self.ProcessValue(arg)
	}
	return
}

func (self *OptParser) ProcessKey(key string) (e error) {
	if self.CurrentKey != "" {
		if defaultValue, ok := self.Defaults[self.CurrentKey]; ok {
			self.Values[self.CurrentKey] = defaultValue
		} else if self.CurrentKey != FINISHED {
			e = errors.New("no value set for key " + self.CurrentKey)
		}
	}
	parts := strings.Split(key, "=")
	if len(parts) > 1 {
		self.CurrentKey = parts[0]
		e = self.ProcessValue(strings.Join(parts[1:], "="))
	} else {
		self.CurrentKey = key
	}
	return
}

func (self *OptParser) ProcessValue(value string) (e error) {
	if self.CurrentKey != "" {
		self.Values[self.CurrentKey] = value
	} else {
		e = errors.New("unable to process " + value)
	}
	self.CurrentKey = ""
	return
}
