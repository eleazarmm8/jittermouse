package main

import (
	"sync"
	"time"

	"github.com/go-vgo/robotgo"
	"github.com/lthibault/jitterbug"
	hook "github.com/robotn/gohook"
)

var wg *sync.WaitGroup

func stop(e hook.Event) {
	hook.End()
}

func main() {
	wg = new(sync.WaitGroup)

	hook.Register(hook.KeyDown, []string{"ctrl", "shift", "l"}, stop)
	go start()
	s := hook.Start()
	<-hook.Process(s)
}

func start() {
	ticker := jitterbug.New(
		time.Minute,
		&jitterbug.Norm{Stdev: time.Millisecond * 600},
	)

	defer ticker.Stop()

	for range ticker.C {
		robotgo.KeyTap("scrolllock") // Non-intrusive workaround for windows
		robotgo.KeyTap("scrolllock")
		robotgo.MoveSmoothRelative(1, 1)
		robotgo.MoveSmoothRelative(-1, -1)
	}
}
