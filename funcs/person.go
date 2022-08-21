package funcs

import (
	"fmt"
	"math/rand"
	"reflect"
)

type Person struct {
	FirstName   string `json:"full_name"`
	LastName    string `json:"last_name"`
	DateOfBirth string `json:"date_of_birth"`
	Sex         string `json:"sex"`
	PESEL       string `json:"pesel"`
}

func (p Person) Introduce() {
	s := reflect.ValueOf(&p).Elem()
	typeOfT := s.Type()
	for i := 0; i < s.NumField(); i++ {
		f := s.Field(i)
		fmt.Printf("%s: %v\n",
			typeOfT.Field(i).Name, f.Interface())
	}
}
func (p *Person) Generate() {
	Y, M, D := Randate()
	S := rand.Intn(8999) + 1000
	p.DateOfBirth = RandateToDOB(Y, M, D)
	p.PESEL = PESEL(Y, M, D, S)
	if S%2 == 0 {
		p.Sex = "Woman"
		p.FirstName = Chosname(womenFN)
		p.LastName = Chosname(womenLN)

	} else {
		p.Sex = "Man"
		p.FirstName = Chosname(menFN)
		p.LastName = Chosname(menLN)

	}
}
