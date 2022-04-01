package benchmarksyncpool

import (
	"fmt"
	"runtime"
	"sync"
	"testing"
)

type BigDataStructure struct {
	ID     string
	Name   string
	Length int
}

var bigDataPool = &sync.Pool{
	New: func() interface{} {
		return new(BigDataStructure)
	},
}

//go test -bench=^Benchmark -benchmem

func BenchmarkWithoutPool(b *testing.B) {
	var s *BigDataStructure
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for j := 0; j < 10000; j++ {
			s = &BigDataStructure{}
			s.ID = "1"
			s.Name = "Item-1"
			s.Length = j
		}
	}
}

func BenchmarkWithPool(b *testing.B) {
	fmt.Println(runtime.NumCPU())
	runtime.GOMAXPROCS(2)
	var s *BigDataStructure
	b.ReportAllocs()
	b.ResetTimer()
	pm := make(map[string]struct{})
	for i := 0; i < b.N; i++ {
		for j := 0; j < 10; j++ {
			s = bigDataPool.Get().(*BigDataStructure)
			pm[fmt.Sprintf("%p", s)] = struct{}{}
			s.ID = "1"
			s.Name = "Item-1"
			s.Length = j
			bigDataPool.Put(s)
		}
	}
	fmt.Println(pm)
}

func BenchmarkWithPoolGC(b *testing.B) {
	var s *BigDataStructure
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for j := 0; j < 10000; j++ {
			s = bigDataPool.Get().(*BigDataStructure)
			s.ID = "1"
			s.Name = "Item-1"
			s.Length = j
			bigDataPool.Put(s)
		}
		runtime.GC()
	}
}
