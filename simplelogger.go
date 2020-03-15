package simplelogger

import (
	"fmt"

	sl "github.com/eshu0/simplelogger/interfaces"
	//kitlog "github.com/go-kit/kit/log"
	kitlevel "github.com/go-kit/kit/log/level"
)

type SimpleLogger struct {

	//inherit from interface
	sl.ISimpleLogger

	// use kitlevel API
	globallevel kitlevel.Option

	//Let's make an array of logging outputs
	channels map[string]sl.ISimpleChannel

}

//
// Simple Logging
//
// these function provide logging to the choosen logfile
//

func NewSimpleLogger(filename string, sessionid string) SimpleLogger {

	ssl := SimpleLogger{}

	channels := make(map[string]sl.ISimpleChannel)

	lg := &SimpleChannel{}
	lg.SetFileName(filename)
	lg.SetSessionID(sessionid)

	channels[lg.sessionid] = lg

	ssl.channels = channels

	return ssl
}

/*
 		SIMPLE LOG CHANNELS
*/

func (ssl *SimpleLogger) AddChannel(log sl.ISimpleChannel) {
	ssl.channels[log.GetSessionID()] = log
}

func (ssl *SimpleLogger) GetChannel(sessionid  string) sl.ISimpleChannel {
	return ssl.channels[sessionid]
}

func (ssl *SimpleLogger) GetChannels() map[string]sl.ISimpleChannel {
	return ssl.channels
}

func (ssl *SimpleLogger) GetSessionIDs() []string {
	var keys []string
	for k := range ssl.channels {
	    keys = append(keys, k)
	}
	return keys
}



func (ssl *SimpleLogger) SetChannelLogLevel(sessionid string,lvl kitlevel.Option) {
	// have to set the filter for the level
	for _, channel := range ssl.channels {

		if(sessionid ==""){
			channel.SetLogLevel(lvl)
		}else{
			if(channel.GetSessionID() == sessionid){
				channel.SetLogLevel(lvl)

			}
		}
	}
}

func (ssl *SimpleLogger) GetChannelLogLevel(sessionid string) kitlevel.Option {
	for _, channel := range ssl.channels {
			if(channel.GetSessionID() == sessionid) {
				return channel.GetLogLevel()
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
				channel.Close()
		}else{
			if(channel.GetSessionID() == sessionid){
					channel.Close()
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
			channel.Open()
		}else{
			if(channel.GetSessionID() == sessionid){
				channel.Open()
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

	channel := &SimpleChannel{}
	channel.SetFileName(logfilename)
	channel.SetSessionID(sessionid)
	channel.Open()

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
		kitlevel.Debug(channel.GetLog()).Log("cmd", cmd, "data", fmt.Sprintf("%s", data))
	}
}

func (ssl *SimpleLogger) LogWarn(cmd string, data ...interface{}) {
	for _, channel := range ssl.channels {
		kitlevel.Warn(channel.GetLog()).Log("cmd", cmd, "data", fmt.Sprintf("%s", data))
	}
}

func (ssl *SimpleLogger) LogInfo(cmd string, data ...interface{}) {
for _, channel := range ssl.channels {
		kitlevel.Info(channel.GetLog()).Log("cmd", cmd, "data", fmt.Sprintf("%s", data))
	}
}
func (ssl *SimpleLogger) LogError(cmd string, data ...interface{}) {
	for _, channel := range ssl.channels {
		kitlevel.Error(channel.GetLog()).Log("cmd", cmd, "data", fmt.Sprintf("%s", data))
	}
}

// the logging functions are here
func (ssl *SimpleLogger) LogDebugf(cmd string, msg string, data ...interface{}) {
	for _, channel := range ssl.channels {
		kitlevel.Debug(channel.GetLog()).Log("cmd", cmd, "data", fmt.Sprintf(msg, data...))
	}
}

func (ssl *SimpleLogger) LogWarnf(cmd string, msg string, data ...interface{}) {
	for _, channel := range ssl.channels {
		kitlevel.Warn(channel.GetLog()).Log("cmd", cmd, "data", fmt.Sprintf(msg, data...))
	}
}

func (ssl *SimpleLogger) LogInfof(cmd string, msg string, data ...interface{}) {
	for _, channel := range ssl.channels {
		kitlevel.Info(channel.GetLog()).Log("cmd", cmd, "data", fmt.Sprintf(msg, data...))
	}
}
func (ssl *SimpleLogger) LogErrorf(cmd string, msg string, data ...interface{}) {
	for _, channel := range ssl.channels {
		kitlevel.Error(channel.GetLog()).Log("cmd", cmd, "data", fmt.Sprintf(msg, data...))
	}
}
