package funcs

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

//Addresses to men and women names databases
var womenFN = "Databases/8_-_WYKAZ_IMION_ŻEŃSKICH_OSÓB_ŻYJĄCYCH_WG_POLA_IMIĘ_PIERWSZE_WYSTĘPUJĄCYCH_W_REJESTRZE_PESEL_BEZ_ZGONÓW.csv"
var womenLN = "Databases/nazwiska_żeńskie-osoby_żyjące_soxLKbB.csv"
var menFN = "Databases/8_-_WYKAZ_IMION_MĘSKICH_OSÓB_ŻYJĄCYCH_WG_POLA_IMIĘ_PIERWSZE_WYSTĘPUJĄCYCH_W_REJESTRZE_PESEL_BEZ_ZGONÓW.csv"
var menLN = "Databases/NAZWISKA_MĘSKIE-OSOBY_ŻYJĄCE_oAcmDus.csv"

type Persons struct {
	IMIEPIERWSZE string
}

func Randate() (int, int, int) {
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

	return Y, M, D
}

func RandateToDOB(Y int, M int, D int) string {
	DOB := fmt.Sprintf("%d-%d-%d", Y, M, D)

	return DOB
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
	rName := rand.Intn(100)
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
