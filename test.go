package main

import (
	"bufio"
	"crypto/md5"
	"errors"
	"fmt"
	"io"
	"os"
	"reflect"
	"runtime"
	"strconv"
	"strings"
	"sync"
	"time"
	"unsafe"
)

type Saiyan struct {
	Name   string
	Power  int
	Father *Saiyan
}

func (s *Saiyan) Super() {
	s.Power += 10000
}

func NewSaiyan(name string, power int) Saiyan {
	return Saiyan{
		Name:  name,
		Power: power,
		Father: &Saiyan{
			Name:   "Goku",
			Power:  9001,
			Father: nil,
		},
	}
}

type contain struct {
	Saiyan
	test string
}

type aaa interface {
	echo()
}

type aa string

func (a aa) echo() {
	fmt.Println(a)
}

func testEcho() int {
	defer func() {
		fmt.Println("defer testEcho")
	}()
	return 12
}

func main() {
	//goku := &Saiyan{"Goku", 9001, &Saiyan{"111", 232, nil}}
	//goku.Super()
	//fmt.Println(goku.Father) // 将会打印出 19001
	//a := NewSaiyan("abc", 1000)
	//b := NewSaiyan("abcd", 2000)
	//fmt.Println(a.Name)
	//fmt.Println(b.Name)
	c := 123
	fmt.Println(uint64(c))

	d := &c

	a := contain{Saiyan{"abc", 24323, nil}, "fasdfas"}
	a.Super()
	fmt.Println(a.Name, a.Saiyan)

	defer func() {
		fmt.Println("abc")
	}()

	fmt.Println("11111")

	fmt.Println(reflect.TypeOf(c), reflect.TypeOf(d))

	fmt.Println(unsafe.Sizeof(c), unsafe.Sizeof(d))

	var first aaa
	var second aa = "abc"
	first = second
	first.echo()

	h := md5.New()
	fmt.Println(h)
	io.WriteString(h, strconv.FormatInt(time.Now().Unix(), 10))
	fmt.Println(h)
	io.WriteString(h, "ganraomaxxxxxxxxx")
	token := fmt.Sprintf("%x", h.Sum(nil))
	fmt.Println(token, h)

	fmt.Println(32 << 20)

	str, _ := os.Getwd()
	fmt.Println(str)

	fmt.Println(testEcho())

	fmt.Println(reflect.TypeOf(fmt.Errorf("%w", errors.New("error"))))

	fmt.Println(strings.Index("abc", "a"))

	t := time.Now().UTC()
	fmt.Println(t)          // Wed Dec 21 08:52:14 +0000 UTC 2011
	fmt.Println(time.Now()) // Wed Dec 21 09:52:14 +0100 RST 2011

	fmt.Println(deferT())

	wg := sync.WaitGroup{}
	b := make(map[int]int)
	for i:= 0;i<=100; i++{
		wg.Add(1)
		go func(wg *sync.WaitGroup, i int, b map[int]int) {
			fmt.Println(fmt.Sprintf("%p", b)  ,1111)
			defer wg.Done()
			fmt.Println(i)

			//fmt.Println(b)
		}(&wg, i, b)
	}
	wg.Wait()
	fmt.Println(b)
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	fmt.Printf("%d Kb\n", m.Alloc / 1024)
	aa:=33
	runtime.SetFinalizer(&aa, func(aa *int){})
	fmt.Printf("y/n?")
	reader := bufio.NewReader(os.Stdin)
	data,_,_ := reader.ReadLine()
	fmt.Printf("%v", string(data))
}

func deferT() int {
	fmt.Println("defer1")
	defer func() {
		fmt.Println("defer3")
	}()
	defer fmt.Println("defer4")
	return deferS()
}

func deferS() (x int) {
	fmt.Println("defer2")
	x = 0
	return
}
