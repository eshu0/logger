package slinterfaces

import (
	//kitlog "github.com/go-kit/kit/log"
	kitlevel "github.com/go-kit/kit/log/level"
)

// main interface for the SimpleLogger
type ISimpleLogger interface {
	//Log Level functions
	GetLogLevel() kitlevel.Option
	SetLogLevel(kitlevel.Option)

	//Print To Screen functions
	GetPrintToScreen() bool
	SetPrintToScreen(bool)

	//Log Level PrintToScreen functions
	GetPrintToScreenLogLevel() kitlevel.Option
	SetPrintToScreenLogLevel(kitlevel.Option)

	LogErrorf(cmd string, message string, data ...interface{})
	LogWarnf(cmd string, message string, data ...interface{})
	LogInfof(cmd string, message string, data ...interface{})
	LogDebugf(cmd string, message string, data ...interface{})

	LogError(cmd string, data ...interface{})
	LogErrorE(cmd string, data error)
	LogWarn(cmd string, data ...interface{})
	LogInfo(cmd string, data ...interface{})
	LogDebug(cmd string, data ...interface{})

	OpenSessionFileLog(logfilename string, sessionid string)
	GetSessionIDs() []string

	CloseChannel(sessionid string)
	CloseAllChannels()

 	OpenChannel(sessionid string)
	OpenAllChannels()

	AddChannel(log ISimpleChannel)
	GetChannel(sessionid  string) ISimpleChannel
	GetChannels() map[string]ISimpleChannel
 	SetChannelLogLevel(sessionid string,lvl kitlevel.Option)

}
