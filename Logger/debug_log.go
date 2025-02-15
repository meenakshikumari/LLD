package main

import "fmt"

type DebugLog struct {
	BaseLogger
}

func NewDebugLog() *DebugLog {
	return &DebugLog{}
}

func (d *DebugLog) Log(level LogLevel, message string) {
	if level == Debug {
		fmt.Println("[Debug] " + message)
	} else {
		d.next.Log(level, message)
	}
}

func (d *DebugLog) SetNext(n Logger) {
	d.next = n
}
