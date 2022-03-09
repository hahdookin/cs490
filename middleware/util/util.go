package util

import (
	"log"
)

// Check ->
func Check(err error) {
	if err != nil {
		log.Fatal(err.Error())
	}
}
