package main

import (
	"encoding/csv"
	"fmt"
	"os"
)

func main() {
	file, err := os.Open("quiz/problems.csv")
	if err != nil {
		newerr := fmt.Errorf("shit hit the fan: %v", err)
		panic(newerr)
	}
	csvReader := csv.NewReader(file)

	data, err := csvReader.ReadAll()

	if err != nil {
		newerr := fmt.Errorf("error while reading csv: %v", err)
		panic(newerr)
	}

	for _, line := range data {
		fmt.Printf("How much is %s ?\n", line[0])
		var input int
		fmt.Scan(&input)
	}
}
