package utils

import (
	"os"
	"log"
)

func LoggingSettings(LogFile string) {
	logfile, err := os.OpenFile(LogFile, os.O_RDWR | os.O_CREATE, 0666)
	if err != nil {
		log.Fatalln(err)
	}
}