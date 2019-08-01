package log

import (
	"log"
	"os"
	"runtime"
	"testing"
)

var textLogger, jsonLogger *Logger

func init() {

	// log file
	file, _ := os.OpenFile("D:/test.log", os.O_TRUNC|os.O_CREATE|os.O_WRONLY, os.ModeAppend)
	runtime.SetFinalizer(file, func(f *os.File) {
		_ = f.Close()
	})

	// text logger
	textLogger = New(
		WithLevel(TraceLevel),
		WithStdLevel(TraceLevel),
		WithOutput(file),
		WithFileLine(true),
		WithFormatter(&TextFormatter{IgnoreBasicFields: false}),
	)

	// json logger
	jsonLogger = New(
		WithLevel(TraceLevel),
		WithStdLevel(TraceLevel),
		WithOutput(file),
		WithFileLine(true),
		WithFormatter(&JsonFormatter{IgnoreBasicFields: true}),
	)

	// overwrite the go std lib log
	log.SetFlags(0)
	log.SetOutput(jsonLogger.Writer())
}

// go test -test.bench BenchmarkLogger_TextInfof -test.count=1 -test.benchtime=1s -test.benchmem
// BenchmarkLogger_TextInfof-4       500000              3596 ns/op             248 B/op          4 allocs/op
func BenchmarkLogger_TextInfof(b *testing.B) {
	for i := 0; i < b.N; i++ {
		textLogger.Infof("BenchmarkLogger_TextInfof \n")
	}
}

// go test -test.bench BenchmarkLogger_TextKvFmt -test.count=1 -test.benchtime=1s -test.benchmem
// BenchmarkLogger_TextKvFmt-4       200000              5485 ns/op             216 B/op          3 allocs/op
func BenchmarkLogger_TextKvFmt(b *testing.B) {
	for i := 0; i < b.N; i++ {
		textLogger.InfoKvln("aaa,bbb,ccc,ddd,eee,fff,ggg,hhh", 1, 2, 3, 4, 5, 6, 7, 8)
	}
}

// go test -test.bench BenchmarkLogger_JsonInfof -test.count=1 -test.benchtime=1s -test.benchmem
// BenchmarkLogger_JsonInfof-4       200000              5980 ns/op            1144 B/op         17 allocs/op
func BenchmarkLogger_JsonInfof(b *testing.B) {
	for i := 0; i < b.N; i++ {
		jsonLogger.Infof("BenchmarkLogger_JsonInfof %d\n", i)
	}
}

// go test -test.bench BenchmarkLogger_JsonKvFmt -test.count=1 -test.benchtime=1s -test.benchmem
// BenchmarkLogger_JsonKvFmt-4       200000              7540 ns/op            1080 B/op         14 allocs/op
func BenchmarkLogger_JsonKvFmt(b *testing.B) {
	for i := 0; i < b.N; i++ {
		jsonLogger.InfoKvln("aaa,bbb,ccc,ddd,eee,fff,ggg,hhh", 1, 2, 3, 4, 5, 6, 7, 8)
	}
}

// go test -test.bench BenchmarkLogger_TextWriter -test.count=1 -test.benchtime=1s -test.benchmem
// BenchmarkLogger_TextWriter-4      300000              3983 ns/op             320 B/op          8 allocs/op
func BenchmarkLogger_TextWriter(b *testing.B) {
	for i := 0; i < b.N; i++ {
		log.Println("BenchmarkLogger_Writer")
	}
}

// go test -test.bench BenchmarkLogger_JsonWriter -test.count=1 -test.benchtime=1s -test.benchmem
// BenchmarkLogger_JsonWriter-4      200000              6185 ns/op            1184 B/op         20 allocs/op
func BenchmarkLogger_JsonWriter(b *testing.B) {
	for i := 0; i < b.N; i++ {
		log.Println("BenchmarkLogger_Writer")
	}
}
