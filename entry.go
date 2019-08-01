package log

import (
	"bytes"
	"runtime"
	"strings"
	"time"
)

type Entry struct {
	Logger *Logger
	Buffer *bytes.Buffer
	KvMap  *KvMap
	Level  Level
	Time   time.Time
	File   string
	Line   int
	Func   string
	Format string
	Args   []interface{}
}

func entry(logger *Logger) *Entry {
	return &Entry{Logger: logger, Buffer: new(bytes.Buffer), KvMap: newKvMap()}
}

func (e *Entry) write(depth int, level Level, format string, args ...interface{}) {
	e.Time = time.Now()
	e.Level = level
	e.Format = format
	e.Args = args
	if e.Logger.opt.fileLine {
		if pc, file, line, ok := runtime.Caller(depth); !ok {
			e.File = "???"
			e.Func = "???"
		} else {
			e.File, e.Line, e.Func = file, line, runtime.FuncForPC(pc).Name()
			e.Func = e.Func[strings.LastIndex(e.Func, "/")+1:]
		}
	}
	e.format()
	e.writer()
	e.release()
}

func (e *Entry) format() {
	if err := e.Logger.opt.formatter.Format(e); err != nil {
		e.Logger.opt.errorHandler(err)
	}
}

func (e *Entry) writer() {
	e.Logger.mux.Lock()
	_, err := e.Logger.opt.output.Write(e.Buffer.Bytes())
	e.Logger.mux.Unlock()
	if err != nil {
		e.Logger.opt.errorHandler(err)
	}
}

func (e *Entry) release() {
	e.Args, e.Line, e.File, e.Format, e.Func = nil, 0, "", "", ""
	e.KvMap.Reset()
	e.Buffer.Reset()
	e.Logger.entryPool.Put(e)
}

func (e *Entry) parseKv(format string, args ...interface{}) {
	if argLen := len(args); format != FmtEmptySeparate && format != FmtLineSeparate && argLen > 0 {
		fmtLen := len(format)
		var offset, index int
		for i, f := range format {
			if index < argLen { // format item <= args item
				if f == 44 { // 44 == ","
					e.KvMap.Put(format[:offset], args[index])
					format = format[offset+1:]
					index++
					offset = 0
					continue
				}
				if i == fmtLen-1 {
					e.KvMap.Put(format, args[index])
				}
				offset++
			}
		}
	}
}

func (e *Entry) log(level Level, format string, args ...interface{}) {
	if e.Logger.opt.level >= level {
		e.write(3, level, format, args...)
	}
}

// usage : kv(Level, "xxx,yyy,zzz", 1, 2, 3) => json:{"xxx":1, "yyy":2, "zzz":3} / text:xxx=1 yyy=2 zzz=3
func (e *Entry) kv(level Level, format string, args ...interface{}) {
	if e.Logger.opt.level >= level {
		e.parseKv(format, args...)
		e.write(3, level, FmtEmptySeparate)
	}
}

func (e *Entry) kvln(level Level, format string, args ...interface{}) {
	if e.Logger.opt.level >= level {
		e.parseKv(format, args...)
		e.write(3, level, FmtLineSeparate)
	}
}
