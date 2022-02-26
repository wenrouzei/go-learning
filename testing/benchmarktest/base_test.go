package benchmarktest

import (
	"bytes"
	"fmt"
	"strconv"
	"testing"
)

//go test -bench=^BenchmarkSlice -benchmem
//TODO silce or map make with size best
func BenchmarkSliceSize(b *testing.B) {
	b.ResetTimer()
	size := 10
	for i := 0; i < b.N; i++ {
		data := make([]int, 0, size)
		for n := 0; n < size; n++ {
			data = append(data, i)
		}
	}
}

func BenchmarkSliceNoneSize(b *testing.B) {
	b.ResetTimer()
	size := 10
	for i := 0; i < b.N; i++ {
		data := make([]int, 0)
		for n := 0; n < size; n++ {
			data = append(data, i)
		}
	}
}

//go test -bench=^BenchmarkInt -benchmem
//TODO strconv better than fmt
func BenchmarkIntToStringWithStrconv(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = strconv.Itoa(i)
	}
}

func BenchmarkIntToStringWithFmt(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = fmt.Sprint(i)
	}
}

//go test -bench=^BenchmarkStringConversion -benchmem -benchtime=10s
//TODO String conversion once better
func BenchmarkStringConversionOnce(b *testing.B) {
	b.ResetTimer()
	var buf bytes.Buffer
	p := []byte("hello word")
	for i := 0; i < b.N; i++ {
		buf.Write(p)
		buf.Reset()
	}
}

func BenchmarkStringConversionEverytime(b *testing.B) {
	b.ResetTimer()
	var buf bytes.Buffer
	for i := 0; i < b.N; i++ {
		buf.Write([]byte("hello word"))
		buf.Reset()
	}
}
