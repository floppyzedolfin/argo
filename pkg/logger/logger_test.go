package logger

import (
	"time"
)

func ExampleLogPrintResults() {
	// reset default values for the purpose of  testing Log
	defaultLevel = Debug
	clock = brokenClock

	Log(Info, "Hello %s", "World!")
	// Output: {"time":"2021-03-29T12:00:00Z","level":"Info","message":"Hello World!","caller":"logger.ExampleLogPrintResults"}
}

func ExampleLogNoOutput() {
	// reset default values for the purpose of  testing Log
	defaultLevel = Warning
	clock = brokenClock

	Log(Info, "Hello %s", "World!")
	// Output:
}

func brokenClock() time.Time {
	return time.Date(2021, time.March, 29, 12, 0, 0, 0, time.UTC)
}
