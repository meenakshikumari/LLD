package main

import "fmt"

type InfoLogger struct {
	BaseLogger
}

func NewInfoLog() *InfoLogger {
	return &InfoLogger{}
}

func (infoLog *InfoLogger) Log(level LogLevel, message string) {
	if level == Info {
		fmt.Println("[INFO] " + message)
	} else if infoLog.next != nil {
		infoLog.next.Log(level, message)
	}
}

func (infoLog *InfoLogger) SetNext(n Logger) {
	infoLog.next = n
}
