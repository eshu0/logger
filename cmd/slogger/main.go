package main

import (
	"bufio"
	"flag"
	"os"
	"strings"

	sl "github.com/eshu0/simplelogger"
)

func main() {

	// this is the dummy logger object
	//&sl.SimpleLogger{}
	filename := flag.String("filename", "slogger.log", "Filename out - defaults to slogger.log")
	session := flag.String("session", "123", "Session - defaults to 123")

	flag.Parse()

	log := sl.NewSimpleLoggerWithFilename(*filename, *sessionid)

	// lets open a flie log using the session
	f1 := log.OpenFileLog()

	reader := bufio.NewReader(os.Stdin)
	for {
		// read the string input
		text, readerr := reader.ReadString('\n')

		if readerr != nil {
			log.LogDebugf("main()", "Reading input has provided following err '%s'", readerr.Error())
			break
			// break out for loop
		}

		// convert CRLF to LF
		text = strings.Replace(text, "\n", "", -1)
		if text == "quit" {
			log.LogDebugf("main()", "Quitting do input '%s'", text)
		} else {
			log.LogDebugf("main()", "'%s'", text)
		}
	}

	//defer the close till the shell has closed
	defer f1.Close()
}
