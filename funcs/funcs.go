package funcs

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"math/rand"
	"os"
	"reflect"
	"strconv"
	"strings"
	"time"
)

//Addresses to men and women names databases
var womenFN = "Databases/8_-_WYKAZ_IMION_ŻEŃSKICH_OSÓB_ŻYJĄCYCH_WG_POLA_IMIĘ_PIERWSZE_WYSTĘPUJĄCYCH_W_REJESTRZE_PESEL_BEZ_ZGONÓW.csv"
var womenLN = "Databases/nazwiska_żeńskie-osoby_żyjące_soxLKbB.csv"
var menFN = "Databases/8_-_WYKAZ_IMION_MĘSKICH_OSÓB_ŻYJĄCYCH_WG_POLA_IMIĘ_PIERWSZE_WYSTĘPUJĄCYCH_W_REJESTRZE_PESEL_BEZ_ZGONÓW.csv"
var menLN = "Databases/NAZWISKA_MĘSKIE-OSOBY_ŻYJĄCE_oAcmDus.csv"

type Person struct {
	FullName    string `json:"full_name"`
	LastName    string `json:"last_name"`
	DateOfBirth string `json:"date_of_birth"`
	Sex         string `json:"sex"`
	PESEL       string `json:"pesel"`
}
type Persons struct {
	IMIEPIERWSZE string
}

var person Person

func Randate() (int, int, int, string) {
	rand.Seed(time.Now().Unix())
	X := time.Now()
	minY := 1971
	maxY := X.Year()
	Y := rand.Intn(maxY-minY) + minY
	minM := 1
	maxM := 12
	M := rand.Intn(maxM-minM) + minM
	minD := 1
	maxD := 30
	D := rand.Intn(maxD-minD) + minD
	DOB := fmt.Sprintf("%d-%d-%d\n", Y, M, D)

	return Y, M, D, DOB
}

func Sex(identity Person) (Person, int) {
	rand.Seed(time.Now().Unix())
	S := rand.Intn(8999) + 1000
	if S%2 == 0 {
		identity.Sex = "Woman"
		identity.FullName = Chosname(womenFN)
		identity.LastName = Chosname(womenLN)

	} else {
		identity.Sex = "Man"
		identity.FullName = Chosname(menFN)
		identity.LastName = Chosname(menLN)

	}
	return identity, S
}

func PESEL(Y int, M int, D int, S int) string {
	var PESEL = [11]int{}
	var wage = [10]int{1, 3, 7, 9, 1, 3, 7, 9, 1, 3}
	PESEL[0] = (Y % 100) / 10
	PESEL[1] = Y % 10
	if M < 10 {
		PESEL[2] = 0
		PESEL[3] = M
	} else {
		PESEL[2] = (M % 100) / 10
		PESEL[3] = M % 10
	}
	if D < 10 {
		PESEL[4] = 0
		PESEL[5] = D
	} else {
		PESEL[4] = (D % 100) / 10
		PESEL[5] = D % 10
	}
	PESEL[6] = S / 1000
	PESEL[7] = S % 1000 / 100
	PESEL[8] = S % 100 / 10
	PESEL[9] = S % 10

	var c int
	for i := 0; i < len(wage); i++ {
		c += (PESEL[i] * wage[i]) % 10
	}
	PESEL[10] = c % 10
	var valuesText []string

	for i := range PESEL {
		number := PESEL[i]
		text := strconv.Itoa(number)
		valuesText = append(valuesText, text)
	}

	result := strings.Join(valuesText, "")
	person.PESEL = result

	return result
}

func checkError(message string, err error) {
	// Error Logging
	if err != nil {
		log.Fatal(message, err)
	}
}

func Chosname(filePath string) string {
	rand.Seed(time.Now().Unix())
	var persons []Persons
	rName := rand.Intn(200)
	isFirstRow := true
	headerMap := make(map[string]int)

	f, _ := os.Open(filePath)
	r := csv.NewReader(f)

	for {
		// Read row
		record, err := r.Read()
		// Stop at EOF.
		if err == io.EOF {
			break
		}
		checkError("Some other error occurred", err)

		if isFirstRow {
			isFirstRow = false
			for _, v := range record {
				headerMap[v] = 0
			}
			continue
		}

		persons = append(persons, Persons{
			IMIEPIERWSZE: record[headerMap[""]],
		})
	}

	return persons[rName].IMIEPIERWSZE
}

func Represent(person Person) {
	s := reflect.ValueOf(&person).Elem()
	typeOfT := s.Type()
	for i := 0; i < s.NumField(); i++ {
		f := s.Field(i)
		fmt.Printf("%s: %v\n",
			typeOfT.Field(i).Name, f.Interface())
	}
}
