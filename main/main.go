package main

import (
	"github.com/antholord/poe-micro/indexer"
	"runtime"
)

func main() {
	go indexer.Run()
	runtime.Goexit()
}
