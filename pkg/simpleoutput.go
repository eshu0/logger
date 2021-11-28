package simplelogger

import (
	"fmt"

	sli "github.com/eshu0/logger/pkg/interfaces"
)

// Simple Output represents an output to be logged to
type SimpleOutput struct {

	//inherit from interface
	sli.ISimpleOutput

	// session id
	sessionid string
}

/*
 Channel Functions after here
*/

func (so SimpleOutput) GetSessionID() string {
	return so.sessionid
}

func (so SimpleOutput) Close() error {
	return nil
}

func (so SimpleOutput) Open() (sli.ISimpleOutput, error) {
	return nil, nil
}

func (so SimpleOutput) GetDetails() string {
	return fmt.Sprintf("Output: %s", so.sessionid)
}
