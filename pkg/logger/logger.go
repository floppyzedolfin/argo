package logger

import (
	"encoding/json"
	"fmt"
	"os"
	"runtime"
	"strings"
	"time"
)

type level int

const (
	Debug level = iota
	Info
	Warning
	Error
	noLog // set the default value to noLog to discard all logs.
)

var (
	// these default parameters exist for testing purposes only - they aren't exposed, so it's "ok" to have them as global variables
	defaultLevel                  = Debug    // change this to set a new logging threshold
	clock        func() time.Time = time.Now // needed for time
)

// Log writes a message on stdout, if the log level is high enough.
func Log(lvl level, format string, a ...interface{}) {
	if lvl >= defaultLevel {
		// skipping 2 layers, in order to get to the caller of this func
		msg := buildLogEntry(lvl, fmt.Sprintf(format, a...), clock, 2)
		marshalledLog, _ := json.Marshal(msg)
		fmt.Printf("%v\n", string(marshalledLog))
	}
}

type logEntry struct {
	Time    string `json:"time"`
	Level   string `json:"level"`
	Message string `json:"message"`
	Caller  string `json:"caller"`
}

func buildLogEntry(lvl level, msg string, t func() time.Time, skip int) logEntry {
	// get info about the caller
	callerFunc := ""
	fpcs := make([]uintptr, 1)
	// read the stack
	n := runtime.Callers(skip+1, fpcs)
	if n != 0 {
		caller := runtime.FuncForPC(fpcs[0] - 1)
		if caller != nil {
			callerFunc = caller.Name()
			path := strings.Split(callerFunc, string(os.PathSeparator))
			callerFunc = path[len(path)-1]
		}
	}

	return logEntry{
		Time:    t().Format(time.RFC3339Nano),
		Level:   levelNames[lvl],
		Message: msg,
		Caller:  callerFunc,
	}
}

var levelNames = map[level]string{
	Debug:   "Debug",
	Info:    "Info",
	Warning: "Warning",
	Error:   "Error",
}
