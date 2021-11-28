package simplelogger

import (
	"testing"
)

func TestNewSimpleLoggerOpenClose(t *testing.T) {

	log := NewSimpleLogger("slogger.log", "123")
	log.PrintDetails()

	// lets open a file log using the session
	openerrors := log.OpenAllChannels()

	if len(openerrors) > 0 {

		t.Errorf("Open failed with %d errors", len(openerrors))
		for i := 0; i < len(openerrors); i++ {
			t.Error(openerrors[i])
		}
	}

	closeerrors := log.CloseAllChannels()

	if len(closeerrors) > 0 {

		t.Errorf("Close failed with %d errors", len(closeerrors))
		for i := 0; i < len(closeerrors); i++ {
			t.Error(closeerrors[i])
		}
	}

}
