package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strings"
	"time"

	"github.com/urfave/cli/v2"
)

func Quiz(c *cli.Context) error {
	csvPath := c.String("csv")
	shuffle := c.Bool("shuffle")

	csvFile, err := os.Open(csvPath)
	if err != nil {
		log.Fatalln("Couldn't open the csv file", err)
	}

	content := csv.NewReader(csvFile)
	userInput := bufio.NewReader(os.Stdin)

	records, err := content.ReadAll()

	if err != nil {
		log.Fatal(err)
	}

	if shuffle {
		shuffleRecords(records)
	}

	correctAmount := 0

	for _, record := range records {
		fmt.Printf("%s=", record[0])
		if equalAnswers(record[1], userInput) {
			correctAmount++
		}
	}

	fmt.Printf("Got %d correct out of %d\n", correctAmount, len(records))
	return nil
}

func shuffleRecords(records [][]string) {
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(records), func(i, j int) { records[i], records[j] = records[j], records[i] })
}

func equalAnswers(csvAnswer string, userInput *bufio.Reader) bool {
	userAnswer, _ := userInput.ReadString('\n')
	userAnswer = strings.ToLower(strings.TrimSpace(userAnswer))
	correctAnswer := strings.ToLower(strings.TrimSpace(csvAnswer))
	return userAnswer == correctAnswer
}
