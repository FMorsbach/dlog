package dlog

import (
	"io"
	"log"
	"os"
	"sync"
)

type DLogger struct {
	*log.Logger
	deb_mu sync.Mutex
	debug  bool
}

var std = New(os.Stderr, "", log.LstdFlags, true)

func New(out io.Writer, prefix string, flag int, debug bool) *DLogger {
	return &DLogger{Logger: log.New(out, prefix, flag), debug: debug}
}

func (d *DLogger) Debug(v ...interface{}) {
	if d.debug {
		d.Print(v...)
	}
}

func (d *DLogger) Debugf(format string, v ...interface{}) {
	if d.debug {
		d.Printf(format, v...)
	}
}

func (d *DLogger) Debugln(v ...interface{}) {
	if d.debug {
		d.Println(v...)
	}
}

func (d *DLogger) SetDebug(b bool) {
	d.deb_mu.Lock()
	defer d.deb_mu.Unlock()
	d.debug = b
}

func Debug(v ...interface{}) {
	std.Debug(v...)
}

func Debugf(format string, v ...interface{}) {
	std.Debugf(format, v...)
}

func Debugln(v ...interface{}) {
	std.Debugln(v...)
}

func SetDebug(b bool) {
	std.SetDebug(b)
}

func Fatal(v ...interface{}) {
	std.Fatal(v...)
}
func Fatalf(format string, v ...interface{}) {
	std.Fatalf(format, v...)
}
func Fatalln(v ...interface{}) {
	std.Fatalln(v...)
}
func Flags() int {
	return std.Flags()
}
func Output(calldepth int, s string) error {
	return std.Output(calldepth, s)
}
func Panic(v ...interface{}) {
	std.Panic(v...)
}
func Panicf(format string, v ...interface{}) {
	std.Panicf(format, v...)
}
func Panicln(v ...interface{}) {
	std.Panicln(v...)
}
func Prefix() string {
	return std.Prefix()
}
func Print(v ...interface{}) {
	std.Print(v...)
}
func Printf(format string, v ...interface{}) {
	std.Printf(format, v...)
}
func Println(v ...interface{}) {
	std.Println(v...)
}
func SetFlags(flag int) {
	std.SetFlags(flag)
}
func SetOutput(w io.Writer) {
	std.SetOutput(w)
}
func SetPrefix(prefix string) {
	std.SetPrefix(prefix)
}
func Writer() io.Writer {
	return std.Writer()
}
