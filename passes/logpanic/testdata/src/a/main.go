package a

import (
	"log"
)

func f() {
	log.Panic("test") // want "avoid log.Panic\\(\\) usage"
}
