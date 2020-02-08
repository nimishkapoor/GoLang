package main

import (
	"encoding/csv"
	"log"
	"os"
	"strconv"
)

func main() {
	rows := readExcel("test.csv")
	rows = calculate(rows)
	writeExcel("new.csv", rows)
}

func readExcel(name string) [][]string {

	f, err := os.Open(name)
	if err != nil {
		log.Fatalf("Cannot open '%s' : %s \n", name, err.Error())
	}

	defer f.Close()

	r := csv.NewReader(f)

	rows, err := r.ReadAll()

	if err != nil {
		log.Fatalln("Cannot read CSV data:", err.Error())
	}

	return rows
}

func calculate(rows [][]string) [][]string {
	sum := 0

	for i := range rows {
		if i == 0 {
			rows[0] = append(rows[0], "Total")
			continue
		}
		item := rows[i][2]

		price, err := strconv.Atoi(rows[i][1])
		if err != nil {
			log.Fatalf("Cannot retrieve price of %s: %s\n", item, err)
		}

		qty, err := strconv.Atoi(rows[i][2])
		if err != nil {
			log.Fatalf("Cannot retrieve quantity of %s: %s\n", item, err)
		}

		total := price * qty

		rows[i] = append(rows[i], strconv.Itoa(total))

		sum += total
	}

	return rows
}

func writeExcel(name string, rows [][]string) {
	f, err := os.Create(name)
	if err != nil {
		log.Fatalf("Cannot open '%s': %s\n", name, err.Error())
	}

	defer func() {
		e := f.Close()
		if e != nil {
			log.Fatalf("Cannot close '%s' : %s \n", name, e.Error())
		}
	}()

	w := csv.NewWriter(f)
	err = w.WriteAll(rows)
}
