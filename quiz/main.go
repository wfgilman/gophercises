package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
	"time"
)

type Problem struct {
	question string
	answer   string
}

func main() {
	// 0. User specifies filename
	fileName := flag.String("f", "problems.csv", "name of the file containing problems")
	flag.Parse()

	// 1. Read CSV file into array
	records := ImportProblems(*fileName)
	problems := ParseRecords(records)

	// 2. Initialize correct and incorrect variables
	var numCorrect int = 0
	var currentAnswer string

	// 3. Loop through array and prompt question, take input
	for index, problem := range problems {
		fmt.Printf("Problem %v: What is %s? ", index+1, problem.question)

		_, err := fmt.Scanf("%s", &currentAnswer)
		if err != nil {
			log.Fatal(err)
		}

		if currentAnswer == problem.answer {
			numCorrect++
		}
	}

	// 4. Return score
	fmt.Printf("Your score is %d of %d", numCorrect, len(problems))
}

func ImportProblems(fileName string) [][]string {
	f, err := os.Open(fileName)
	if err != nil {
		fmt.Printf("Could not open file named %s", fileName)
		os.Exit(1)
	}

	r := csv.NewReader(f)

	records, err := r.ReadAll()
	if err != nil {
		fmt.Printf("Could not read file named %s. Err: %v", fileName, err)
		os.Exit(1)
	}

	return records
}

func ParseRecords(records [][]string) []Problem {
	problems := make([]Problem, len(records))
	for index, record := range records {
		problems[index] = Problem{
			question: record[0],
			answer:   strings.TrimSpace(record[1]),
		}
	}
	return problems
}

func timer() {
	time.Sleep(30 * time.Millisecond)
}
