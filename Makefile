files:
	node data/build.js ./crater.png ./crops.png ./characters.png

run: files
	go run main.go

build: files
	go build -o mars-game-osx main.go
	appify -name "Mars Game" -icon ./icon.png ./mars-game-osx

