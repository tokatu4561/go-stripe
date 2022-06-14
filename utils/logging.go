package utils

import (
	"os"
	"log"
	"io"
)

func LoggingSettings(LogFile string) {
	logfile, err := os.OpenFile(LogFile, os.O_RDWR | os.O_CREATE, 0666)
	if err != nil {
		log.Fatalln(err)
	}

	multiLogFile := io.MultiWriter(os.Stdout, logfile)
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
	log.SetOutput(multiLogFile)
	
}