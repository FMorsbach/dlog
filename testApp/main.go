package main

import "github.com/FMorsbach/dlog"

func main() {
	dlog.Debug("Debug 0")
	dlog.Println("Print 0")

	dlog.SetDebug(false)

	dlog.Debug("Debug 1")
	dlog.Println("Print 1")
}
