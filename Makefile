
$(pkged.go):
	pkger

run: pkged.go
	go run *.go

build: pkged.go
	go build -o mars-game-osx *.go

