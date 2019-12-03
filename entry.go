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
	Map    map[string]interface{}
	Level  Level
	Time   time.Time
	File   string
	Line   int
	Func   string
	Format string
	Args   []interface{}
}

func entry(logger *Logger) *Entry {
	return &Entry{Logger: logger, Buffer: new(bytes.Buffer), Map: make(map[string]interface{}, 5)}
}

func (e *Entry) write(level Level, format string, args ...interface{}) {
	if e.Logger.opt.level < level {
		return
	}
	e.Time = time.Now()
	e.Level = level
	e.Format = format
	e.Args = args
	if e.Logger.opt.fileLine {
		if pc, file, line, ok := runtime.Caller(2); !ok {
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
	e.Buffer.Reset()
	// FIXED ME 'release the Map'
	e.Logger.entryPool.Put(e)
}
