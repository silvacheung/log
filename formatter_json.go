package log

import (
	"fmt"
	"github.com/json-iterator/go"
	"strconv"
	"strings"
	"time"
)

type JsonFormatter struct {
	// 忽略基础字段
	IgnoreBasicFields bool
}

func (f *JsonFormatter) Format(e *Entry) error {

	if !f.IgnoreBasicFields {
		e.Map["level"] = LevelNameMapping[e.Level]
		e.Map["time"] = e.Time.Format(time.RFC3339)
		if e.File != "" {
			e.Map["file"] = e.File + ":" + strconv.Itoa(e.Line)
			e.Map["func"] = e.Func
		}

		switch e.Format {
		case FmtEmptySeparate, FmtLineSeparate:
			e.Map["message"] = fmt.Sprint(e.Args...)
		default:
			e.Map["message"] = fmt.Sprintf(strings.TrimSuffix(e.Format, FmtLineSeparate), e.Args...)
		}

		return jsoniter.NewEncoder(e.Buffer).Encode(e.Map)
	}

	switch e.Format {
	case FmtEmptySeparate, FmtLineSeparate:
		// log.Info/Infoln(obj...)
		for _, arg := range e.Args {
			if err := jsoniter.NewEncoder(e.Buffer).Encode(arg); err != nil {
				return err
			}
		}
	default:
		// log.Infof("%s%d", "xxx", 1)
		e.Buffer.WriteString(fmt.Sprintf(e.Format, e.Args...))
	}

	return nil
}
