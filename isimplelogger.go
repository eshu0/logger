package slinterfaces

import (
	"os"

	kitlog "github.com/go-kit/kit/log"
	kitlevel "github.com/go-kit/kit/log/level"
)

// main interface for the SimpleLogger
type ISimpleLogger interface {
	// logging
	GetLog() kitlog.Logger
	SetLog(log kitlog.Logger)

	//log functions
	GetLogLevel() kitlevel.Option
	SetLogLevel(kitlevel.Option)
	//SetLogPrefix(string)

	LogErrorf(cmd string, message string, data ...interface{})
	LogWarnf(cmd string, message string, data ...interface{})
	LogInfof(cmd string, message string, data ...interface{})
	LogDebugf(cmd string, message string, data ...interface{})

	LogError(cmd string, data ...interface{})
	LogWarn(cmd string, data ...interface{})
	LogInfo(cmd string, data ...interface{})
	LogDebug(cmd string, data ...interface{})

	// This opens a session
	OpenSessionFileLog(session ISession) *os.File
}
