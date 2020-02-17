package a

import (
	l "log"
)

func falias() {
	l.Panic("test") // want "avoid log.Panic\\(\\) usage"
}
