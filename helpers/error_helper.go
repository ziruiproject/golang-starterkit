package helpers

import (
	"log"
)

func PanicOnError(err error) {
	if err != nil {
		log.Println(err)
		panic(err)
	}
}
