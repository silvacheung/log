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
		e.KvMap.Put("level", LevelNameMapping[e.Level])
		e.KvMap.Put("time", e.Time.Format(time.RFC3339))
		if e.File != "" {
			e.KvMap.Put("file", e.File+":"+strconv.Itoa(e.Line))
			e.KvMap.Put("func", e.Func)
		}

		switch e.Format {
		case FmtEmptySeparate, FmtLineSeparate:
			e.KvMap.Put("message", fmt.Sprint(e.Args...))
		default:
			e.Format = strings.TrimSuffix(e.Format, FmtLineSeparate)
			e.KvMap.Put("message", fmt.Sprintf(e.Format, e.Args...))
		}

		return jsoniter.NewEncoder(e.Buffer).Encode(e.KvMap.Map())
	}

	switch e.Format {
	case FmtEmptySeparate, FmtLineSeparate:
		// log.InfoKv("xxx,yyy", 1, 2)
		if len(e.KvMap.Map()) > 0 {
			return jsoniter.NewEncoder(e.Buffer).Encode(e.KvMap.Map())
		}
		// log.Info(obj...)
		for _, arg := range e.Args {
			if err := jsoniter.NewEncoder(e.Buffer).Encode(arg); err != nil {
				return err
			}
		}
	default:
		// log.Infof("%s%d", "xxx", 1)
		e.Format = strings.TrimSuffix(e.Format, FmtLineSeparate)
		e.Buffer.WriteString(fmt.Sprintf(e.Format, e.Args...))
	}

	return nil
}
