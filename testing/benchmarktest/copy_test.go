package benchmarktest

import (
	"fmt"
	"sync"
	"testing"
)

//go test -bench=^BenchmarkOnly -benchmem -benchtime=10s

type aa struct {
	A string
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
				fmt.Println(i)
				return
			}
		}
	}
}

func bbChOut(wg *sync.WaitGroup, bbCh <-chan *aa) {
	defer wg.Done()
	var i int
	for {
		select {
		case _, ok := <-bbCh:
			if ok {
				i++
			} else {
				fmt.Println(i)
				return
			}
		}
	}
}

func BenchmarkOnlyCopy(b *testing.B) {
	b.ResetTimer()
	var aaCh = make(chan aa, 100)
	var wg, wg1 sync.WaitGroup
	wg.Add(1)
	go aaChOut(&wg, aaCh)
	for i := 0; i < b.N; i++ {
		wg1.Add(1)
		go func() {
			defer wg1.Done()
			cc := aa{
				A: "a",
			}
			aaCh <- cc
		}()
	}
	wg1.Wait()
	close(aaCh)
	wg.Wait()
}

func BenchmarkOnlyPoint(b *testing.B) {
	b.ResetTimer()
	var bbCh = make(chan *aa, 100)
	var wg, wg1 sync.WaitGroup
	wg.Add(1)
	go bbChOut(&wg, bbCh)
	for i := 0; i < b.N; i++ {
		wg1.Add(1)
		go func() {
			defer wg1.Done()
			cc := aa{
				A: "a",
			}
			bbCh <- &cc
		}()
	}
	wg1.Wait()
	close(bbCh)
	wg.Wait()
}
