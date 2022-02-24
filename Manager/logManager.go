package Manager

import (
	"log"
	"os"
)

type LogManager struct {
	InfoLogger  *log.Logger
	ErrorLogger *log.Logger
}

func (logManager *LogManager) InitLog() {
	file, err := os.OpenFile("logs.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}
	logManager.InfoLogger = log.New(file, "INFO :", log.Ldate|log.Ltime|log.Lshortfile)
	logManager.ErrorLogger = log.New(file, "ERROR :", log.Ldate|log.Ltime|log.Lshortfile)
}

func (logManager *LogManager) WriteInfoLog(message string) {
	logManager.InfoLogger.Println(message)
}

func (logManager *LogManager) WriteErrorLog(message string) {
	logManager.ErrorLogger.Println(message)
}
