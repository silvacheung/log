package log

import (
	"fmt"
	"strconv"
	"time"
)

type TextFormatter struct {
	IgnoreBasicFields bool
}

func (f *TextFormatter) Format(e *Entry) error {

	if !f.IgnoreBasicFields {
		e.Buffer.WriteString("level=")
		e.Buffer.WriteString(LevelNameMapping[e.Level])
		e.Buffer.WriteString(" time=")
		e.Buffer.WriteString(e.Time.Format(time.RFC3339)) // allocs
		if e.File != "" {
			e.Buffer.WriteString(" file=")
			e.Buffer.WriteString(e.File)
			e.Buffer.WriteString(":")
			e.Buffer.WriteString(strconv.Itoa(e.Line)) // allocs
			e.Buffer.WriteString(" func=")
			e.Buffer.WriteString(e.Func)
		}
		e.Buffer.WriteString(" ")
	}

	e.KvMap.Range(func(k string, v interface{}) bool {
		e.Buffer.WriteString(k)
		e.Buffer.WriteString("=")
		e.Buffer.WriteString(fmt.Sprint(v)) // allocs
		e.Buffer.WriteString(" ")
		return true
	})

	if !f.IgnoreBasicFields {
		e.Buffer.WriteString("message=")
	}

	switch e.Format {
	case FmtEmptySeparate:
		e.Buffer.WriteString(fmt.Sprint(e.Args...)) // allocs
	case FmtLineSeparate:
		e.Buffer.WriteString(fmt.Sprintln(e.Args...)) // allocs
	default:
		e.Buffer.WriteString(fmt.Sprintf(e.Format, e.Args...)) // allocs
	}

	return nil
}
