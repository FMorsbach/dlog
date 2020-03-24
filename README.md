# dlog
dlog is an extension of the standard golang log package to support a toggable debug log level. 

It can be used like the [standard log package](https://golang.org/pkg/log/) but extends the functionality by adding a switch to toggle debug logs and corresponding debug methods:

- `func Debug(v ...interface{})`
- `func Debugf(format string, v ...interface{})`
- `func Debugln(v ...interface{})`
- `func SetDebug(b bool)`
