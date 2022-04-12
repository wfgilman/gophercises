package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"log"
	"os"
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
	problems := ImportProblems(*fileName)

	// 2. Initialize correct and incorrect variables
	var numCorrect int = 0
	var numProblems int = len(problems)
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
	fmt.Printf("Your score is %d of %d", numCorrect, numProblems)
}

func ImportProblems(fileName string) []Problem {
	var problems []Problem

	f, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}

	r := csv.NewReader(f)

	records, err := r.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	for _, record := range records {
		problem := Problem{
			question: record[0],
			answer:   record[1],
		}

		problems = append(problems, problem)
	}

	return problems
}
