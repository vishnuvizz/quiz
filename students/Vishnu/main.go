package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"log"
	"os"
)

func main() {
	fmt.Println("Hello Hello! Welcome to the quiz game")

	fileName := flag.String("fileName", "problems.csv", "a file name in csv format")
	flag.Parse()

	timeout = flag.Int("timeout", 30, "time limit for the quiz, if it exceeds for a question, it exits program")
	flag.Parse()

	records := readcsvfiles(*fileName)

	var marks int

	for _, value := range records {

		fmt.Printf("Question: %s = ", value[0])
		var userAnswer string
		fmt.Scanln(&userAnswer)
		if userAnswer == value[1] {
			marks++
		}

	}
	fmt.Println("You scored", marks, "out of", len(records))

}

func readcsvfiles(filePath string) [][]string {
	f, err := os.Open(filePath)
	if err != nil {
		log.Fatal("Unable to open file: ", err)
	}
	defer f.Close()

	csvReader := csv.NewReader(f)
	records, err := csvReader.ReadAll()
	if err != nil {
		log.Fatal("Unable to parse file as csv for " + filePath)
	}
	return records
}
