package simplelogger

import (
	sli "github.com/eshu0/simplelogger/pkg/interfaces"
)

type AppLogger struct {

	//inherit from interface
	sli.IAppLogger

	Log     sli.ISimpleLogger `json:"-"`
	Started bool              `json:"-"`
}

/*
	START/FINISH LOG FUNCTIONS
*/

// the logging functions are here
func (al *AppLogger) StartLogging() {

	if al.Log == nil {
		log := NewApplicationNowLogger()

		// lets open a file log using the session
		log.OpenAllChannels()

		al.Log = log
	} else {
		al.Log.OpenAllChannels()
	}

	al.Started = true
}

// the logging functions are here
func (al *AppLogger) FinishLogging() {

	if al.Log != nil {
		al.Log.CloseAllChannels()
	}
}

/*
	APP LOG FUNCTIONS
*/

// the logging functions are here
func (al *AppLogger) LogDebug(cmd string, data ...interface{}) {
	if !al.Started {
		al.StartLogging()
	}
	if al.Log != nil {
		al.Log.LogDebug(cmd, data...)
	}
}

func (al *AppLogger) LogWarn(cmd string, data ...interface{}) {
	if !al.Started {
		al.StartLogging()
	}
	if al.Log != nil {
		al.Log.LogWarn(cmd, data...)
	}
}

func (al *AppLogger) LogInfo(cmd string, data ...interface{}) {
	if !al.Started {
		al.StartLogging()
	}
	if al.Log != nil {
		al.Log.LogInfo(cmd, data...)
	}
}

func (al *AppLogger) LogError(cmd string, data ...interface{}) {
	if !al.Started {
		al.StartLogging()
	}
	if al.Log != nil {
		al.Log.LogError(cmd, data...)
	}
}

// This Log error allows errors to be logged .Error() is the data written
func (al *AppLogger) LogErrorE(cmd string, data error) {
	if !al.Started {
		al.StartLogging()
	}
	if al.Log != nil {
		al.Log.LogErrorE(cmd, data)
	}
}

// This Log error allows errors to be logged .Error() is the data written
func (al *AppLogger) LogErrorEf(cmd string, msg string, e error) {
	if !al.Started {
		al.StartLogging()
	}
	if al.Log != nil {
		al.Log.LogErrorEf(cmd, msg, e)
	}
}

// the logging functions are here
func (al *AppLogger) LogDebugf(cmd string, msg string, data ...interface{}) {
	if !al.Started {
		al.StartLogging()
	}
	if al.Log != nil {
		al.Log.LogDebugf(cmd, msg, data...)
	}
}

func (al *AppLogger) LogWarnf(cmd string, msg string, data ...interface{}) {
	if !al.Started {
		al.StartLogging()
	}
	if al.Log != nil {
		al.Log.LogWarnf(cmd, msg, data...)
	}
}

func (al *AppLogger) LogInfof(cmd string, msg string, data ...interface{}) {
	if !al.Started {
		al.StartLogging()
	}
	if al.Log != nil {
		al.Log.LogInfof(cmd, msg, data...)
	}
}

func (al *AppLogger) LogErrorf(cmd string, msg string, data ...interface{}) {
	if !al.Started {
		al.StartLogging()
	}
	if al.Log != nil {
		al.Log.LogErrorf(cmd, msg, data...)
	}
}
