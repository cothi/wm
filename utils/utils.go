package utils

import "log"

func UtilsError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
