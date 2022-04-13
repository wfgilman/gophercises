package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

type Problem struct {
	question string
	answer   string
}

func main() {
	// 0. User specifies filename and time limit
	fileName := flag.String("f", "problems.csv", "name of the file containing problems")
	timerDurationSec := flag.Int("t", 2, "duration of the quiz in seconds")
	shuffle := flag.Bool("s", false, "shuffle problems")
	flag.Parse()

	// 1. Read CSV file into array
	records := ImportProblems(*fileName)
	problems := ParseRecords(records)

	if *shuffle {
		rand.Seed(time.Now().UnixNano())
		rand.Shuffle(len(problems), func(i int, j int) {
			problems[i], problems[j] = problems[j], problems[i]
		})
	}

	// 2. Initialize correct and incorrect variables
	var numCorrect int = 0

	// 3. Start timer
	timer := time.NewTimer(time.Duration(*timerDurationSec) * time.Second)

	// 4. Loop through array and prompt question, take input
	for index, problem := range problems {
		fmt.Printf("Problem %v: What is %s? ", index+1, problem.question)
		answerChannel := make(chan string)
		go func() {
			var answer string
			fmt.Scanf("%s", &answer)
			answerChannel <- answer
		}()

		select {
		case <-timer.C:
			{
				fmt.Printf("\nTime is up! Your score is %d of %d", numCorrect, len(problems))
				return
			}
		case answer := <-answerChannel:
			if answer == problem.answer {
				numCorrect++
			}
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
