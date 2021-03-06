run: 
	go run main.go

build:
	go build -o mars-game-osx main.go
	appify -name "Mars Game" -icon ./icon.png ./mars-game-osx

downloads:
	gh api repos/:owner/:repo/releases | jq -c '.[] | .assets[] | [(.browser_download_url | match("download/[a-zA-Z-0-9\/.]+"; "g") | .string), .download_count]'

perf: build
	PERFON=true ./mars-game-osx > frames.csv
	go tool pprof -svg cpuperf-0.perf
