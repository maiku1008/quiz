package questions

import (
	"fmt"
	"math/rand"
)

// Generate a sequence of questions
func Generate(qty int) []*Question {
	questions := []*Question{}
	for i := 1; i <= qty; i++ {
		var q = NewQuestion()
		answer := randInt(5, 100)
		q.X, q.Y = q.generateAddends(answer)
		q.Question = q.stringify()
		q.fillAnswers(answer)
		questions = append(questions, q)
	}
	return questions
}

type Question struct {
	Question   string
	X, Y       int
	A, B, C, D int
}

func NewQuestion() *Question {
	return &Question{}
}

// addend + addend = sum
func (q *Question) generateAddends(sum int) (int, int) {
	x := randInt(1, sum)
	y := sum - x
	return x, y
}

// Fill the Question object with 3 unique bogus answers and a correct one
func (q *Question) fillAnswers(answer int) {
	options := []*int{&q.A, &q.B, &q.C, &q.D}
	correct := rand.Intn(3)
	set := map[int]bool{}
	for i, option := range options {
		if i == correct {
			*option = answer
			continue
		}
		n := rand.Intn(answer)
		for set[n] {
			n = rand.Intn(answer)
		}
		*option = n
		set[n] = true
	}
}

func (q *Question) stringify() string {
	return fmt.Sprintf("%d + %d = ?", q.X, q.Y)
}

func randInt(min, max int) int {
	if max-min <= 0 {
		return 0
	}
	return min + rand.Intn(max-min)
}
