package server

import (
	"context"
	"io"
	"log"
	"strconv"

	"github.com/micuffaro/quiz/quiz_server/internal/dao/memory"
	"github.com/micuffaro/quiz/quiz_server/internal/questions"
	"github.com/micuffaro/quiz/quiz_server/internal/score"

	pb "github.com/micuffaro/quiz/quiz"
)

const numberOfQuestions = 10

type Server struct {
	pb.UnimplementedQuizServer
	Storage *memory.Storage
}

func (q *Server) GetQuestions(questionRequest *pb.QuestionRequest, stream pb.Quiz_GetQuestionsServer) error {
	name := questionRequest.GetName()
	log.Printf("Received GetQuestions from player: %v", name)

	// Create or update the player
	player, err := q.Storage.GetPlayer(name)
	if err != nil {
		log.Printf("error while getting player %s: %v", name, err)
		return err
	}
	if player.Name == "" && player.Score == 0 {
		err := q.Storage.CreatePlayer(name)
		if err != nil {
			log.Printf("error while creating player %s: %v", name, err)
			return err
		}
	}

	for _, question := range questions.Generate(numberOfQuestions) {
		resp := &pb.QuestionResponse{
			Question: question.Question,
			A:        int32(question.A),
			B:        int32(question.B),
			C:        int32(question.C),
			D:        int32(question.D),
			X:        int32(question.X),
			Y:        int32(question.Y),
		}
		err := stream.Send(resp)
		if err != nil {
			return err
		}
	}
	return nil
}

func (q *Server) GiveAnswers(stream pb.Quiz_GiveAnswersServer) error {
	score := 0
	for {
		answer, err := stream.Recv()
		if err == io.EOF {
			return stream.SendAndClose(&pb.AnswerResponse{
				Message: "ok",
			})
		}
		if err != nil {
			return err
		}
		correct := answer.X + answer.Y
		if answer.Answer == correct {
			score++
			err = q.Storage.IncreasePlayerScore(answer.Name, 1)
			if err != nil {
				return err
			}
		}
		log.Printf("player %s got correct %d answers", answer.Name, score)
	}
}

func (q *Server) GetScore(ctx context.Context, in *pb.ScoreRequest) (*pb.ScoreResponse, error) {
	name := in.GetName()
	// Create or update the player
	player, err := q.Storage.GetPlayer(name)
	if err != nil {
		log.Printf("error while getting player %s: %v", name, err)
		return nil, err
	}
	if player.Name == "" && player.Score == 0 {
		err := q.Storage.CreatePlayer(name)
		if err != nil {
			log.Printf("error while creating player %s: %v", name, err)
			return nil, err
		}
	}
	log.Printf("Received GetScore request from player: %v", name)
	players, err := q.Storage.ListAllPlayers()
	if err != nil {
		log.Printf("error while listing all players: %v", err)
		return nil, err
	}
	percentage := score.GetPercentage(name, players)
	log.Println("player score: ", player.Score)
	score := strconv.Itoa(player.Score)
	return &pb.ScoreResponse{
		Score:   score,
		Percent: percentage,
		Message: "ok",
	}, nil
}
