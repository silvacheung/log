package log

import (
	"io"
	"sync"
	"unsafe"
)

var std *Logger

type Logger struct {
	opt       *option
	mux       *MutexWrap
	entryPool *sync.Pool
}

func init() {
	std = New()
}

func New(opts ...Option) *Logger {
	logger := &Logger{opt: initOptions(opts...), mux: &MutexWrap{}}
	logger.mux.NoLock(logger.opt.noLock)
	logger.entryPool = &sync.Pool{New: func() interface{} { return entry(logger) }}
	return logger
}

func StdLogger() *Logger {
	return std
}

func SetOptions(opts ...Option) {
	std.SetOptions(opts...)
}

func (l *Logger) SetOptions(opts ...Option) {
	defer l.mux.NoLock(l.opt.noLock)
	for _, optFun := range opts {
		optFun(l.opt)
	}
}

func Writer() io.Writer {
	return std
}

func (l *Logger) Writer() io.Writer {
	return l
}

func (l *Logger) Write(data []byte) (int, error) {
	l.entry().write(l.opt.stdLevel, FmtEmptySeparate, *(*string)(unsafe.Pointer(&data)))
	return 0, nil
}

func (l *Logger) entry() *Entry {
	return l.entryPool.Get().(*Entry)
}

func (l *Logger) Debugf(format string, args ...interface{}) {
	l.entry().write(DebugLevel, format, args...)
}

func (l *Logger) Infof(format string, args ...interface{}) {
	l.entry().write(InfoLevel, format, args...)
}

func (l *Logger) Warnf(format string, args ...interface{}) {
	l.entry().write(WarnLevel, format, args...)
}

func (l *Logger) Errorf(format string, args ...interface{}) {
	l.entry().write(ErrorLevel, format, args...)
}

func (l *Logger) Fatalf(format string, args ...interface{}) {
	l.entry().write(FatalLevel, format, args...)
}

func (l *Logger) Panicf(format string, args ...interface{}) {
	l.entry().write(PanicLevel, format, args...)
}

func (l *Logger) Tracef(format string, args ...interface{}) {
	l.entry().write(TraceLevel, format, args...)
}

func (l *Logger) Debug(args ...interface{}) {
	l.entry().write(DebugLevel, FmtEmptySeparate, args...)
}

func (l *Logger) Info(args ...interface{}) {
	l.entry().write(InfoLevel, FmtEmptySeparate, args...)
}

func (l *Logger) Warn(args ...interface{}) {
	l.entry().write(WarnLevel, FmtEmptySeparate, args...)
}

func (l *Logger) Error(args ...interface{}) {
	l.entry().write(ErrorLevel, FmtEmptySeparate, args...)
}

func (l *Logger) Fatal(args ...interface{}) {
	l.entry().write(FatalLevel, FmtEmptySeparate, args...)
}

func (l *Logger) Panic(args ...interface{}) {
	l.entry().write(PanicLevel, FmtEmptySeparate, args...)
}

func (l *Logger) Trace(args ...interface{}) {
	l.entry().write(TraceLevel, FmtEmptySeparate, args...)
}

func (l *Logger) Debugln(args ...interface{}) {
	l.entry().write(DebugLevel, FmtLineSeparate, args...)
}

func (l *Logger) Infoln(args ...interface{}) {
	l.entry().write(InfoLevel, FmtLineSeparate, args...)
}

func (l *Logger) Warnln(args ...interface{}) {
	l.entry().write(WarnLevel, FmtLineSeparate, args...)
}

func (l *Logger) Errorln(args ...interface{}) {
	l.entry().write(ErrorLevel, FmtLineSeparate, args...)
}

func (l *Logger) Fatalln(args ...interface{}) {
	l.entry().write(FatalLevel, FmtLineSeparate, args...)
}

func (l *Logger) Panicln(args ...interface{}) {
	l.entry().write(PanicLevel, FmtLineSeparate, args...)
}

func (l *Logger) Traceln(args ...interface{}) {
	l.entry().write(TraceLevel, FmtLineSeparate, args...)
}

// std logger

func Debugf(format string, args ...interface{}) {
	std.entry().write(DebugLevel, format, args...)
}

func Infof(format string, args ...interface{}) {
	std.entry().write(InfoLevel, format, args...)
}

func Warnf(format string, args ...interface{}) {
	std.entry().write(WarnLevel, format, args...)
}

func Errorf(format string, args ...interface{}) {
	std.entry().write(ErrorLevel, format, args...)
}

func Fatalf(format string, args ...interface{}) {
	std.entry().write(FatalLevel, format, args...)
}

func Panicf(format string, args ...interface{}) {
	std.entry().write(PanicLevel, format, args...)
}

func Tracef(format string, args ...interface{}) {
	std.entry().write(TraceLevel, format, args...)
}

func Debug(args ...interface{}) {
	std.entry().write(DebugLevel, FmtEmptySeparate, args...)
}

func Info(args ...interface{}) {
	std.entry().write(InfoLevel, FmtEmptySeparate, args...)
}

func Warn(args ...interface{}) {
	std.entry().write(WarnLevel, FmtEmptySeparate, args...)
}

func Error(args ...interface{}) {
	std.entry().write(ErrorLevel, FmtEmptySeparate, args...)
}

func Fatal(args ...interface{}) {
	std.entry().write(FatalLevel, FmtEmptySeparate, args...)
}

func Panic(args ...interface{}) {
	std.entry().write(PanicLevel, FmtEmptySeparate, args...)
}

func Trace(args ...interface{}) {
	std.entry().write(TraceLevel, FmtEmptySeparate, args...)
}

func Debugln(args ...interface{}) {
	std.entry().write(DebugLevel, FmtLineSeparate, args...)
}

func Infoln(args ...interface{}) {
	std.entry().write(InfoLevel, FmtLineSeparate, args...)
}

func Warnln(args ...interface{}) {
	std.entry().write(WarnLevel, FmtLineSeparate, args...)
}

func Errorln(args ...interface{}) {
	std.entry().write(ErrorLevel, FmtLineSeparate, args...)
}

func Fatalln(args ...interface{}) {
	std.entry().write(FatalLevel, FmtLineSeparate, args...)
}

func Panicln(args ...interface{}) {
	std.entry().write(PanicLevel, FmtLineSeparate, args...)
}

func Traceln(args ...interface{}) {
	std.entry().write(TraceLevel, FmtLineSeparate, args...)
}
