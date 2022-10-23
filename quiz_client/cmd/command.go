package cmd

import (
	"context"
	"fmt"
	"io"
	"log"

	pb "github.com/micuffaro/quiz/quiz"
	"github.com/spf13/cobra"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	server     string
	playerName string

	rootCmd = &cobra.Command{
		Use:   "quiz",
		Short: "Quiz is a client for quiz",
	}
	versionCmd = &cobra.Command{
		Use:   "version",
		Short: "Print the version number of Quiz",
		Long:  `All software has versions. This is Quiz'`,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Quiz version 1.0")
		},
	}
	playCmd = &cobra.Command{
		Use:   "play",
		Short: "Start the quiz",
		RunE:  playFunc,
	}
	getScoreCmd = &cobra.Command{
		Use:   "score",
		Short: "Get a player score",
		RunE:  getScoreFunc,
	}
)

func init() {
	rootCmd.AddCommand(versionCmd)
	rootCmd.AddCommand(playCmd)
	rootCmd.AddCommand(getScoreCmd)
	rootCmd.PersistentFlags().StringVar(&server, "server", "localhost:50051", "the address to connect to")
	rootCmd.PersistentFlags().StringVar(&playerName, "player", "", "The current player name (required)")
	rootCmd.MarkPersistentFlagRequired("player")
}

func Execute() error {
	return rootCmd.Execute()
}

func playFunc(cmd *cobra.Command, args []string) error {
	conn, err := grpc.Dial(server, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("could not get grpc connection: %v", err)
		return err
	}
	defer conn.Close()

	c := pb.NewQuizClient(conn)
	if err != nil {
		log.Fatalf("could not start quiz client: %v", err)
		return err
	}
	ctx := context.Background()
	name := playerName

	qStream, err := c.GetQuestions(ctx, &pb.QuestionRequest{Name: name})
	if err != nil {
		log.Fatalf("could not stream questions to user: %v", err)
		return err
	}
	// stream the questions and store them in memory
	questions := []*pb.QuestionResponse{}
	for {
		question, err := qStream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("client.GetQuestions failed: %v", err)
		}
		questions = append(questions, question)
	}

	// have the user answer them
	answers := []*pb.AnswerRequest{}
	for _, question := range questions {
		fmt.Printf("%d + %d = ?\n", question.X, question.Y)
		fmt.Printf("A. %d\n", question.A)
		fmt.Printf("B. %d\n", question.B)
		fmt.Printf("C. %d\n", question.C)
		fmt.Printf("D. %d\n", question.D)

		answer := getPlayerAnswer(question)

		answerReq := &pb.AnswerRequest{
			Name:   name,
			X:      question.X,
			Y:      question.Y,
			Answer: answer,
		}
		answers = append(answers, answerReq)
	}

	// stream back answers from client
	aStream, err := c.GiveAnswers(ctx)
	if err != nil {
		log.Fatalf("could not stream answers from client: %v", err)
		return err
	}
	for _, answer := range answers {
		err = aStream.Send(answer)
		if err != nil {
			log.Fatalf("client.GiveAnswers failed: %v", err)
			return err
		}
	}
	_, err = aStream.CloseAndRecv()
	if err != nil {
		log.Fatalf("could not close stream failed: %v", err)
		return err
	}
	return nil
}

func getScoreFunc(cmd *cobra.Command, args []string) error {
	conn, err := grpc.Dial(server, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("could not get grpc connection: %v", err)
		return err
	}
	defer conn.Close()

	c := pb.NewQuizClient(conn)
	if err != nil {
		log.Fatalf("could not start quiz client: %v", err)
		return err
	}
	ctx := context.Background()
	// get scores
	resp, err := c.GetScore(ctx, &pb.ScoreRequest{Name: playerName})
	if err != nil {
		log.Fatalf("could not get score for %s: %v", playerName, err)
		return err
	}
	fmt.Printf("%s, your score is %s, which is better than %s of total players.\n", playerName, resp.Score, resp.Percent)
	return nil
}

func getPlayerAnswer(question *pb.QuestionResponse) int32 {
	answers := map[string]int32{
		"A": question.A,
		"B": question.B,
		"C": question.C,
		"D": question.D,
	}

	var answer string
	fmt.Println("Answer (A, B, C or D): ")
	fmt.Scanln(&answer)
	for answer != "A" && answer != "B" && answer != "C" && answer != "D" {
		fmt.Println("Answer (A, B, C or D): ")
		fmt.Scanln(&answer)
	}
	return answers[answer]
}
