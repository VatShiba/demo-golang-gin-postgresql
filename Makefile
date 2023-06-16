server:
	go run cmd/main.go
start:
	nodemon --exec go run cmd/main.go --signal SIGTERM