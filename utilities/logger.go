package utilities

import (
	"log"
	"os"
)

var Logger *log.Logger

func InitLogger() {
	Logger = log.New(os.Stdout, "[Kahawatein Log] ", log.LstdFlags)
}
