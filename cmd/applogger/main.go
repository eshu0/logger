package main

import (
	sl "github.com/eshu0/logger/pkg"
)

type TestApp struct {
	sl.AppLogger
}

func main() {
	ta := TestApp{}
	ta.LogInfo("Logging Info!")
	ta.LogError("Logging Error!")
	ta.LogDebug("Logging Debug!")
	ta.LogWarn("Logging LogWarn!")
	if ta.Log != nil {
		ta.Log.PrintDetails()
	}
	ta.FinishLogging()
}
