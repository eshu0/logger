package shellframework

import (
	"fmt"
	"os"

	"github.com/eshu0/simplelogger/interfaces"
	kitlog "github.com/go-kit/kit/log"
	kitlevel "github.com/go-kit/kit/log/level"
)

type SimpleLogger struct {
	slinterfaces.SimpleLogger
	loglevel kitlevel.Option
	log      kitlog.Logger
}

//
// Simple Logging
//
// these function provide logging to the choosen logfile
//

func NewSimpleLogger(logger kitlog.Logger) ShellLogger {
	ssl := SimpleLogger{}
	ssl.log = logger
	return ssl
}

//func (ssl *ShellLogger) SetLogPrefix(prefix string) {
//ssl.log = kitlog.With(ssl.log, "session_id", session.ID())
//}

func (ssl *SimpleLogger) SetLog(log kitlog.Logger) {
	ssl.log = log
}

func (ssl *SimpleLogger) GetLog() kitlog.Logger {
	return ssl.log
}

func (ssl *SimpleLogger) SetLogLevel(lvl kitlevel.Option) {
	ssl.loglevel = lvl
	// have to set the filter for the level
	ssl.log = kitlevel.NewFilter(ssl.log, lvl)
}

func (ssl *SimpleLogger) GetLogLevel() kitlevel.Option {
	return ssl.loglevel
}

// the logging functions are here
func (ssl *SimpleLogger) LogDebug(cmd string, data ...interface{}) {
	kitlevel.Debug(ssl.log).Log("cmd", cmd, "data", fmt.Sprintf("%s", data))
}

func (ssl *SimpleLogger) LogWarn(cmd string, data ...interface{}) {
	kitlevel.Warn(ssl.log).Log("cmd", cmd, "data", fmt.Sprintf("%s", data))
}

func (ssl *SimpleLogger) LogInfo(cmd string, data ...interface{}) {
	kitlevel.Info(ssl.log).Log("cmd", cmd, "data", fmt.Sprintf("%s", data))
}
func (ssl *SimpleLogger) LogError(cmd string, data ...interface{}) {
	kitlevel.Error(ssl.log).Log("cmd", cmd, "data", fmt.Sprintf("%s", data))
}

// the logging functions are here
func (ssl *SimpleLogger) LogDebugf(cmd string, msg string, data ...interface{}) {
	kitlevel.Debug(ssl.log).Log("cmd", cmd, "data", fmt.Sprintf(msg, data...))
}

func (ssl *SimpleLogger) LogWarnf(cmd string, msg string, data ...interface{}) {
	kitlevel.Warn(ssl.log).Log("cmd", cmd, "data", fmt.Sprintf(msg, data...))
}

func (ssl *SimpleLogger) LogInfof(cmd string, msg string, data ...interface{}) {
	kitlevel.Info(ssl.log).Log("cmd", cmd, "data", fmt.Sprintf(msg, data...))
}
func (ssl *SimpleLogger) LogErrorf(cmd string, msg string, data ...interface{}) {
	kitlevel.Error(ssl.log).Log("cmd", cmd, "data", fmt.Sprintf(msg, data...))
}

func (ssl *SimpleLogger) OpenSessionFileLog(logfilename string,sessionid string) *os.File {
	f, err := os.OpenFile(logfilename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	// check error
	if err != nil {
		panic(err)
	}

	//logger :=
	logger := kitlog.NewLogfmtLogger(f)                                            //(f, session.ID()+" ", log.LstdFlags)
	logger = kitlog.With(logger, "session_id", sessionid, "ts", kitlog.DefaultTimestampUTC) //, "caller", kitlog.DefaultCaller)

	// check log is valid
	if logger == nil {
		panic("logger is nil")
	}

	ssl.log = logger

	// default to show everything
	ssl.SetLogLevel(kitlevel.AllowAll())

	//ssl.loglevel = -1
	return f
}
