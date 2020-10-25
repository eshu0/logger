package main

import (
	sl "github.com/eshu0/simplelogger/pkg"
)

func main() {

	// Open a log
	slog := sl.NewApplicationLogger()

	// lets open a flie log using the session
	slog.OpenAllChannels()

	//defer the close till the shell has closed
	defer slog.CloseAllChannels()

	slog.LogInfo("Logging Info!")
	slog.LogError("Logging Error!")
	slog.LogDebug("Logging Debug!")
	slog.LogDebug("Logging LogWarn!")
}
