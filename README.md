## Quiz

Build a quiz server and client application talking to each other.

## Components:

- gRPC for communication between client and server
- in-memory database, with data access object for future extension
- CLI using https://github.com/spf13/cobra for the client

## User stories/Use cases:

- User should be able to get questions with a number of answers
- User should be able to select just one answer per question.
- User should be able to answer all the questions and then post his/hers answers and get back how many correct answers they had, displayed to the user.
- User should see how good he/she did compare to others that have taken the quiz, "You were better than 60% of all quizzers"
- If two users have the same score, sort alphabetically

## How to run

Run the server first:
```
$ go run quiz_server/main.go
```

In another terminal, run the client
```
# Start the quiz for player mike
$ go run quiz_client/main.go play --player mike
# See the score for player mike
$ go run quiz_client/main.go score --player mike
```

## Development

For modifying the grpc interface, use
```
protoc --go_out=./quiz --go_opt=paths=source_relative \
    --go-grpc_out=. --go-grpc_opt=paths=source_relative \
    quiz.proto
```
