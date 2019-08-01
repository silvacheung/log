package log

import (
	"io"
	"sync"
	"unsafe"
)

var stdLogger *Logger

type Logger struct {
	opt       *option
	mux       *MutexWrap
	entryPool *sync.Pool
}

func init() {
	stdLogger = New()
}

func New(opts ...Option) *Logger {
	logger := &Logger{opt: initOptions(opts...), mux: &MutexWrap{}}
	logger.mux.NoLock(logger.opt.noLock)
	logger.entryPool = &sync.Pool{New: func() interface{} { return entry(logger) }}
	return logger
}

func StdLogger() *Logger {
	return stdLogger
}

func SetOptions(opts ...Option) {
	stdLogger.SetOptions(opts...)
}

func (l *Logger) SetOptions(opts ...Option) {
	defer l.mux.NoLock(l.opt.noLock)
	for _, optFun := range opts {
		optFun(l.opt)
	}
}

func Writer() io.Writer {
	return stdLogger
}

func (l *Logger) Writer() io.Writer {
	return l
}

func (l *Logger) Write(data []byte) (int, error) {
	l.entry().log(l.opt.stdLevel, FmtEmptySeparate, *(*string)(unsafe.Pointer(&data)))
	return 0, nil
}

func (l *Logger) entry() *Entry {
	return l.entryPool.Get().(*Entry)
}

func (l *Logger) InfoKv(format string, args ...interface{}) {
	l.entry().kv(InfoLevel, format, args...)
}

func (l *Logger) DebugKv(format string, args ...interface{}) {
	l.entry().kv(DebugLevel, format, args...)
}

func (l *Logger) WarnKv(format string, args ...interface{}) {
	l.entry().kv(WarnLevel, format, args...)
}

func (l *Logger) ErrorKv(format string, args ...interface{}) {
	l.entry().kv(ErrorLevel, format, args...)
}

func (l *Logger) FatalKv(format string, args ...interface{}) {
	l.entry().kv(FatalLevel, format, args...)
}

func (l *Logger) PanicKv(format string, args ...interface{}) {
	l.entry().kv(PanicLevel, format, args...)
}

func (l *Logger) TraceKv(format string, args ...interface{}) {
	l.entry().kv(TraceLevel, format, args...)
}

func (l *Logger) InfoKvln(format string, args ...interface{}) {
	l.entry().kvln(InfoLevel, format, args...)
}

func (l *Logger) DebugKvln(format string, args ...interface{}) {
	l.entry().kvln(DebugLevel, format, args...)
}

func (l *Logger) WarnKvln(format string, args ...interface{}) {
	l.entry().kvln(WarnLevel, format, args...)
}

func (l *Logger) ErrorKvln(format string, args ...interface{}) {
	l.entry().kvln(ErrorLevel, format, args...)
}

func (l *Logger) FatalKvln(format string, args ...interface{}) {
	l.entry().kvln(FatalLevel, format, args...)
}

func (l *Logger) PanicKvln(format string, args ...interface{}) {
	l.entry().kvln(PanicLevel, format, args...)
}

func (l *Logger) TraceKvln(format string, args ...interface{}) {
	l.entry().kvln(TraceLevel, format, args...)
}

func (l *Logger) Debugf(format string, args ...interface{}) {
	l.entry().log(DebugLevel, format, args...)
}

func (l *Logger) Infof(format string, args ...interface{}) {
	l.entry().log(InfoLevel, format, args...)
}

func (l *Logger) Warnf(format string, args ...interface{}) {
	l.entry().log(WarnLevel, format, args...)
}

func (l *Logger) Errorf(format string, args ...interface{}) {
	l.entry().log(ErrorLevel, format, args...)
}

func (l *Logger) Fatalf(format string, args ...interface{}) {
	l.entry().log(FatalLevel, format, args...)
}

func (l *Logger) Panicf(format string, args ...interface{}) {
	l.entry().log(PanicLevel, format, args...)
}

func (l *Logger) Tracef(format string, args ...interface{}) {
	l.entry().log(TraceLevel, format, args...)
}

func (l *Logger) Debug(args ...interface{}) {
	l.entry().log(DebugLevel, FmtEmptySeparate, args...)
}

func (l *Logger) Info(args ...interface{}) {
	l.entry().log(InfoLevel, FmtEmptySeparate, args...)
}

func (l *Logger) Warn(args ...interface{}) {
	l.entry().log(WarnLevel, FmtEmptySeparate, args...)
}

func (l *Logger) Error(args ...interface{}) {
	l.entry().log(ErrorLevel, FmtEmptySeparate, args...)
}

func (l *Logger) Fatal(args ...interface{}) {
	l.entry().log(FatalLevel, FmtEmptySeparate, args...)
}

func (l *Logger) Panic(args ...interface{}) {
	l.entry().log(PanicLevel, FmtEmptySeparate, args...)
}

func (l *Logger) Trace(args ...interface{}) {
	l.entry().log(TraceLevel, FmtEmptySeparate, args...)
}

func (l *Logger) Debugln(args ...interface{}) {
	l.entry().log(DebugLevel, FmtLineSeparate, args...)
}

func (l *Logger) Infoln(args ...interface{}) {
	l.entry().log(InfoLevel, FmtLineSeparate, args...)
}

func (l *Logger) Warnln(args ...interface{}) {
	l.entry().log(WarnLevel, FmtLineSeparate, args...)
}

func (l *Logger) Errorln(args ...interface{}) {
	l.entry().log(ErrorLevel, FmtLineSeparate, args...)
}

func (l *Logger) Fatalln(args ...interface{}) {
	l.entry().log(FatalLevel, FmtLineSeparate, args...)
}

func (l *Logger) Panicln(args ...interface{}) {
	l.entry().log(PanicLevel, FmtLineSeparate, args...)
}

func (l *Logger) Traceln(args ...interface{}) {
	l.entry().log(TraceLevel, FmtLineSeparate, args...)
}

// std logger

func InfoKv(format string, args ...interface{}) {
	stdLogger.entry().kv(InfoLevel, format, args...)
}

func DebugKv(format string, args ...interface{}) {
	stdLogger.entry().kv(DebugLevel, format, args...)
}

func WarnKv(format string, args ...interface{}) {
	stdLogger.entry().kv(WarnLevel, format, args...)
}

func ErrorKv(format string, args ...interface{}) {
	stdLogger.entry().kv(ErrorLevel, format, args...)
}

func FatalKv(format string, args ...interface{}) {
	stdLogger.entry().kv(FatalLevel, format, args...)
}

func PanicKv(format string, args ...interface{}) {
	stdLogger.entry().kv(PanicLevel, format, args...)
}

func TraceKv(format string, args ...interface{}) {
	stdLogger.entry().kv(TraceLevel, format, args...)
}

func InfoKvln(format string, args ...interface{}) {
	stdLogger.entry().kvln(InfoLevel, format, args...)
}

func DebugKvln(format string, args ...interface{}) {
	stdLogger.entry().kvln(DebugLevel, format, args...)
}

func WarnKvln(format string, args ...interface{}) {
	stdLogger.entry().kvln(WarnLevel, format, args...)
}

func ErrorKvln(format string, args ...interface{}) {
	stdLogger.entry().kvln(ErrorLevel, format, args...)
}

func FatalKvln(format string, args ...interface{}) {
	stdLogger.entry().kvln(FatalLevel, format, args...)
}

func PanicKvln(format string, args ...interface{}) {
	stdLogger.entry().kvln(PanicLevel, format, args...)
}

func TraceKvln(format string, args ...interface{}) {
	stdLogger.entry().kvln(TraceLevel, format, args...)
}

func Debugf(format string, args ...interface{}) {
	stdLogger.entry().log(DebugLevel, format, args...)
}

func Infof(format string, args ...interface{}) {
	stdLogger.entry().log(InfoLevel, format, args...)
}

func Warnf(format string, args ...interface{}) {
	stdLogger.entry().log(WarnLevel, format, args...)
}

func Errorf(format string, args ...interface{}) {
	stdLogger.entry().log(ErrorLevel, format, args...)
}

func Fatalf(format string, args ...interface{}) {
	stdLogger.entry().log(FatalLevel, format, args...)
}

func Panicf(format string, args ...interface{}) {
	stdLogger.entry().log(PanicLevel, format, args...)
}

func Tracef(format string, args ...interface{}) {
	stdLogger.entry().log(TraceLevel, format, args...)
}

func Debug(args ...interface{}) {
	stdLogger.entry().log(DebugLevel, FmtEmptySeparate, args...)
}

func Info(args ...interface{}) {
	stdLogger.entry().log(InfoLevel, FmtEmptySeparate, args...)
}

func Warn(args ...interface{}) {
	stdLogger.entry().log(WarnLevel, FmtEmptySeparate, args...)
}

func Error(args ...interface{}) {
	stdLogger.entry().log(ErrorLevel, FmtEmptySeparate, args...)
}

func Fatal(args ...interface{}) {
	stdLogger.entry().log(FatalLevel, FmtEmptySeparate, args...)
}

func Panic(args ...interface{}) {
	stdLogger.entry().log(PanicLevel, FmtEmptySeparate, args...)
}

func Trace(args ...interface{}) {
	stdLogger.entry().log(TraceLevel, FmtEmptySeparate, args...)
}

func Debugln(args ...interface{}) {
	stdLogger.entry().log(DebugLevel, FmtLineSeparate, args...)
}

func Infoln(args ...interface{}) {
	stdLogger.entry().log(InfoLevel, FmtLineSeparate, args...)
}

func Warnln(args ...interface{}) {
	stdLogger.entry().log(WarnLevel, FmtLineSeparate, args...)
}

func Errorln(args ...interface{}) {
	stdLogger.entry().log(ErrorLevel, FmtLineSeparate, args...)
}

func Fatalln(args ...interface{}) {
	stdLogger.entry().log(FatalLevel, FmtLineSeparate, args...)
}

func Panicln(args ...interface{}) {
	stdLogger.entry().log(PanicLevel, FmtLineSeparate, args...)
}

func Traceln(args ...interface{}) {
	stdLogger.entry().log(TraceLevel, FmtLineSeparate, args...)
}
