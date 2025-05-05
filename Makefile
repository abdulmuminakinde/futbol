MAKEFLAGS = --silent

build:
	go build -o footy

run: build
	./footy

set: build
	./footy setToken

reset:
	rm ~/.config/futbol/token.txt
