package utils

import (
	"io"
	"log"
	"os"
)

func LoggingSettings(logFile string) {
<<<<<<< HEAD
	logfile, err := os.OpenFile(logFile, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
=======
	logfile, err := os.OpenFile(logFile,os.O_RDWR|os.O_CREATE|os.O_APPEND,0666)
>>>>>>> origin/main
	if err != nil {
		log.Fatalln(err)
	}
	multiLogFile := io.MultiWriter(os.Stdout, logfile)
	log.SetFlags(log.Ldate|log.Ltime|log.Lshortfile)
	log.SetOutput(multiLogFile)

	log.Println("test")
}