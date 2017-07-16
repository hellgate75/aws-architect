package log

import (
	"fmt"
	"time"
)

type Logger struct {
	NoDebug bool
}

func (l *Logger) Log(value interface{}) {
	var LogDate time.Time = time.Now()
	fmt.Printf("[%s] INFO %v\n", LogDate.Format("2006-01-02 15:04:05.000"), value)
}

func (l *Logger) Error(err error) {
	var LogDate time.Time = time.Now()
	fmt.Printf("[%s] ERROR %s\n", LogDate.Format("2006-01-02 15:04:05.000"), err.Error())
}

func (l *Logger) WarningE(err error) {
	var LogDate time.Time = time.Now()
	fmt.Printf("[%s] WARNING %s\n", LogDate.Format("2006-01-02 15:04:05.000"), err.Error())
}

func (l *Logger) Debug(value interface{}) {
	if !l.NoDebug {
		var LogDate time.Time = time.Now()
		fmt.Printf("[%s] DEBUG %v\n", LogDate.Format("2006-01-02 15:04:05.000"), value)
	}
}
