
$(pkged.go):

pkged.go:
	pkger

run: pkged.go
	go run *.go

build: pkged.go
	go build -o mars-game-osx main.go
	appify -name "Mars Game" -icon ./icon.png ./mars-game-osx

