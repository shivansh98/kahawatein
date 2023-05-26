package utilities

import "log"

func Panic(err error) {
	log.Default().Printf("got an error %v", err)
	Panic(err)
}

