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

type problem struct {
	questionPrompt string
	correctAnswer  string
}

func main() {

	csvFileName := flag.String("csv", "problems.csv", "A csv file in the format of 'question,answer'")
	timeLimit := flag.Int("limit", 30, "The time limit for the quiz in seconds")

	flag.Parse()

	CORRECT_SCORE := 0
	INCORRECT_SCORE := 0

	records, err := readCsvFile(*csvFileName)

	if err != nil {
		fmt.Print("Error loading data")
	}

	problems := parseLines(records)

	timer := time.NewTimer(time.Duration(*timeLimit) * time.Second)
	for i, p := range problems {
		fmt.Printf("Problem #%d: %s = ", i+1, p.questionPrompt)
		answerCh := make(chan string)
		go func() {
			answerCh <- getUserAnswer(p.questionPrompt)
		}()
		select {
		case <-timer.C:
			fmt.Println("Correct:", CORRECT_SCORE, "Incorrect:", INCORRECT_SCORE)
			return
		case userAnswer := <-answerCh:
			if strings.Compare(userAnswer, p.correctAnswer) == 0 {
				fmt.Println("Correct!")
				CORRECT_SCORE++
			} else {
				fmt.Println("Incorrect!")
				fmt.Printf("The correct answer is %s\n", p.correctAnswer)
				INCORRECT_SCORE++
			}
		}
	}

	fmt.Println("Correct:", CORRECT_SCORE, "Incorrect:", INCORRECT_SCORE)

}

func parseLines(lines [][]string) []problem {
	ret := make([]problem, len(lines))
	for i, line := range lines {
		ret[i] = problem{
			questionPrompt: line[0],
			correctAnswer:  line[1],
		}
	}
	return ret
}

func readCsvFile(filePath string) ([][]string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal("Error loading file "+filePath, err)
		return nil, err
	}
	defer file.Close()
	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		log.Fatal("Error loading file as CSV "+filePath, err)
		return nil, err
	}
	return records, nil
}

func getUserAnswer(question_prompt string) string {

	var userAnswer string
	fmt.Scanf("%s\n", &userAnswer)

	return userAnswer

}
