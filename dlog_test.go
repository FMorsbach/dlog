package dlog

import (
	"log"
	"os"
	"testing"
)

func TestCustomLogger(t *testing.T) {

	logger := New(os.Stderr, "", log.LstdFlags, true)

	logger.Println("Print0")
	logger.Debug("Debug0")

	logger.SetDebug(false)

	logger.Println("Print1")
	logger.Debug("Debug1")
}

func TestDefaultLogger(t *testing.T) {

	Debug("Debug0")

	SetDebug(false)

	Debug("Debug1")
}
