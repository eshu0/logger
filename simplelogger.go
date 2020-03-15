package simplelogger

import (
	"fmt"
	"os"

	sl "github.com/eshu0/simplelogger/interfaces"
	kitlog "github.com/go-kit/kit/log"
	kitlevel "github.com/go-kit/kit/log/level"
)

type SimpleLogger struct {
	//inherit from interface
	sl.ISimpleLogger

	// use kitlevel API
	loglevel kitlevel.Option

	//Let's make an array of logging outputs
	log []kitlog.Logger

	// filename for the log
	logfilename string

	// session id
	sessionid string
}

//
// Simple Logging
//
// these function provide logging to the choosen logfile
//

func NewSimpleLogger(logger kitlog.Logger) SimpleLogger {
	ssl := SimpleLogger{}

	logs := []kitlog.Logger{}
	logs = append(logs, logger)

	ssl.log = logs
	return ssl
}

func NewSimpleLoggerWithFilename(filename string) SimpleLogger {
	ssl := SimpleLogger{}
	ssl.logfilename = filename
	logs := []kitlog.Logger{}
	ssl.log = logs
	return ssl
}

//func (ssl *ShellLogger) SetLogPrefix(prefix string) {
//ssl.log = kitlog.With(ssl.log, "session_id", session.ID())
//}

// Deprecated: SetLog exists for historical compatibility
// and should not be used. Use AddLog instead
func (ssl *SimpleLogger) SetLog(log kitlog.Logger) {
	ssl.AddLog(log)
}

func (ssl *SimpleLogger) AddLog(log kitlog.Logger) {
	logs := ssl.log
	logs = append(logs, log)
	ssl.log = logs
}

// Deprecated: GetLog exists for historical compatibility
// and should not be used. Use GetLogs instead
func (ssl *SimpleLogger) GetLog() kitlog.Logger {
	return ssl.log[0]
}

func (ssl *SimpleLogger) GetLogs() []kitlog.Logger {
	return ssl.log
}

func (ssl *SimpleLogger) SetLogLevel(lvl kitlevel.Option) {
	ssl.loglevel = lvl
	// have to set the filter for the level
	for i := 0; i < len(ssl.log); i++ {
		ssl.log[i] = kitlevel.NewFilter(ssl.log[i], lvl)
	}
}

func (ssl *SimpleLogger) GetLogLevel() kitlevel.Option {
	return ssl.loglevel
}

// the logging functions are here
func (ssl *SimpleLogger) LogDebug(cmd string, data ...interface{}) {
	for i := 0; i < len(ssl.log); i++ {
		kitlevel.Debug(ssl.log[i]).Log("cmd", cmd, "data", fmt.Sprintf("%s", data))
	}
}

func (ssl *SimpleLogger) LogWarn(cmd string, data ...interface{}) {
	for i := 0; i < len(ssl.log); i++ {
		kitlevel.Warn(ssl.log[i]).Log("cmd", cmd, "data", fmt.Sprintf("%s", data))
	}
}

func (ssl *SimpleLogger) LogInfo(cmd string, data ...interface{}) {
	for i := 0; i < len(ssl.log); i++ {
		kitlevel.Info(ssl.log[i]).Log("cmd", cmd, "data", fmt.Sprintf("%s", data))
	}
}
func (ssl *SimpleLogger) LogError(cmd string, data ...interface{}) {
	for i := 0; i < len(ssl.log); i++ {
		kitlevel.Error(ssl.log[i]).Log("cmd", cmd, "data", fmt.Sprintf("%s", data))
	}
}

// the logging functions are here
func (ssl *SimpleLogger) LogDebugf(cmd string, msg string, data ...interface{}) {
	for i := 0; i < len(ssl.log); i++ {
		kitlevel.Debug(ssl.log[i]).Log("cmd", cmd, "data", fmt.Sprintf(msg, data...))
	}
}

func (ssl *SimpleLogger) LogWarnf(cmd string, msg string, data ...interface{}) {
	for i := 0; i < len(ssl.log); i++ {
		kitlevel.Warn(ssl.log[i]).Log("cmd", cmd, "data", fmt.Sprintf(msg, data...))
	}
}

func (ssl *SimpleLogger) LogInfof(cmd string, msg string, data ...interface{}) {
	for i := 0; i < len(ssl.log); i++ {
		kitlevel.Info(ssl.log[i]).Log("cmd", cmd, "data", fmt.Sprintf(msg, data...))
	}
}
func (ssl *SimpleLogger) LogErrorf(cmd string, msg string, data ...interface{}) {
	for i := 0; i < len(ssl.log); i++ {
		kitlevel.Error(ssl.log[i]).Log("cmd", cmd, "data", fmt.Sprintf(msg, data...))
	}
}

/*
func (ssl *SimpleLogger) OpenTimeNowFileLog(logfolder string, sessionid string) *os.File {
	currentTime := time.Now()
	return OpenSessionFileLog(logfilename, currentTime.Format("2006-01-02-15-04-05"), )

}
*/

func (ssl *SimpleLogger) OpenSessionFileLog(logfilename string, sessionid string) *os.File {
	ssl1 := NewSimpleLoggerWithFilename(logfilename)
	ssl = &ssl1
	return ssl.OpenFileLog()
}

func (ssl *SimpleLogger) OpenFileLog() *os.File {
	f, err := os.OpenFile(ssl.logfilename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	// check error
	if err != nil {
		panic(err)
	}

	//logger :=
	logger := kitlog.NewLogfmtLogger(f)                                                         //(f, session.ID()+" ", log.LstdFlags)
	logger = kitlog.With(logger, "session_id", ssl.sessionid, "ts", kitlog.DefaultTimestampUTC) //, "caller", kitlog.DefaultCaller)

	// check log is valid
	if logger == nil {
		panic("logger is nil")
	}

	ssl.AddLog(logger)

	// default to show everything
	ssl.SetLogLevel(kitlevel.AllowAll())

	//ssl.loglevel = -1
	return f
}
