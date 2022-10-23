package questions

import "testing"

func TestGenerateQuestions(t *testing.T) {
	questions := Generate(10000)
	for _, question := range questions {
		answer := question.X + question.Y
		choices := []int{
			question.A,
			question.B,
			question.C,
			question.D,
		}
		unique := map[int]bool{}
		foundCorrect := false
		for _, choice := range choices {
			if choice == answer {
				foundCorrect = true
			}
			unique[choice] = true
		}
		if !foundCorrect {
			t.Fatalf("could not find correct choice %d in question: %+v", answer, question)
		}
		if len(unique) != 4 {
			t.Fatalf("question does not have 4 unique choices: %+v", question)
		}
	}
}
