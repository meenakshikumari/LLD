package main

type Logger interface {
	Log(LogLevel, string)
	SetNext(next Logger)
}

func CreateChainOfResp() Logger {
	infoLog := NewInfoLog()
	dLog := NewDebugLog()
	eLog := NewErrorLog()

	infoLog.SetNext(dLog)
	dLog.SetNext(eLog)

	return infoLog
}
