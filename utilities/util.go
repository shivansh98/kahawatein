package utilities

import "log"

func CallPanic(err error) {
	log.Default().Printf("got an error %v", err)
	panic(err)
}
