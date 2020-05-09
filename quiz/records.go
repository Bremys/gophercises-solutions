package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

// RecordsIterate receives an array of records, listens to user input
// and returns the number of correct answers by the user.
func RecordsIterate(records [][]string) int {
	correctAmount := 0
	userInput := bufio.NewReader(os.Stdin)

	for _, record := range records {
		fmt.Printf("%s=", record[0])
		if equalAnswers(record[1], userInput) {
			correctAmount++
		}
	}
	return correctAmount
}

func RecordsIterateTimeout(records [][]string, timeout time.Duration) int {
	correctAmount := 0
	userInput := bufio.NewReader(os.Stdin)

	for _, record := range records {
		fmt.Printf("%s=", record[0])

		ch := make(chan bool, 1)

		go func(ch chan bool, answer string, userInput *bufio.Reader) {
			ch <- equalAnswers(answer, userInput)
		}(ch, record[1], userInput)

		select {
		case same := <-ch:
			if same {
				correctAmount++
			}

		case <-time.After(timeout):
			fmt.Println()
			return correctAmount
		}
	}
	return correctAmount
}

func equalAnswers(csvAnswer string, userInput *bufio.Reader) bool {
	userAnswer, _ := userInput.ReadString('\n')
	userAnswer = strings.ToLower(strings.TrimSpace(userAnswer))
	correctAnswer := strings.ToLower(strings.TrimSpace(csvAnswer))
	return userAnswer == correctAnswer
}
