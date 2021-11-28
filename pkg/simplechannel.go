package simplelogger

import (
	"fmt"
	"os"

	sli "github.com/eshu0/logger/pkg/interfaces"
	kitlog "github.com/go-kit/log"
	kitlevel "github.com/go-kit/log/level"
)

// Simple Channel represents and output channel to be logged to
// kitlog does the hard work this simply wraps
type SimpleChannel struct {

	//inherit from interface
	sli.ISimpleChannel

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

/*
 Channel Functions after here
*/

func (lo SimpleChannel) SetFileName(filename string) {
	lo.filename = filename
}

func (lo SimpleChannel) GetFileName() string {
	return lo.filename
}

func (lo SimpleChannel) SetSessionID(sessionid string) {
	lo.sessionid = sessionid
}

func (lo SimpleChannel) GetSessionID() string {
	return lo.sessionid
}

func (lo SimpleChannel) SetLogLevel(lvl kitlevel.Option) {
	lo.level = lvl
	lo.log = kitlevel.NewFilter(lo.log, lvl)
}

func (lo SimpleChannel) GetLogLevel() kitlevel.Option {
	return lo.level
}

func (lo SimpleChannel) SetLog(log kitlog.Logger) {
	lo.log = log
}

func (lo SimpleChannel) GetLog() kitlog.Logger {
	return lo.log
}

func (lo SimpleChannel) Close() {
	if lo.fileptr != nil {
		lo.fileptr.Close()
	}
}

func (lo SimpleChannel) Open() {

	if len(lo.filename) > 0 {
		f, err := os.OpenFile(lo.filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)

		// check error
		if err != nil {
			fmt.Printf("Err %v for %s\n", err, lo.filename)
			panic(err)
		}

		logger := kitlog.NewLogfmtLogger(f)                                                        //(f, session.ID()+" ", log.LstdFlags)
		logger = kitlog.With(logger, "session_id", lo.sessionid, "ts", kitlog.DefaultTimestampUTC) //, "caller", kitlog.DefaultCaller)

		// check log is valid
		if logger == nil {
			panic("logger is nil")
		}

		lo.SetLog(logger)
		lo.fileptr = f
	} else {
		panic("")
	}

}
