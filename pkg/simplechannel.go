package simplelogger

import (
	"errors"
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

//func (sc SimpleChannel) SetFileName(filename string) {
//	sc.filename = filename
//}

//func (sc SimpleChannel) GetFileName() string {
//	return sc.filename
//}

func (sc SimpleChannel) SetSessionID(sessionid string) {
	sc.sessionid = sessionid
}

func (sc SimpleChannel) GetSessionID() string {
	return sc.sessionid
}

func (sc SimpleChannel) SetLogLevel(lvl kitlevel.Option) {
	sc.level = lvl
	sc.log = kitlevel.NewFilter(sc.log, lvl)
}

func (sc SimpleChannel) GetLogLevel() kitlevel.Option {
	return sc.level
}

func (sc SimpleChannel) SetLog(log kitlog.Logger) {
	sc.log = log
}

func (sc SimpleChannel) GetLog() kitlog.Logger {
	return sc.log
}

func (sc SimpleChannel) Close() error {
	if sc.fileptr != nil {
		return sc.fileptr.Close()
	}
	return nil
}

func (sc SimpleChannel) Open() error {

	if len(sc.filename) > 0 {
		f, err := os.OpenFile(sc.filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)

		// check error
		if err != nil {
			return err
		}

		logger := kitlog.NewLogfmtLogger(f)                                                        //(f, session.ID()+" ", log.LstdFlags)
		logger = kitlog.With(logger, "session_id", sc.sessionid, "ts", kitlog.DefaultTimestampUTC) //, "caller", kitlog.DefaultCaller)

		// check log is valid
		if logger == nil {
			return errors.New("logger is nil")
		}

		sc.SetLog(logger)
		sc.fileptr = f
		return nil

	} else {
		return errors.New("filename was missing from the channel")
	}

}

func (sc SimpleChannel) GetDetails() string {
	return fmt.Sprintf("%s : %s", sc.sessionid, sc.filename)
}
