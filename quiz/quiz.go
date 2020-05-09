package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"math/rand"
	"os"
	"time"
)

// QuizArgs represents all the information needed in order to correctly run
// our quiz.
type QuizArgs struct {
	csvPath   string
	shuffle   bool
	timeLimit time.Duration
	timeout   bool
}

// Quiz runs a quiz described in a csv file, where each record represents
// a question and an answer. It prints the amount of correct answers by the
// user.
func Quiz(args *QuizArgs) error {
	csvFile, err := os.Open(args.csvPath)
	defer csvFile.Close()

	if err != nil {
		log.Fatalln("Couldn't open the csv file", err)
	}

	content := csv.NewReader(csvFile)
	records, err := content.ReadAll()

	if err != nil {
		log.Fatal("Couldn't read csv file", err)
	}

	if args.shuffle {
		shuffleRecords(records)
	}

	correctAmount := 0
	if args.timeout {
		correctAmount = RecordsIterateTimeout(records, args.timeLimit)
	} else {
		correctAmount = RecordsIterate(records)
	}

	fmt.Printf("Score: %d/%d\n", correctAmount, len(records))
	return nil
}

func shuffleRecords(records [][]string) {
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(records), func(i, j int) { records[i], records[j] = records[j], records[i] })
}
