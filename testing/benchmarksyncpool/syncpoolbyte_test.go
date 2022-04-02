package benchmarksyncpool

import (
	"runtime"
	"sync"
	"sync/atomic"
	"testing"
)

var bytePool = &sync.Pool{
	New: func() interface{} {
		atomic.AddUint32(&times, 1)
		return make([]byte, 0, 255)
	},
}
var times uint32

//go test -bench=^BenchmarkByte -benchmem

func BenchmarkByteWithoutPool(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		s := make([]byte, 0, 255)
		s = append(s, byte('a'))
	}
}

func BenchmarkByteWithPool(b *testing.B) {
	//fmt.Println(runtime.NumCPU())
	//runtime.GOMAXPROCS(1)
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		s := bytePool.Get().([]byte)
		s = append(s, byte('a'))
		bytePool.Put(s)
	}
	//fmt.Println(times, b.N, runtime.NumCPU())
}

func BenchmarkByteWithPoolGC(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		s := bytePool.Get().([]byte)
		s = append(s, byte('a'))
		bytePool.Put(s)
		runtime.GC()
	}
	//fmt.Println(times, b.N, runtime.NumCPU())
}
