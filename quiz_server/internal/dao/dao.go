package dao

import "github.com/micuffaro/quiz/quiz_server/internal/models"

type PlayerDAO interface {
	GetPlayer(name string) (models.Player, error)
	CreatePlayer(name string) error
	IncreasePlayerScore(name string, score int) error
	ListAllPlayers() ([]models.Player, error)
}
