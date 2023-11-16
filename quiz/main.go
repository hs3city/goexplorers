package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("problems.csv")
	if err != nil {
		fmt.Printf(os.Getwd())
		newerr := fmt.Errorf("shit hit the fan: %v", err)
		panic(newerr)
	}
	csvReader := csv.NewReader(file)

	data, err := csvReader.ReadAll()

	if err != nil {
		newerr := fmt.Errorf("error while reading csv: %v", err)
		panic(newerr)
	}

	totalAnswers := len(data)
	correctAnswers := 0

	for _, line := range data {
		question, answer := line[0], line[1]
		fmt.Printf("How much is %s ?\n", question)
		var input int
		fmt.Scan(&input)

		verify, err := strconv.Atoi(strings.TrimSpace(answer))

		if err != nil {
			newerr := fmt.Errorf("failed to convert to int: %v", err)
			panic(newerr)
		}

		if input == verify {
			correctAnswers++
		}
	}

	fmt.Printf("Total answers: %d\n", totalAnswers)
	fmt.Printf("Correct answers: %d\n", correctAnswers)
	fmt.Printf("Score: %f%%\n", float64(correctAnswers)/float64(totalAnswers)*100)
}
