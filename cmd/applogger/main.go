package main

import (
	sl "github.com/eshu0/simplelogger/pkg"
)

type TestApp struct {
	sl.AppLogger
}

func main() {
	ta := TestApp{}
	ta.Start()

	ta.LogInfo("Logging Info!")
	ta.LogError("Logging Error!")
	ta.LogDebug("Logging Debug!")
	ta.LogWarn("Logging LogWarn!")

	ta.Finish()
}
