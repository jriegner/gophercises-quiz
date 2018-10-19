package main

import (
	"encoding/csv"
	"fmt"
	"os"
)

var config configuration

type problem struct {
	question string
	answer   string
}

func main() {
	config.init()
	config.parse()

	f, err := os.Open(config.csv)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	lines, err := csv.NewReader(f).ReadAll()
	if err != nil {
		panic(err)
	}

	var problems []problem
	for _, line := range lines {
		p := problem{
			question: line[0],
			answer:   line[1],
		}
		problems = append(problems, p)
	}
	var right, total int
	for i, p := range problems {
		total++
		fmt.Printf("%d. Question: %s?", i+1, p.question)
		var a string
		_, err := fmt.Scanln(&a)
		if err != nil {
			panic(err)
		}

		if a == p.answer {
			right++
		}
	}

	fmt.Printf("You answered right in %d of %d questions.", right, total)
}
