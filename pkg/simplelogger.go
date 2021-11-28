package simplelogger

import (
	"fmt"
	"math/rand"
	"os"
	"time"

	sli "github.com/eshu0/logger/pkg/interfaces"
	kitlevel "github.com/go-kit/log/level"
)

const (
	LOG_EXTENSION = ".log"
	CMD_STRING    = "cmd"
	DATA_STRING   = "data"
	CASE_DEBUG    = "debug"
	CASE_WARN     = "warn"
	CASE_INFO     = "info"
	CASE_ERROR    = "error"
	CASE_DEBUGF   = "debugf"
	CASE_WARNF    = "warnf"
	CASE_INFOF    = "infof"
	CASE_ERRORF   = "errorf"
)

type SimpleLogger struct {

	//inherit from interface
	sli.ISimpleLogger

	// use kitlevel API
	globallevel kitlevel.Option

	//Let's make an array of logging outputs
	channels map[string]sli.ISimpleChannel

	printtoscreen sli.PrintLevel
}

//
// Simple Logging
//
// these function provide logging to the choosen logfile
//

// This is the simplest application log generator
// The os.args[0] is used for filename and the session is random
func NewApplicationLogger() SimpleLogger {
	return NewApplicationSessionLogger(RandomSessionID())
}

// This is application log generator when the session is required
// The os.args[0] is used for filename
func NewApplicationSessionLogger(sessionid string) SimpleLogger {

	appname, err := os.Executable()

	if err != nil {
		appname = "unknown"
	}

	return NewSimpleLogger(appname+LOG_EXTENSION, sessionid)
}

func NewApplicationNowLogger() SimpleLogger {
	return NewAppSessionNowLogger(RandomSessionID())
}

func NewAppSessionNowLogger(sessionid string) SimpleLogger {

	appname, err := os.Executable()

	if err != nil {
		appname = "unknown"
	}
	filename := appname + "-" + time.Now().Format("2006-01-02-15-04-05")
	return NewSimpleLogger(filename+LOG_EXTENSION, sessionid)
}

func NewApplicationDayLogger() SimpleLogger {
	return NewAppSessionDayLogger(RandomSessionID())
}

func NewAppSessionDayLogger(sessionid string) SimpleLogger {

	appname, err := os.Executable()

	if err != nil {
		appname = "unknown"
	}
	filename := appname + "-" + time.Now().Format("2006-01-02")
	return NewSimpleLogger(filename+LOG_EXTENSION, sessionid)
}

// This lets you specify the filename and the session
func NewSimpleLogger(filename string, sessionid string) SimpleLogger {

	ssl := SimpleLogger{}

	channels := make(map[string]sli.ISimpleChannel)

	lg := SimpleChannel{filename: filename, sessionid: sessionid}
	//lg.SetFileName(filename)
	//lg.SetSessionID(sessionid)

	fmt.Println(lg.GetDetails())
	channels[lg.sessionid] = lg

	ssl.channels = channels

	// by default we print everything being logged to the screen
	ssl.SetPrintToScreen(sli.PrintInfo)
	return ssl
}

/*
	SIMPLE LOG CHANNELS
*/

func (ssl SimpleLogger) PrintDetails() {
	for key, channel := range ssl.channels {
		fmt.Printf("[%s] = %s ", key, channel.GetDetails())
	}
}

func (ssl SimpleLogger) AddChannel(log sli.ISimpleChannel) {
	ssl.channels[log.GetSessionID()] = log
}

func (ssl SimpleLogger) GetChannel(sessionid string) sli.ISimpleChannel {
	return ssl.channels[sessionid]
}

func (ssl SimpleLogger) GetChannels() map[string]sli.ISimpleChannel {
	return ssl.channels
}

func (ssl SimpleLogger) GetSessionIDs() []string {
	var keys []string
	for k := range ssl.channels {
		keys = append(keys, k)
	}
	return keys
}

func (ssl SimpleLogger) SetChannelLogLevel(sessionid string, lvl kitlevel.Option) {
	// have to set the filter for the level
	for _, channel := range ssl.channels {

		if len(sessionid) > 0 {
			if channel.GetSessionID() == sessionid {
				channel.SetLogLevel(lvl)
			}
		} else {
			channel.SetLogLevel(lvl)
		}
	}
}

func (ssl SimpleLogger) GetChannelLogLevel(sessionid string) kitlevel.Option {
	for _, channel := range ssl.channels {
		if channel.GetSessionID() == sessionid {
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

func (ssl SimpleLogger) GetPrintToScreen() sli.PrintLevel {
	return ssl.printtoscreen
}

func (ssl SimpleLogger) SetPrintToScreen(toggle sli.PrintLevel) {
	ssl.printtoscreen = toggle
}

func (ssl SimpleLogger) CloseChannel(sessionid string) []error {

	results := []error{}

	// have to set the filter for the level
	for _, channel := range ssl.channels {
		if len(sessionid) > 0 {
			if channel.GetSessionID() == sessionid {
				err := channel.Close()
				if err != nil {
					results = append(results, err)
				}
			}
		} else {
			err := channel.Close()
			if err != nil {
				results = append(results, err)
			}
		}
	}

	return results
}

func (ssl SimpleLogger) CloseAllChannels() []error {
	return ssl.CloseChannel("")
}

func (ssl SimpleLogger) OpenChannel(sessionid string) []error {
	results := []error{}
	for _, channel := range ssl.channels {
		// have to set the filter for the channel
		if len(sessionid) > 0 {
			if channel.GetSessionID() == sessionid {
				err := channel.Open()
				if err != nil {
					results = append(results, err)
				}
			}
		} else {
			err := channel.Open()
			if err != nil {
				results = append(results, err)
			}
		}
	}
	return results
}

func (ssl SimpleLogger) OpenAllChannels() []error {
	return ssl.OpenChannel("")
}

func (ssl SimpleLogger) SetLogLevel(lvl kitlevel.Option) {
	ssl.globallevel = lvl
	ssl.SetChannelLogLevel("", lvl)
}

func (ssl SimpleLogger) GetLogLevel() kitlevel.Option {
	return ssl.globallevel
}

func (ssl SimpleLogger) OpenSessionFileLog(logfilename string, sessionid string) error {

	channel := SimpleChannel{}
	channel.SetFileName(logfilename)
	channel.SetSessionID(sessionid)

	err := channel.Open()
	if err != nil {
		return err
	}

	ssl.AddChannel(channel)

	// default to show everything
	ssl.SetLogLevel(kitlevel.AllowAll())

	return nil
}

/*
	LOGGING after here
*/

// I am not sure i like this method too much however it works
func log(ssl SimpleLogger, lvl string, cmd string, msg string, data ...interface{}) {

	for _, channel := range ssl.channels {
		log := channel.GetLog()

		if log != nil {
			switch lvl {
			case CASE_DEBUG:
				kitlevel.Debug(log).Log(CMD_STRING, cmd, DATA_STRING, fmt.Sprintf("%s", data...))
				printscreen(ssl, lvl, cmd, fmt.Sprintf("%s", data...))
			case CASE_WARN:
				kitlevel.Warn(log).Log(CMD_STRING, cmd, DATA_STRING, fmt.Sprintf("%s", data...))
				printscreen(ssl, lvl, cmd, fmt.Sprintf("%s", data...))
			case CASE_INFO:
				kitlevel.Info(log).Log(CMD_STRING, cmd, DATA_STRING, fmt.Sprintf("%s", data...))
				printscreen(ssl, lvl, cmd, fmt.Sprintf("%s", data...))
			case CASE_ERROR:
				kitlevel.Error(log).Log(CMD_STRING, cmd, DATA_STRING, fmt.Sprintf("%s", data...))
				printscreen(ssl, lvl, cmd, fmt.Sprintf("%s", data...))
			case CASE_DEBUGF:
				kitlevel.Debug(log).Log(CMD_STRING, cmd, DATA_STRING, fmt.Sprintf(msg, data...))
				printscreen(ssl, lvl, cmd, fmt.Sprintf(msg, data...))
			case CASE_WARNF:
				kitlevel.Warn(log).Log(CMD_STRING, cmd, DATA_STRING, fmt.Sprintf(msg, data...))
				printscreen(ssl, lvl, cmd, fmt.Sprintf(msg, data...))
			case CASE_INFOF:
				kitlevel.Info(log).Log(CMD_STRING, cmd, DATA_STRING, fmt.Sprintf(msg, data...))
				printscreen(ssl, lvl, cmd, fmt.Sprintf(msg, data...))
			case CASE_ERRORF:
				kitlevel.Error(log).Log(CMD_STRING, cmd, DATA_STRING, fmt.Sprintf(msg, data...))
				printscreen(ssl, lvl, cmd, fmt.Sprintf(msg, data...))
			}
		} else {
			printscreen(ssl, lvl, "log", fmt.Sprintf("log nil %s", channel.GetSessionID()))
		}
	}

}

func printscreenfmt(lvl string, cmd string, msg string) {
	if len(msg) > 0 {
		fmt.Printf("%s: %s - %s\n", lvl, cmd, msg)
	} else {
		fmt.Printf("%s: - %s\n", lvl, cmd)
	}
}

func printscreen(ssl SimpleLogger, lvl string, cmd string, msg string) {
	if ssl.GetPrintToScreen() == sli.PrintNone {
		return
	}

	switch lvl {
	case CASE_DEBUG:
		if ssl.GetPrintToScreen() == sli.PrintDebug {
			printscreenfmt("Debug", cmd, msg)
		}
	case CASE_WARN:
		printscreenfmt("Warning", cmd, msg)
	case CASE_INFO:
		printscreenfmt("Info", cmd, msg)
	case CASE_ERROR:
		printscreenfmt("Error", cmd, msg)
	case CASE_DEBUGF:
		if ssl.GetPrintToScreen() == sli.PrintDebug {
			printscreenfmt("Debug", cmd, msg)
		}
	case CASE_WARNF:
		printscreenfmt("Warning", cmd, msg)
	case CASE_INFOF:
		printscreenfmt("Info", cmd, msg)
	case CASE_ERRORF:
		printscreenfmt("Error", cmd, msg)
	}

}

// the logging functions are here
func (ssl SimpleLogger) LogDebug(cmd string, data ...interface{}) {
	log(ssl, CASE_DEBUG, cmd, "%s", data...)
}

func (ssl SimpleLogger) LogWarn(cmd string, data ...interface{}) {
	log(ssl, CASE_WARN, cmd, "%s", data...)
}

func (ssl SimpleLogger) LogInfo(cmd string, data ...interface{}) {
	log(ssl, CASE_INFO, cmd, "%s", data...)
}

func (ssl SimpleLogger) LogError(cmd string, data ...interface{}) {
	log(ssl, CASE_ERROR, cmd, "%s", data...)
}

// This Log error allows errors to be logged .Error() is the data written
func (ssl SimpleLogger) LogErrorE(cmd string, e error) {
	log(ssl, CASE_ERROR, cmd, "%s", e.Error())
}

// This Log error allows errors to be logged where .Error() will be passed into the string
func (ssl SimpleLogger) LogErrorEf(cmd string, msg string, e error) {
	log(ssl, CASE_ERROR, cmd, msg, e.Error())
}

// the logging functions are here
func (ssl SimpleLogger) LogDebugf(cmd string, msg string, data ...interface{}) {
	log(ssl, CASE_DEBUGF, cmd, msg, data...)
}

func (ssl SimpleLogger) LogWarnf(cmd string, msg string, data ...interface{}) {
	log(ssl, CASE_WARNF, cmd, msg, data...)
}

func (ssl SimpleLogger) LogInfof(cmd string, msg string, data ...interface{}) {
	log(ssl, CASE_INFOF, cmd, msg, data...)
}

func (ssl SimpleLogger) LogErrorf(cmd string, msg string, data ...interface{}) {
	log(ssl, CASE_ERRORF, cmd, msg, data...)
}
