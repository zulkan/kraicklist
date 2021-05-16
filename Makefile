BINARY=engine

build:
	go build -o ${BINARY} app/*.go

run:
	go run app/main.go
