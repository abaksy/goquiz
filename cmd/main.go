package main

import (
	"flag"
	"fmt"

	quiz "github.com/abaksy/goquiz/pkg"
)

func main() {

	fileName := flag.String("file", "problems.csv", "Filename to read quiz problems from")
	timeDuration := flag.Int("time", 30, "Time for quiz to run")
	flag.Parse()

	// Read questions from CSV file
	questions := quiz.ParseQuestions(*fileName)

	// Run timed quiz
	score := quiz.RunTimedQuiz(questions, *timeDuration)

	// Print final score
	fmt.Printf("\n\nYou scored %v out of %v!\n", score, len(questions))
}
