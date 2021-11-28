package slinterfaces

import (
	kitlog "github.com/go-kit/log"
	kitlevel "github.com/go-kit/log/level"
)

// main interface for the Simple Channel
type ISimpleChannel interface {
	GetLog() kitlog.Logger
	SetLog(log kitlog.Logger)

	GetFileName() string
	SetFileName(sessionid string)

	GetSessionID() string
	SetSessionID(sessionid string)

	//log functions
	GetLogLevel() kitlevel.Option
	SetLogLevel(kitlevel.Option)

	Open() error
	Close() error

	GetDetails() string
}
