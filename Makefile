MAKEFLAGS = --silent

build:
	go build -o footy

run: build
	./footy

set: build
	./footy setToken

today: build
	./footy today

reset:
	rm ~/.config/futbol/token.txt
