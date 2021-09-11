package logger

import (
	"log"
	"os"
)

type LoggerInstance struct {
	InfoLogger    *log.Logger
	WarningLogger *log.Logger
	ErrorLogger   *log.Logger
}

func NewLogger(FileName string) *LoggerInstance {

	file, err := os.OpenFile(FileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}
	InfoLogger := log.New(file, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
	WarningLogger := log.New(file, "WARNING: ", log.Ldate|log.Ltime|log.Lshortfile)
	ErrorLogger := log.New(file, "Error: ", log.Ldate|log.Ltime|log.Lshortfile)
	return &LoggerInstance{InfoLogger: InfoLogger, WarningLogger: WarningLogger, ErrorLogger: ErrorLogger}

}
