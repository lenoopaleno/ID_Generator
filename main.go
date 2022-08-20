package main

import (
	"IDGen/funcs"
)

func main() {
	var Identity funcs.Person
	var S int
	Y, M, D, DOB := funcs.Randate()
	Identity.DateOfBirth = DOB
	Identity, S = funcs.Sex(Identity)
	Identity.PESEL = funcs.PESEL(Y, M, D, S)
	funcs.Represent(Identity)
}
