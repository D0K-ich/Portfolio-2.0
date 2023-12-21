package logger

import (
	"github.com/gofiber/fiber/v2/log"
	"os"
)

func Init() {
	f, err := os.OpenFile("../../logs/Dok's.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatal(err)
	}
	log.SetOutput(f)
	log.Trace("Logger loaded")
}
