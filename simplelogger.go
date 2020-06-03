package simplelogger

import (
	"fmt"
	"os"
	"math/rand"
	"time"

	sl "github.com/eshu0/simplelogger/interfaces"
	kitlevel "github.com/go-kit/kit/log/level"
)

type SimpleLogger struct {

	//inherit from interface
	sl.ISimpleLogger

	// use kitlevel API
	globallevel kitlevel.Option

	//Let's make an array of logging outputs
	channels map[string]sl.ISimpleChannel

	printtoscreen sl.PrintLevel 
}

//
// Simple Logging
//
// these function provide logging to the choosen logfile
//

// This is the simplest application log generator
// The os.args[0] is used for filename and the session is random
func NewApplicationLogger() *SimpleLogger {
	return NewApplicationSessionLogger(RandomSessionID())
}

// This is application log generator when the session is required
// The os.args[0] is used for filename
func NewApplicationSessionLogger(sessionid string) *SimpleLogger {
	
	appname, err:= os.Executable()
	
	if(err != nil) {
		appname = "unknown"
	}
	
	return NewSimpleLogger(appname + ".log", sessionid)
}

// This lets you specify the filename and the session
func NewSimpleLogger(filename string, sessionid string) *SimpleLogger {

	ssl := &SimpleLogger{}

	channels := make(map[string]sl.ISimpleChannel)

	lg := &SimpleChannel{}
	lg.SetFileName(filename)
	lg.SetSessionID(sessionid)

	channels[lg.sessionid] = lg

	ssl.channels = channels

	// by default we print everything being logged to the screen
	ssl.SetPrintToScreen(sl.PrintInfo)
	//ssl.SetPrintToScreenLogLevel(kitlevel.AllowAll())

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


// Generates Random session string
func RandomSessionID() string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	seededRand := rand.New(rand.NewSource(time.Now().UnixNano()))
	b := make([]byte, 6)
	for i := range b {
		b[i] = charset[seededRand.Intn(len(charset))]
	}
	return string(b)
}

func (ssl *SimpleLogger) GetPrintToScreen() sl.PrintLevel {
	return ssl.printtoscreen
}

func (ssl *SimpleLogger) SetPrintToScreen(toggle sl.PrintLevel){
	ssl.printtoscreen = toggle
}

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

// I am not sure i like this method too much however it works
func log(ssl *SimpleLogger, lvl string, cmd string, msg string, data ...interface{}) {

	for _, channel := range ssl.channels {
		log := channel.GetLog()

		if log != nil {
			switch lvl {
				case "debug" :
							kitlevel.Debug(log).Log("cmd", cmd, "data", fmt.Sprintf("%s", data...))	
							printscreen(ssl, lvl, cmd, fmt.Sprintf("%s", data...))
							break	
				case "warn" :
							kitlevel.Warn(log).Log("cmd", cmd, "data", fmt.Sprintf("%s", data...))
							printscreen(ssl, lvl, cmd, fmt.Sprintf("%s", data...))
							break	
				case "info" :
							kitlevel.Info(log).Log("cmd", cmd, "data", fmt.Sprintf("%s", data...))	
							printscreen(ssl, lvl, cmd, fmt.Sprintf("%s", data...))
							break	
				case "error" :
							kitlevel.Error(log).Log("cmd", cmd, "data", fmt.Sprintf("%s", data...))
							printscreen(ssl, lvl, cmd, fmt.Sprintf("%s", data...))
							break		
				case "debugf" :
							kitlevel.Debug(log).Log("cmd", cmd, "data", fmt.Sprintf(msg, data...))
							printscreen(ssl, lvl, cmd, fmt.Sprintf(msg, data...) )
							break
				case "warnf" :
							kitlevel.Warn(log).Log("cmd", cmd, "data", fmt.Sprintf(msg, data...))
							printscreen(ssl, lvl, cmd, fmt.Sprintf(msg, data...) )
							break
				case "infof" :
							kitlevel.Info(log).Log("cmd", cmd, "data", fmt.Sprintf(msg, data...))
							printscreen(ssl, lvl, cmd, fmt.Sprintf(msg, data...) )
							break
			  	case "errorf" :
							kitlevel.Error(log).Log("cmd", cmd, "data", fmt.Sprintf(msg, data...))
							printscreen(ssl, lvl, cmd, fmt.Sprintf(msg, data...) )
							break
			}
		} else {
			printscreen(ssl, lvl, "log", fmt.Sprintf("log nil %s",channel.GetSessionID()))
		}
	}

}

func printscreen(ssl *SimpleLogger, lvl string, cmd string, msg string) {


	if ssl.GetPrintToScreen() != sl.PrintNone {
		switch lvl {
			case "debug" :
				if ssl.GetPrintToScreen() != sl.PrintDebug {
						fmt.Println(fmt.Sprintf("[%s] [%s] - %s", lvl,cmd, msg))
				}
			case "warn" :
						fmt.Println(fmt.Sprintf("[%s] [%s] - %s", lvl,cmd, msg))
			case "info" :
						fmt.Println(fmt.Sprintf("[%s] [%s] - %s", lvl,cmd, msg))
			case "error" :
						fmt.Println(fmt.Sprintf("[%s] [%s] - %s", lvl,cmd, msg))
			case "debugf" :
				if ssl.GetPrintToScreen() != sl.PrintDebug {
						fmt.Println(fmt.Sprintf("[%s] [%s] - %s", lvl,cmd, msg))
				}
			case "warnf" :
						fmt.Println(fmt.Sprintf("[%s] [%s] - %s", lvl,cmd, msg))
			case "infof" :
						fmt.Println(fmt.Sprintf("[%s] [%s] - %s", lvl,cmd, msg))
			case "errorf" :
						fmt.Println(fmt.Sprintf("[%s] [%s] - %s", lvl,cmd, msg))
		}
	}

}

// the logging functions are here
func (ssl *SimpleLogger) LogDebug(cmd string, data ...interface{}) {
	log(ssl, "debug", cmd, "%s", data)
}

func (ssl *SimpleLogger) LogWarn(cmd string, data ...interface{}) {
	log(ssl, "warn", cmd,  "%s", data)
}

func (ssl *SimpleLogger) LogInfo(cmd string, data ...interface{}) {
	log(ssl, "info", cmd, "%s",  data)
}

func (ssl *SimpleLogger) LogError(cmd string, data ...interface{}) {
	log(ssl, "error", cmd,  "%s", data)
}

// This Log error allows errors to be logged .Error() is the data written
func (ssl *SimpleLogger) LogErrorE(cmd string, data error) {
	log(ssl, "error", cmd,  "%s", data.Error())
}

// the logging functions are here
func (ssl *SimpleLogger) LogDebugf(cmd string, msg string, data ...interface{}) {
	log(ssl, "debugf", cmd, msg, data...)
}

func (ssl *SimpleLogger) LogWarnf(cmd string, msg string, data ...interface{}) {
	log(ssl, "warnf", cmd, msg, data...)
}

func (ssl *SimpleLogger) LogInfof(cmd string, msg string, data ...interface{}) {
	log(ssl, "infof", cmd, msg, data...)
}

func (ssl *SimpleLogger) LogErrorf(cmd string, msg string, data ...interface{}) {
	log(ssl, "errorf", cmd, msg, data...)
}
