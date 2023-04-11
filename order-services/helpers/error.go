package helpers

import "log"

func FailError(err error, msg string) {
	if err != nil {
		log.Panicf("%s: %s", msg, err)
	}
}
