package perf

import (
	"fmt"
	"os"
	"runtime/pprof"
	"time"
)

var startTime time.Time
var cpupath string
var cpufile *os.File
var count int
var frames int
var second <-chan time.Time

func Update() {
	if startTime.IsZero() {
		startTime = time.Now()
		second = time.Tick(time.Second)
		start()
	}
	frames++
	drift := time.Since(startTime).Seconds()

	select {
	case <-second:
		if frames > 60 {
			frames = 60
		}
		fmt.Printf("%f,%d\n", drift, frames)
		frames = 0
	default:
	}

	if drift > 60 {
		end()
		os.Exit(0)
	}

}

func start() {
	cpupath = fmt.Sprintf("./cpuperf-%d.perf", count)
	fmt.Println("Starting new profile:", cpupath)
	cpufile, err := os.Create(cpupath)
	if err != nil {
		panic(err)
	}
	pprof.StartCPUProfile(cpufile)
}

func end() {
	pprof.StopCPUProfile()
	count++
	cpufile.Close()
}
