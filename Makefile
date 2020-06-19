files:
	node data/build.js ./crater.png ./crops.png ./characters.png ./maps/base.json ./mars-base.png

run: files
	go run main.go

build: files
	go build -o mars-game-osx main.go
	appify -name "Mars Game" -icon ./icon.png ./mars-game-osx

downloads:
	gh api repos/:owner/:repo/releases | jq -c '.[] | .assets[] | [(.browser_download_url | match("download/[a-zA-Z-0-9\/.]+"; "g") | .string), .download_count]'
