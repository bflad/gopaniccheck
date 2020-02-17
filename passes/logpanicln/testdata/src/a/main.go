package a

import (
	"log"
)

func f() {
	log.Panicln("test") // want "avoid log.Panicln\\(\\) usage"
}
