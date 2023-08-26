package main

import (
	"flag"

	quiz "github.com/abaksy/goquiz/pkg"
	"github.com/fatih/color"
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
	boldGreen := quiz.Green.Add(color.Bold)
	boldGreen.Printf("\n\nYou scored %v out of %v!\n", score, len(questions))
}
