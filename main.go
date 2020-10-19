package main

import (
	"hstransfer/lib"
	"os"
)
import . "fmt"

func main() {
	Printf("%v", os.Args)
	runOnce := lib.ArgsContainFlag("--run-once") // TODO use flag instead
	mode := os.Getenv("HSTRANSFER_MODE")
	Printf("# hstransfer[runOnce=%v]\n", runOnce)
	if mode == "uploader" {
		uploader(runOnce)
	} else {
		downloader(runOnce)
	}
}
