package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"

	sl "github.com/eshu0/logger/pkg"
	sli "github.com/eshu0/logger/pkg/interfaces"
)

func main() {

	// this is the dummy logger object
	//&sl.SimpleLogger{}
	filename := flag.String("filename", "slogger.log", "Filename out - defaults to slogger.log")
	session := flag.String("session", "123", "Session - defaults to 123")

	flag.Parse()

	log := sl.NewSimpleLogger(*filename, *session)

	// lets open a flie log using the session
	log.OpenAllChannels()

	//defer the close till the shell has closed
	defer log.CloseAllChannels()

	reader := bufio.NewReader(os.Stdin)

	for {
		// read the string input
		text, readerr := reader.ReadString('\n')

		if readerr != nil {
			log.LogDebugf("main()", "Reading input has provided following err '%s'", readerr.Error())
			break
			// break out for loop
		}

		log.LogDebugf("main()", "input was: '%s'", text)

		// convert CRLF to LF
		text = strings.Replace(text, "\n", "", -1)
		if strings.ToLower(text) == "quit" || strings.ToLower(text) == "exit" {
			fmt.Println("bye bye")
			break
		} else {

			inputs := strings.Split(text, " ")

			if len(inputs) == 1 {

				if strings.ToLower(inputs[0]) == "sessionids" {
					sessionids := log.GetSessionIDs()
					for _, SessionID := range sessionids {
						fmt.Printf("'%s'", SessionID)
						log.LogInfof("main()", "Get Session ID: '%s'", SessionID)
					}
				} else if strings.ToLower(inputs[0]) == "sessions" {
					for _, channel := range log.GetChannels() {
						fmt.Printf("Session ID: '%s'", channel.GetSessionID())
						fmt.Printf("FileName: '%s'", channel.GetFileName())
						fmt.Println("----")
						log.LogInfof("main()", "Get Session: '%s'", channel.GetSessionID())
						log.LogInfof("main()", "Get FileName: '%s'", channel.GetFileName())

					}
				} else if strings.ToLower(inputs[0]) == "printscreenon" {
					fmt.Println("Setting Printing to screen on")
					log.SetPrintToScreen(sli.PrintInfo)
				} else if strings.ToLower(inputs[0]) == "printscreenoff" {
					fmt.Println("Setting Printing to screen off")
					log.SetPrintToScreen(sli.PrintNone)
				} else if strings.ToLower(inputs[0]) == "printstatus" {

					if log.GetPrintToScreen() == sli.PrintInfo {
						fmt.Println("Printing to screen")
					} else if log.GetPrintToScreen() == sli.PrintDebug {
						fmt.Println("Printing to screen with debug")
					} else {
						fmt.Println("Not Printing to screen")
					}
				}

			} else {

				if len(inputs) >= 2 {
					fmt.Printf("Logged to '%s' with %s", inputs[0], inputs[1])
					if strings.ToLower(inputs[0]) == "debug" {
						log.LogDebugf("main()", "'%s'", inputs[1])
					} else if strings.ToLower(inputs[0]) == "info" {
						log.LogInfof("main()", "'%s'", inputs[1])
					} else if strings.ToLower(inputs[0]) == "error" {
						log.LogErrorf("main()", "'%s'", inputs[1])
					} else if strings.ToLower(inputs[0]) == "warn" {
						log.LogWarnf("main()", "'%s'", inputs[1])
					} else if strings.ToLower(inputs[0]) == "add" && strings.ToLower(inputs[1]) == "session" {
						channel := sl.SimpleChannel{}
						channel.SetSessionID(inputs[2])
						channel.SetFileName(inputs[3])
						channel.Open()
						log.AddChannel(channel)
						//logger := kitlog.NewLogfmtLogger(f1)
						//logger = kitlog.With(logger, "session_id", inputs[2], "ts", kitlog.DefaultTimestampUTC)
						//log.AddLog(logger)
					}
				} else {
					fmt.Printf("'%s' was split but only had %d inputs", text, len(inputs))
					log.LogDebugf("main()", "'%s' was split but only had %d inputs", text, len(inputs))
				}
			}

		}
	}

}
