# Install
go mod tidy
go mod download

# go install air for live reload
go install github.com/air-verse/air@latest

# RUN SERVER
go run cmd/main.go
air

# BUILD
go build -o app ./cmd
air