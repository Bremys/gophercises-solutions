package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"strings"

	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:    "csv",
				Aliases: []string{"c"},
				Usage:   "Load csv from `FILE`",
				Value:   "problems.csv",
			},
		},
		Action: mainAction,
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}

}

func mainAction(c *cli.Context) error {
	csvPath := c.String("csv")
	fmt.Println(csvPath)

	csvFile, err := os.Open(csvPath)
	if err != nil {
		log.Fatalln("Couldn't open the csv file", err)
	}

	content := csv.NewReader(csvFile)
	userInput := bufio.NewReader(os.Stdin)
	correctAmount, size := 0, 0

	for {
		// Read each record from csv
		record, err := content.Read()
		if err == io.EOF {
			break
		}

		if err != nil {
			log.Fatal(err)
		}

		if HandleSingleQuizLine(record, userInput) {
			correctAmount++
		}
		size++
	}

	fmt.Printf("Got %d correct out of %d\n", correctAmount, size)
	return nil
}

func HandleSingleQuizLine(record []string, userInput *bufio.Reader) bool {
	fmt.Printf("%s=", record[0])
	answer, _ := userInput.ReadString('\n')
	answer = strings.TrimSpace(answer)
	return answer == record[1]
}
