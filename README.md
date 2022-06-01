# Credit Score
Microservices project

## Requirements
- Golang

## How to run
- `cd auth-service && go run main.go`, Because this project doesn't have docker yet, you need to keep the auth-service running splited terminal
- `cd debts-service && go run main.go` to setup the Debts Service, it depends on the Auth service
- `cd financial-events && go run main.go`
- `cd score-service && go run main.go`

## How to consume the API
- Import the RequestsCollection.json to your Insomnia App to test