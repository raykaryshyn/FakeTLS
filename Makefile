all: client server

client: client.c
	gcc -g client.c -o client

server: server.go
	go build -o server server.go
