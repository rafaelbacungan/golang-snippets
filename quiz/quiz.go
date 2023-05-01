package quiz

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"strings"
	"time"
)

func Quiz() {
	csvFilename := flag.String("csv", "problems.csv", "a csv file in the format of"+
		"'question,answer'")
	flag.Parse()

	timeLimit := flag.Int("limit", 30, "the time limit for the quiz in seconds")
	flag.Parse()

	file, err := os.Open(*csvFilename)
	if err != nil {
		exit(fmt.Sprintf("Failed to open the CSV file: %s\n", *csvFilename))
	}
	r := csv.NewReader(file)
	lines, err := r.ReadAll()
	if err != nil {
		exit("Failed to parse the provided CSV File.")
	}

	problems := parseLines(lines)
	fmt.Println(problems)

	timer := time.NewTimer(time.Duration(*timeLimit) * time.Second)
	// <-timer.C

	correct := 0
	for i, p := range problems {
		select {
		case <-timer.C:
			fmt.Printf("Problem #%d: %s = \n", i+1, p.q)
			return
		default:
			var answer string
			fmt.Scanf("%s\n", &answer)
			if answer == p.a {
				fmt.Println("Correct!")
				correct++
			}
		}

	}

	fmt.Printf("You scored %d out of %d.\n", correct, len(problems))
}

func parseLines(lines [][]string) []problem {
	ret := make([]problem, len(lines))
	for i, line := range lines {
		ret[i] = problem{
			q: line[0],
			a: strings.TrimSpace(line[1]),
		}
	}
	return ret
}

type problem struct {
	q string
	a string
}

func exit(msg string) {
	fmt.Println(msg)
	os.Exit(1)
}
