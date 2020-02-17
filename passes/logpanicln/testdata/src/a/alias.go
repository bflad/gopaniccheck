package a

import (
	l "log"
)

func falias() {
	l.Panicln("test") // want "avoid log.Panicln\\(\\) usage"
}
