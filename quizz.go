package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"log"
	"os"
)

type Problem struct {
	question string
	answer   string
}

func main() {
	fileName, err := os.Open("problems.csv")
	if err != nil {
		log.Fatalf("erreur :%v", err)
	}
	defer fileName.Close()

	reader := csv.NewReader(fileName)
	csvContent, err := reader.ReadAll()
	if err != nil {
		log.Fatalf("could not read the csv file : %v", err)
	}
	problems := make([]Problem, len(csvContent))
	for i, content := range csvContent {
		problems[i] = Problem{
			question: content[0],
			answer:   content[1],
		}
	}

	goodAnswers := 0
	scanner := bufio.NewScanner(os.Stdin)
	for i, problem := range problems {
		fmt.Printf("Problem %d - %s ?\n", (i + 1), problem.question)
		scanner.Scan()
		if scanner.Text() == problem.answer {
			goodAnswers++
		}
	}

	fmt.Printf("%d good answers on %d questions\n", goodAnswers, len(problems))

}
