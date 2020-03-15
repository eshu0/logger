package slinterfaces

import (
	//kitlog "github.com/go-kit/kit/log"
	kitlevel "github.com/go-kit/kit/log/level"
)

// main interface for the SimpleLogger
type ISimpleLogger interface {
	//log functions
	GetLogLevel() kitlevel.Option
	SetLogLevel(kitlevel.Option)

	LogErrorf(cmd string, message string, data ...interface{})
	LogWarnf(cmd string, message string, data ...interface{})
	LogInfof(cmd string, message string, data ...interface{})
	LogDebugf(cmd string, message string, data ...interface{})

	LogError(cmd string, data ...interface{})
	LogWarn(cmd string, data ...interface{})
	LogInfo(cmd string, data ...interface{})
	LogDebug(cmd string, data ...interface{})

	OpenSessionFileLog(logfilename string, sessionid string)

	CloseChannel(sessionid string)
	CloseAllChannels()

 	OpenChannel(sessionid string)
	OpenAllChannels()

	AddChannel(log ISimpleChannel)
	GetChannel(sessionid  string) ISimpleChannel
	GetChannels() map[string]ISimpleChannel
 	SetChannelLogLevel(sessionid string,lvl kitlevel.Option)

}
