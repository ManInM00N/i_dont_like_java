package binary

import (
	"log"
	"os"
)

var (
	DebugLog *log.Logger
	InfoLog  *log.Logger
	F        *os.File
)

func LogInit() {
	log.SetFlags(log.Ltime)
	F, _ = os.OpenFile("temp.log", os.O_CREATE|os.O_APPEND|os.O_RDWR, os.ModePerm)
	log.SetOutput(F)
	InfoLog = log.New(F, "[Info]  - ", log.Ldate|log.Ltime)
	DebugLog = log.New(F, "[Debug] - ", log.Ldate|log.Ltime)
}
