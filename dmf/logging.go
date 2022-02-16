package dmf

import (
	"fmt"
	"io"
	"time"
)

type Log struct {
	Writer io.Writer
}

func (log *Log) formatString(s string) string {
	var now = time.Now()
	var prefix = fmt.Sprintf("[%s]", now.Format("2006-01-02 15:04:05"))
	return fmt.Sprintf("%s %s\n", prefix, s)
}

func (log *Log) write(s string) {
	io.WriteString(log.Writer, s)
}

func (log *Log) Info(s string) {
	log.write(log.formatString(s))
}

func (log *Log) InfoFormat(s string, params ...interface{}) {
	log.Info(fmt.Sprintf(s, params))
}

func (log *Log) Error(s string) {
	var value = log.formatString(s)
	log.write(value)
	panic(value)
}
