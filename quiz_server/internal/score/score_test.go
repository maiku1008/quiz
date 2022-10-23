package score

import (
	"testing"

	"github.com/micuffaro/quiz/quiz_server/internal/models"
)

func TestGetPercentage(t *testing.T) {
	players := []models.Player{
		{Name: "Gopher", Score: 7},
		{Name: "Alice", Score: 15},
		{Name: "Vera", Score: 14},
		{Name: "Pichael", Score: 4},
	}

	var testCases = []struct {
		name    string
		percent string
	}{
		{name: "Alice", percent: "75.00%"},
		{name: "Vera", percent: "50.00%"},
		{name: "Gopher", percent: "25.00%"},
		{name: "Pichael", percent: "0.00%"},
	}

	for _, tc := range testCases {
		got := GetPercentage(tc.name, players)
		if tc.percent != got {
			t.Errorf("player name: %s, want %s, got %s", tc.name, tc.percent, got)
		}
	}
}

func TestGetPercentageSameScore(t *testing.T) {
	players := []models.Player{
		{Name: "Alice", Score: 8},
		{Name: "Vera", Score: 8},
		{Name: "Gopher", Score: 7},
		{Name: "Pichael", Score: 7},
	}

	var testCases = []struct {
		name    string
		percent string
	}{
		{name: "Alice", percent: "75.00%"},
		{name: "Vera", percent: "50.00%"},
		{name: "Gopher", percent: "25.00%"},
		{name: "Pichael", percent: "0.00%"},
	}

	for _, tc := range testCases {
		got := GetPercentage(tc.name, players)
		if tc.percent != got {
			t.Errorf("player name: %s, want %s, got %s", tc.name, tc.percent, got)
		}
	}
}
