# Install
go mod tidy
go mod download

# RUN SERVER
go run cmd/main.go

# BUILD
go build -o app ./cmd