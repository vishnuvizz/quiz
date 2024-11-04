package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"log"
	"os"
	"time"
)

func main() {
	fmt.Println("Hello Hello! Welcome to the quiz game")

	fileName := flag.String("fileName", "problems.csv", "a file name in csv format")
	timeout := flag.Int("timeout", 30, "time limit for the quiz, if it exceeds for a question, it exits program")
	flag.Parse()

	records := readcsvfiles(*fileName)

	var marks int

	for _, value := range records {
		var userAnswer string
		fmt.Printf("Question: %s = ", value[0])
		timer := time.NewTimer(time.Duration(*timeout) * time.Second)

		answerCh := make(chan string)
		go func() {
			fmt.Scanln(&userAnswer)
			answerCh <- userAnswer
		}()

		select {
		case <-timer.C:
			fmt.Println("Time out!")
			fmt.Println("You scored", marks, "out of", len(records))
			return
		case userAnswer := <-answerCh:
			if userAnswer == value[1] {
				fmt.Scanln(&userAnswer)
				marks++
			}
			timer.Stop()
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
