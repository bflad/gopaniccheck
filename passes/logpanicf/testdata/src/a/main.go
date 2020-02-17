package a

import (
	"log"
)

func f() {
	log.Panicf("%s", "test") // want "avoid log.Panicf\\(\\) usage"
}
