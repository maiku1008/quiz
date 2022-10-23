package score

import (
	"fmt"
	"sort"

	"github.com/micuffaro/quiz/quiz_server/internal/models"
)

// GetPercentage gets how good the player has scored
// against a list of other players, in percentage
func GetPercentage(playerName string, players []models.Player) string {
	sort.Slice(players, func(i, j int) bool {
		// If two players have the same score, sort alphabetically
		if players[i].Score == players[j].Score {
			return players[i].Name > players[j].Name
		}
		return players[i].Score < players[j].Score
	})
	index := 0
	for i, player := range players {
		if player.Name == playerName {
			index = i
			break
		}
	}
	position := float64(index)
	percent := position / float64(len(players)) * 100
	return fmt.Sprintf("%.2f%%", percent)
}
