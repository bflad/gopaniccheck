package a

import (
	l "log"
)

func falias() {
	l.Panicf("%s", "test") // want "avoid log.Panicf\\(\\) usage"
}
