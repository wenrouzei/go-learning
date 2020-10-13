package main

import (
	"fmt"
	"io"
	"log"
	"math"
	"net/http"
	"os"
	"strconv"
	"sync"
)

var wg sync.WaitGroup

type part struct {
	i     int
	start int
	end   int64
}

var ch chan *part

func main() {
	url := "http://down-lhlt.resources.3737.com/185_100216_299999.apk"
	ch = make(chan *part)
	size := 50 << 20
	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatal(err)
	}
	client := http.Client{}
	res, err := client.Do(request)
	if err != nil {
		log.Fatal(err)
	}
	res.Body.Close()
	num := res.Header.Get("Content-Length")
	len, _ := strconv.ParseInt(num, 10, 64)
	partNum := int(math.Ceil(float64(len) / float64(size)))
	fmt.Println(num, len, partNum, size)
	go func() {
		for i := 0; i < partNum; i++ {
			start := i * size
			var end int64
			if i == partNum-1 {
				end = len - 1
			} else {
				end = int64(start + size - 1)
			}
			part := part{
				i:     i,
				start: start,
				end:   end,
			}
			ch <- &part
		}
	}()

	for i := 0; i < 3; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for p := range ch {
				request.Header.Set("Range", fmt.Sprintf("bytes=%d-%d", p.start, p.end))
				res, _ := client.Do(request)
				defer res.Body.Close()
				file := fmt.Sprintf("%d", p.i)
				f, _ := os.Create(file)
				defer f.Close()
				io.Copy(f, res.Body)
			}
		}()
	}
	wg.Wait()

}
