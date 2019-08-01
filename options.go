package log

import (
	"fmt"
	"io"
	"os"
	"runtime/debug"
)

const (
	FmtEmptySeparate = ""
	FmtLineSeparate  = "\n"
)

// log level
type Level uint8

// const log level
const (
	PanicLevel Level = iota
	FatalLevel
	ErrorLevel
	WarnLevel
	InfoLevel
	DebugLevel
	TraceLevel
)

// log level string name mapping
var LevelNameMapping = map[Level]string{
	PanicLevel: "PANIC",
	FatalLevel: "FATAL",
	ErrorLevel: "ERROR",
	WarnLevel:  "WARN",
	InfoLevel:  "INFO",
	DebugLevel: "DEBUG",
	TraceLevel: "TRACE",
}

// type log error handler
type ErrorHandler func(err error)

// default log error handler
var errorHandler ErrorHandler = func(err error) {
	if _, err := fmt.Fprintf(os.Stderr, "log error :%v \n", err); err != nil {
		panic(err)
	}
	debug.PrintStack()
}

// log options
type option struct {
	output       io.Writer
	level        Level
	stdLevel     Level // the go std lib log level
	formatter    Formatter
	fileLine     bool
	noLock       bool
	errorHandler ErrorHandler
}

func initOptions(opts ...Option) (o *option) {

	o = &option{}

	for _, opt := range opts {
		opt(o)
	}

	if o.output == nil {
		o.output = os.Stderr
	}

	if o.formatter == nil {
		o.formatter = &TextFormatter{}
	}

	if o.errorHandler == nil {
		o.errorHandler = errorHandler
	}

	return
}

type Option func(*option)

// SET
func WithOutput(output io.Writer) Option {
	return func(o *option) {
		o.output = output
	}
}

func WithLevel(level Level) Option {
	return func(o *option) {
		o.level = level
	}
}

func WithStdLevel(level Level) Option {
	return func(o *option) {
		o.stdLevel = level
	}
}

func WithFormatter(formatter Formatter) Option {
	return func(o *option) {
		o.formatter = formatter
	}
}

func WithFileLine(fileLine bool) Option {
	return func(o *option) {
		o.fileLine = fileLine
	}
}

func WithNoLock(noLock bool) Option {
	return func(o *option) {
		o.noLock = noLock
	}
}

func WithErrorHandler(handler ErrorHandler) Option {
	return func(o *option) {
		o.errorHandler = handler
	}
}

// GET
func (o *option) GetOutput() io.Writer {
	return o.output
}

func (o *option) GetLevel() Level {
	return o.level
}

func (o *option) GetStdLevel() Level {
	return o.stdLevel
}

func (o *option) GetFormatter() Formatter {
	return o.formatter
}

func (o *option) GetFileLine() bool {
	return o.fileLine
}

func (o *option) GetNoLock(off bool) bool {
	return o.noLock
}

func (o *option) GetErrorHandler() ErrorHandler {
	return o.errorHandler
}
