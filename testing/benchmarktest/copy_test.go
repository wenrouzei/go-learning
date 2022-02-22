package benchmarktest

import (
	"sync"
	"testing"
	"time"
)

//go test -bench=^BenchmarkOnly -benchmem -benchtime=10s

type aa struct {
	A string
	B time.Time
}

func aaChOut(wg *sync.WaitGroup, aaCh <-chan aa) {
	defer wg.Done()
	var i int
	for {
		select {
		case _, ok := <-aaCh:
			if ok {
				i++
			} else {
				//fmt.Println(i)
				return
			}
		}
	}
}

func bbChOut(wg *sync.WaitGroup, aaCh <-chan *aa) {
	defer wg.Done()
	var i int
	for {
		select {
		case _, ok := <-aaCh:
			if ok {
				i++
			} else {
				//fmt.Println(i)
				return
			}
		}
	}
}

func BenchmarkOnlyCopy(b *testing.B) {
	b.ResetTimer()
	var aaCh = make(chan aa, 100)
	var wg sync.WaitGroup
	wg.Add(1)
	go aaChOut(&wg, aaCh)
	for i := 0; i < b.N; i++ {
		cc := aa{
			A: "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaabbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbb",
			B: time.Now(),
		}
		aaCh <- cc
	}
	close(aaCh)
	wg.Wait()
}

func BenchmarkOnlyPointer(b *testing.B) {
	b.ResetTimer()
	var aaCh = make(chan *aa, 100)
	var wg sync.WaitGroup
	wg.Add(1)
	go bbChOut(&wg, aaCh)
	for i := 0; i < b.N; i++ {
		cc := aa{
			A: "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaabbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbb",
			B: time.Now(),
		}
		aaCh <- &cc
	}
	close(aaCh)
	wg.Wait()
}
