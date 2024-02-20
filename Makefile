run-exhibition:
	go run cmd/exhibition/main.go

test-coverage:
	mkdir -p coverage
	go test -race -short -v -coverprofile coverage/cover.out ./...
	go tool cover -html=coverage/cover.out

gen-swag:
	swag init -d ./cmd/exhibition,./handler/exhibihandler -o ./cmd/exhibition/doc --pd