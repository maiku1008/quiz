package memory

import (
	"fmt"
	"sync"

	"github.com/micuffaro/quiz/quiz_server/internal/dao"
	"github.com/micuffaro/quiz/quiz_server/internal/models"
)

func NewStorage() *Storage {
	return &Storage{
		store: map[string]models.Player{},
	}
}

type Storage struct {
	sync.RWMutex
	store map[string]models.Player
}

var _ dao.PlayerDAO = (*Storage)(nil)

func (s *Storage) GetPlayer(name string) (models.Player, error) {
	s.Lock()
	defer s.Unlock()

	player := s.store[name]
	return player, nil
}

func (s *Storage) CreatePlayer(name string) error {
	s.Lock()
	defer s.Unlock()

	if existing, ok := s.store[name]; ok {
		return fmt.Errorf("player with name %s already exists", existing.Name)
	}

	s.store[name] = models.Player{
		Name:  name,
		Score: 0,
	}
	return nil
}

func (s *Storage) IncreasePlayerScore(name string, score int) error {
	s.Lock()
	defer s.Unlock()

	player := s.store[name]
	player.Score += score
	s.store[name] = player
	return nil
}

func (s *Storage) ListAllPlayers() ([]models.Player, error) {
	players := make([]models.Player, len(s.store))
	idx := 0
	for _, player := range s.store {
		players[idx] = player
		idx++
	}
	return players, nil
}
