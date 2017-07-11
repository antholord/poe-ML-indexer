package main

import (
	"github.com/antholord/poe-ML-indexer/indexer"
	"runtime"
)

func main() {
	go indexer.Run()
	runtime.Goexit()
}
