package main

import (
	"os"
	"fmt"
	"time"
)

var LogColors = map[int]int {
	DEBUG: 102,
	INFO: 28,
	WARN: 214,
	ERROR: 196,
}

const TIME_FORMAT = "2006-01-02T15:04:05.000000"

func colorize(c int, s string) (r string) {
	return fmt.Sprintf("\033[38;5;%dm%s\033[0m", c, s)
}

const (
	DEBUG = iota
	INFO
	WARN
	ERROR
)

var LogPrefixes = map[int]string {
	DEBUG:	"DEBUG",
	INFO:	"INFO ",
	WARN:	"WARN ",
	ERROR:	"ERROR",
}

type Logger struct {
	LogLevel int
	Prefix string
}

func (l* Logger) Fdebug(format string, n ...interface{}) {
	l.Flog(DEBUG, format, n...)
}

func (l* Logger) Finfo(format string, n ...interface{}) {
	l.Flog(INFO, format, n...)
}

func (l* Logger) Fwarn(format string, n ...interface{}) {
	l.Flog(WARN, format, n...)
}

func (l* Logger) Ferror(format string, n ...interface{}) {
	l.Flog(ERROR, format, n...)
}

func (l* Logger) Flog(level int, s string, n ...interface{}) {
	if level >= l.LogLevel {
		l.Write(l.LogPrefix(level), fmt.Sprintf(s,  n...))
	}
}

func (l* Logger) Debug(n ...interface{}) {
	l.Log(DEBUG, n...)
}

func (l* Logger) Info(n ...interface{}) {
	l.Log(INFO, n...)
}

func (l* Logger) Warn(n ...interface{}) {
	l.Log(WARN, n...)
}

func (l* Logger) Error(n ...interface{}) {
	l.Log(ERROR, n...)
}

func (l* Logger) LogPrefix(i int) (s string ) {
	s = time.Now().Format(TIME_FORMAT)
	if l.Prefix != "" {
		s = s + " [" + l.Prefix + "]"
	}
	s = s + " " + l.LogLevelPrefix(i)
	return
}

func (l* Logger) LogLevelPrefix(level int) (s string) {
	color := LogColors[level]
	prefix := LogPrefixes[level]
	return colorize(color, prefix)
}

func (l* Logger) Log(level int, n ...interface{}) {
	if level >= l.LogLevel {
		all := append([]interface{} { l.LogPrefix(level) }, n...)
		l.Write(all...)
	}
}

func (self *Logger) Write(n... interface{}) {
	fmt.Fprintln(os.Stderr, n...)
}
