package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	csvFileName := flag.String("csv", "problems.csv", "filename of csv file containing question separated by ans")

	totalDuration := flag.Int("time", 30, "Total Time Duration of the Quiz")

	flag.Parse()

	file, err := os.Open(*csvFileName) //csvFileName returns pointer to the string

	if err != nil {
		exit(fmt.Sprintf("Failed to Open the file name: %s", *csvFileName))
	}

	r := csv.NewReader(file)
	lines, err := r.ReadAll()

	if err != nil {
		exit("Failed to parse the csv file")
	}
	problems := parseLines(lines)

	score := 0

	timer := time.NewTimer(time.Duration(*totalDuration) * time.Second)

	for i, v := range problems {
		fmt.Printf("Problem #%v:%v = \n", i+1, v.q)
		answerCh := make(chan string)
		go func() {
			var answer string
			fmt.Scanf("%s", &answer)
			answerCh <- answer
		}()

		select {
		case <-timer.C: //timer and answer on different go routines and don't interfere with each other
			fmt.Printf("You scored %v out of %v", score, len(problems))
			return
		case answer := <-answerCh: //blocks until one of the case can run if we used default then case would get stuck on default
			if answer == v.a {
				score++
			}
		}

	}
	fmt.Printf("You scored %v out of %v", score, len(problems))
}

type problem struct {
	q string
	a string
}

func parseLines(lines [][]string) []problem {
	problems := make([]problem, len(lines))

	for i, line := range lines {
		problems[i] = problem{
			q: line[0],
			a: strings.TrimSpace(line[1]), //trims whitespace if exists in answer
		}
	}

	return problems
}

func exit(msg string) {
	fmt.Println(msg)
	os.Exit(1)
}
