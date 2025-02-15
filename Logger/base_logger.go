package main

type BaseLogger struct {
	next Logger
}

func (l *BaseLogger) SetNext(n Logger) {
	l.next = n
}

func (l *BaseLogger) Log(level LogLevel, message string) {
	if l.next != nil {
		l.next.Log(level, message)
	}
}
