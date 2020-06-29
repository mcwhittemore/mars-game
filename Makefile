files:
	node data/build.js ./*.png ./maps/*.json ./items/*.json

run: files
	go run main.go

build: files
	go build -o mars-game-osx main.go
	appify -name "Mars Game" -icon ./icon.png ./mars-game-osx

downloads:
	gh api repos/:owner/:repo/releases | jq -c '.[] | .assets[] | [(.browser_download_url | match("download/[a-zA-Z-0-9\/.]+"; "g") | .string), .download_count]'

perf: build
	PERFON=true ./mars-game-osx > frames.csv
	go tool pprof -svg cpuperf-0.perf

world:
	go run main.go world-builder ./maps/base.json ./items/base-structure.json
	make files
