package main

import (
	"IDGen/funcs"
)

func main() {
	var Identity funcs.Person
	Identity.Generate()
	Identity.Introduce()
}
