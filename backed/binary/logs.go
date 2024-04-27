package binary

import (
	"log"
	"os"
)

var F *os.File

func LogInit() {
	log.SetFlags(log.Ldate | log.Ltime)
	F, _ = os.OpenFile("temp.log", os.O_CREATE|os.O_APPEND|os.O_RDWR, os.ModePerm)
	log.SetOutput(F)
}
