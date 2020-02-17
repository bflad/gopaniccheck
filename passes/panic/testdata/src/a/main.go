package a

func f() {
	panic(true) // want "avoid panic\\(\\) usage"
}
