package main

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

func randate() (int, int, int) {
	rand.Seed(time.Now().Unix())
	minY := 1971
	maxY := 2070
	Y := rand.Intn(maxY-minY) + minY
	minM := 1
	maxM := 12
	M := rand.Intn(maxM-minM) + minM
	minD := 1
	maxD := 30
	D := rand.Intn(maxD-minD) + minD
	return Y, M, D
}

func sex() int {
	rand.Seed(time.Now().Unix())
	S := rand.Intn(8999) + 1000
	return S
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
	valuesText := []string{}

	for i := range PESEL {
		number := PESEL[i]
		text := strconv.Itoa(number)
		valuesText = append(valuesText, text)
	}

	result := strings.Join(valuesText, "")
	fmt.Printf("PESEL: %s\n", result)

	return result
}

type Person struct {
	IMIEPIERWSZE string
}

func chosname(filePath string) string {
	var persons []Person
	rand.Seed(time.Now().Unix())
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

		persons = append(persons, Person{
			IMIEPIERWSZE: record[headerMap["IMIEPIERWSZE"]],
		})
	}

	return persons[rName].IMIEPIERWSZE
}

func checkError(message string, err error) {
	// Error Logging
	if err != nil {
		log.Fatal(message, err)
	}
}

func main() {
	Y, M, D := randate()
	fmt.Printf("Date of birth: %d-%d-%d\n", Y, M, D)
	S := sex()
	if S%2 == 0 {
		fmt.Println("Sex: WOMAN")
		fmt.Printf("First name: %s\n", chosname("8_-_WYKAZ_IMION_ŻEŃSKICH_OSÓB_ŻYJĄCYCH_WG_POLA_IMIĘ_PIERWSZE_WYSTĘPUJĄCYCH_W_REJESTRZE_PESEL_BEZ_ZGONÓW.csv"))
		fmt.Printf("Last name: %s\n", chosname("nazwiska_żeńskie-osoby_żyjące_soxLKbB.csv"))
	} else {
		fmt.Println("Sex: MAN")
		fmt.Printf("First name: %s\n", chosname("8_-_WYKAZ_IMION_MĘSKICH_OSÓB_ŻYJĄCYCH_WG_POLA_IMIĘ_PIERWSZE_WYSTĘPUJĄCYCH_W_REJESTRZE_PESEL_BEZ_ZGONÓW.csv"))
		fmt.Printf("Last name: %s\n", chosname("NAZWISKA_MĘSKIE-OSOBY_ŻYJĄCE_oAcmDus.csv"))

	}
	PESEL(Y, M, D, S)
}
