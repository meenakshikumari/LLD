package main

import "fmt"

type ErrorLog struct {
	*BaseLogger
}

func NewErrorLog() *ErrorLog {
	return &ErrorLog{}
}

func (d *ErrorLog) Log(level LogLevel, message string) {
	if level == Error {
		fmt.Println("[Error] " + message)
	} else {
		d.next.Log(level, message)
	}
}

func (d *ErrorLog) SetNext(n Logger) {
	d.next = n
}
