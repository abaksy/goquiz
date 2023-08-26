package quiz

import (
	"bufio"
	"encoding/csv"
	"os"
	"strings"
	"time"
)

// Read data from CSV file and return it as array of array of strings.
// Each sub-array in the array contains [question answer]
func ParseQuestions(fileName string) [][]string {

	r, err := os.Open(fileName)
	if err != nil {
		panic(err)
	}
	defer r.Close()

	csvReader := csv.NewReader(r)
	questions, err := csvReader.ReadAll()
	if err != nil {
		panic(err)
	}

	return questions
}

// Run timed quiz and return the final score
// The timer sends a value through its channel C when it completes
func RunTimedQuiz(questions [][]string, timer int) int {
	scanner := bufio.NewScanner(os.Stdin)

	Green.Printf("Press enter to start the quiz....")
	scanner.Scan()

	quizTimer := time.NewTimer(time.Duration(timer) * time.Second)

	totalScore := 0
	i := 0

	for {
		select {
		// Timer triggered
		case <-quizTimer.C:
			Magenta.Println("Time's Up!")
			return totalScore
		// Ask question
		default:
			if i >= len(questions) {
				Green.Println("Quiz Complete!")
				return totalScore
			}
			question, correctAnswer := questions[i][0], questions[i][1]
			BoldBlue.Printf("%v = ", question)

			scanner.Scan()
			// Holds the string that scanned
			answer := scanner.Text()

			if strings.Compare(answer, correctAnswer) == 0 {
				totalScore += 1
			}
			i += 1
		}
	}
}
