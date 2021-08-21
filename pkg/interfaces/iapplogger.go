package slinterfaces

// main interface for the SimpleLogger
type IAppLogger interface {
	LogErrorf(cmd string, message string, data ...interface{})
	LogWarnf(cmd string, message string, data ...interface{})
	LogInfof(cmd string, message string, data ...interface{})
	LogDebugf(cmd string, message string, data ...interface{})

	LogError(cmd string, data ...interface{})
	LogErrorE(cmd string, data error)
	LogErrorEf(cmd string, message string, e error)
	LogWarn(cmd string, data ...interface{})
	LogInfo(cmd string, data ...interface{})
	LogDebug(cmd string, data ...interface{})

	StartLogging()
	FinishLogging()
}
