package simplelogger

import (
	"fmt"
	"os"

	sl "github.com/eshu0/simplelogger/interfaces"
	kitlog "github.com/go-kit/kit/log"
	kitlevel "github.com/go-kit/kit/log/level"
)


type Channel struct{

	// session id
	sessionid string

	//Let's make an array of logging outputs
	log kitlog.Logger

	// filename for the log
	filename string

	fileptr *os.File

	// use kitlevel API
	level kitlevel.Option

}

type SimpleLogger struct {

	//inherit from interface
	sl.ISimpleLogger

	// use kitlevel API
	globallevel kitlevel.Option

	//Let's make an array of logging outputs
	channels map[string]*Channel

}

//
// Simple Logging
//
// these function provide logging to the choosen logfile
//

func NewSimpleLogger(filename string, sessionid string) SimpleLogger {

	ssl := SimpleLogger{}

	channels := make(map[string]*Channel)

	lg := Channel{}
	lg.SetFileName(filename)
	lg.SetSessionID(sessionid)

	channels[lg.sessionid] = &lg

	ssl.channels = channels

	return ssl
}

/*
 		SIMPLE LOG CHANNELS
*/

func (ssl *SimpleLogger) AddChannel(log Channel) {
	ssl.channels[log.sessionid] = &log
}

func (ssl *SimpleLogger) GetChannel(sessionid  string) *Channel {
	return ssl.channels[sessionid]
}

func (ssl *SimpleLogger) GetChannels() map[string]*Channel {
	return ssl.channels
}

func (ssl *SimpleLogger) SetChannelLogLevel(sessionid string,lvl kitlevel.Option) {
	// have to set the filter for the level
	for _, channel := range ssl.channels {

		if(sessionid ==""){
			channel.level = lvl
			channel.log = kitlevel.NewFilter(channel.log, lvl)
		}else{
			if(channel.sessionid == sessionid){
				channel.level = lvl
				channel.log = kitlevel.NewFilter(channel.log, lvl)
			}
		}
	}
}

func (ssl *SimpleLogger) GetChannelLogLevel(sessionid string) kitlevel.Option {
	for _, channel := range ssl.channels {
			if(channel.sessionid == sessionid) {
				return channel.level
			}
	}
	return nil
}

/*
 		SIMPLE LOG FUNCTIONS
*/


func (ssl *SimpleLogger) CloseChannel(sessionid string) {
	// have to set the filter for the level
	for _, channel := range ssl.channels {

		if(sessionid ==""){
			if(channel.fileptr != nil){
				channel.fileptr.Close()
			}
		}else{
			if(channel.sessionid == sessionid){
				if(channel.fileptr != nil){
					channel.fileptr.Close()
				}
			}
		}
	}
}

func (ssl *SimpleLogger) CloseAllChannels() {
	ssl.CloseChannel("")
}

func (ssl *SimpleLogger) OpenChannel(sessionid string) {
	// have to set the filter for the level
	for _, channel := range ssl.channels {

		if(sessionid ==""){
			channel.OpenFileLog()
		}else{
			if(channel.sessionid == sessionid){
				channel.OpenFileLog()
			}
		}
	}
}

func (ssl *SimpleLogger) OpenAllChannels() {
	ssl.OpenChannel("")
}

func (ssl *SimpleLogger) SetLogLevel(lvl kitlevel.Option) {
	ssl.globallevel = lvl
	ssl.SetChannelLogLevel("",lvl)
}

func (ssl *SimpleLogger) GetLogLevel() kitlevel.Option {
	return ssl.globallevel
}

func (ssl *SimpleLogger) OpenSessionFileLog(logfilename string, sessionid string) {

	channel := Channel{}
	channel.SetFileName(logfilename)
	channel.SetSessionID(sessionid)
	channel.OpenFileLog()

	ssl.AddChannel(channel)

	// default to show everything
	ssl.SetLogLevel(kitlevel.AllowAll())
}

/*
 		LOGGING after here
*/

// the logging functions are here
func (ssl *SimpleLogger) LogDebug(cmd string, data ...interface{}) {
	for _, channel := range ssl.channels {
		kitlevel.Debug(channel.log).Log("cmd", cmd, "data", fmt.Sprintf("%s", data))
	}
}

func (ssl *SimpleLogger) LogWarn(cmd string, data ...interface{}) {
	for _, channel := range ssl.channels {
		kitlevel.Warn(channel.log).Log("cmd", cmd, "data", fmt.Sprintf("%s", data))
	}
}

func (ssl *SimpleLogger) LogInfo(cmd string, data ...interface{}) {
for _, channel := range ssl.channels {
		kitlevel.Info(channel.log).Log("cmd", cmd, "data", fmt.Sprintf("%s", data))
	}
}
func (ssl *SimpleLogger) LogError(cmd string, data ...interface{}) {
	for _, channel := range ssl.channels {
		kitlevel.Error(channel.log).Log("cmd", cmd, "data", fmt.Sprintf("%s", data))
	}
}

// the logging functions are here
func (ssl *SimpleLogger) LogDebugf(cmd string, msg string, data ...interface{}) {
	for _, channel := range ssl.channels {
		kitlevel.Debug(channel.log).Log("cmd", cmd, "data", fmt.Sprintf(msg, data...))
	}
}

func (ssl *SimpleLogger) LogWarnf(cmd string, msg string, data ...interface{}) {
	for _, channel := range ssl.channels {
		kitlevel.Warn(channel.log).Log("cmd", cmd, "data", fmt.Sprintf(msg, data...))
	}
}

func (ssl *SimpleLogger) LogInfof(cmd string, msg string, data ...interface{}) {
	for _, channel := range ssl.channels {
		kitlevel.Info(channel.log).Log("cmd", cmd, "data", fmt.Sprintf(msg, data...))
	}
}
func (ssl *SimpleLogger) LogErrorf(cmd string, msg string, data ...interface{}) {
	for _, channel := range ssl.channels {
		kitlevel.Error(channel.log).Log("cmd", cmd, "data", fmt.Sprintf(msg, data...))
	}
}


/*
 Channel Functions after here
*/

func (lo *Channel) SetFileName(filename string) {
	lo.filename = filename
}

func (lo *Channel) GetFileName() string{
	return lo.filename
}

func (lo *Channel) SetSessionID(sessionid string) {
	lo.sessionid = sessionid
}

func (lo *Channel) GetSessionID() string {
	return lo.sessionid
}

func (lo *Channel) OpenFileLog() {

	f, err := os.OpenFile(lo.filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	// check error
	if err != nil {
		panic(err)
	}

	//logger :=
	logger := kitlog.NewLogfmtLogger(f)                                                         //(f, session.ID()+" ", log.LstdFlags)
	logger = kitlog.With(logger, "session_id", lo.sessionid, "ts", kitlog.DefaultTimestampUTC) //, "caller", kitlog.DefaultCaller)

	// check log is valid
	if logger == nil {
		panic("logger is nil")
	}

	lo.fileptr = f

}
