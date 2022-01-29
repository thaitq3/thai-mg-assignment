package server

import (
	"time"

	"word-counting/utils"
)

func Serve() {
	InitDependencies()

	Start()

	utils.WaitShutdownSignal()
	// actions on shutdown
	utils.WaitOrTimeout(time.Minute*3, Stop())
}
