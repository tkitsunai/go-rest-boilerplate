test:
	go test -v ./...

run:
	go run cmd/api/main.go

setupAir:
	go install github.com/cosmtrek/air@latest

runAir:
	air
