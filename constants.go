package main

const timeFormat = "2006-01-02 15:04:05Z"
const VERSION = "0.1.5"

var (
	GITCOMMIT string
	parser    = NewOptParser()
)
