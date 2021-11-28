package slinterfaces

import (
	kitlevel "github.com/go-kit/log/level"
)

type PrintLevel int

const (
	PrintNone  PrintLevel = 0
	PrintInfo  PrintLevel = 1
	PrintDebug PrintLevel = 2
)

// main interface for the SimpleLogger
type ISimpleLogger interface {

	//Print To Screen functions
	GetPrintToScreen() PrintLevel
	SetPrintToScreen(PrintLevel)

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

	OpenSessionFileLog(logfilename string, sessionid string) error
	GetSessionIDs() []string

	CloseChannel(sessionid string) []error
	CloseAllChannels() []error

	OpenChannel(sessionid string) []error
	OpenAllChannels() []error

	AddChannel(log ISimpleChannel)
	GetChannel(sessionid string) ISimpleChannel
	GetChannels() map[string]ISimpleChannel
	SetChannelLogLevel(sessionid string, lvl kitlevel.Option)
}
