package main

import (
	"fmt"
	"strings"
	"unicode"
)

func main() {
	fmt.Println(strings.Contains("seafood", "foo")) // true
	fmt.Println(strings.Contains("seafood", "bar")) // false
	fmt.Println(strings.Contains("seafood", ""))    // true
	fmt.Println(strings.Contains("", ""))           // true
	fmt.Println("##########################################")
	fmt.Println(strings.ContainsAny("team", "i"))     // false
	fmt.Println(strings.ContainsAny("fail", "ui"))    // true
	fmt.Println(strings.ContainsAny("ure", "ui"))     // true
	fmt.Println(strings.ContainsAny("failure", "ui")) // true
	fmt.Println(strings.ContainsAny("foo", ""))       // false
	fmt.Println(strings.ContainsAny("", ""))          // false
	fmt.Println("##########################################")
	// Finds whether a string contains a particular Unicode code point.
	// The code point for the lowercase letter "a", for example, is 97.
	fmt.Println(strings.ContainsRune("aardvark", 97))
	fmt.Println(strings.ContainsRune("timeout", 97))
	fmt.Println("##########################################")
	fmt.Println(strings.Count("cheese", "e")) // 3
	fmt.Println(strings.Count("five", ""))    // before & after each rune 5=4+1
	fmt.Println("##########################################")
	fmt.Printf("Fields are: %q", strings.Fields("  foo bar  baz   ")) //  ["foo" "bar" "baz"]
	f := func(c rune) bool {
		return !unicode.IsLetter(c) && !unicode.IsNumber(c)
	}
	fmt.Printf("Fields are: %q", strings.FieldsFunc("  foo1;bar2,baz3...", f)) // ["foo1" "bar2" "baz3"]
	fmt.Println("##########################################")
	fmt.Println(strings.HasPrefix("Gopher", "Go")) // true
	fmt.Println(strings.HasPrefix("Gopher", "C"))  // false
	fmt.Println(strings.HasPrefix("Gopher", ""))   // true

	fmt.Println(strings.HasSuffix("Amigo", "go"))  // true
	fmt.Println(strings.HasSuffix("Amigo", "O"))   // false
	fmt.Println(strings.HasSuffix("Amigo", "Ami")) // false
	fmt.Println(strings.HasSuffix("Amigo", ""))    // true
	fmt.Println("##########################################")
	s := []string{"foo", "bar", "baz"}
	fmt.Println(strings.Join(s, ", ")) // foo,bar,baz
	fmt.Println("##########################################")
	fmt.Println(strings.Index("chicken", "ken")) // 4
	fmt.Println(strings.Index("chicken", "dmr")) // -1

	// 子串中的任意字符在源串出现的位置
	fmt.Println(strings.IndexAny("chicken", "aeiouy")) // 2
	fmt.Println(strings.IndexAny("crwth", "aeiouy"))   // -1

	// IndexByte，字符在字符串中出现的位置
	fmt.Println(strings.IndexByte("golang", 'g'))  // 0
	fmt.Println(strings.IndexByte("gophers", 'h')) // 3
	fmt.Println(strings.IndexByte("golang", 'x'))  // -1

	// IndexFunc 满足条件的作为筛选条件
	f1 := func(c rune) bool {
		return unicode.Is(unicode.Han, c)
	}
	fmt.Println(strings.IndexFunc("Hello, 世界", f1))    // 7
	fmt.Println(strings.IndexFunc("Hello, world", f1)) // -1

	// 某个字符在源串中的位置
	fmt.Println(strings.IndexRune("chicken", 'k')) // 4
	fmt.Println(strings.IndexRune("chicken", 'd')) // -1
	fmt.Println("##########################################")
	rot13 := func(r rune) rune { // r是遍历的每一个字符
		switch {
		case r >= 'A' && r <= 'Z':
			return 'A' + (r-'A'+13)%26
		case r >= 'a' && r <= 'z':
			return 'a' + (r-'a'+13)%26
		}
		return r
	}
	fmt.Println(strings.Map(rot13, "'Twas brillig and the slithy gopher..."))

	fmt.Println("##########################################")
	fmt.Println("ba" + strings.Repeat("na", 2))
	fmt.Println("##########################################")
	fmt.Println(strings.Replace("oink oink oink", "k", "ky", 2)) // oinky oinkky oink
	fmt.Println(strings.Replace("oink oink oink", "oink", "moo", -1))
	fmt.Println("##########################################")

	fmt.Printf("%q\n", strings.Split("a,b,c", ","))                        // ["a","b","c"]
	fmt.Printf("%q\n", strings.Split("a man a plan a canal panama", "a ")) // ["" "man " "plan " "canal panama"]
	fmt.Printf("%q\n", strings.Split(" xyz ", ""))                         // [" " "x" "y" "z" " "]
	fmt.Printf("%q\n", strings.Split("", "Bernardo O'Higgins"))            // [""]

	// SplitN 定义返回之后的切片中包含的长度，最后一部分是未被处理的。
	fmt.Printf("%q\n", strings.SplitN("a,b,c", ",", 2)) // ["a", "b,c"]
	z := strings.SplitN("a,b,c", ",", 0)
	fmt.Printf("%q (nil = %v)\n", z, z == nil) // [] (nil = true)

	// 使用sep分割，分割出来的字符串中包含sep，可以限定分割之后返回的长度。
	fmt.Printf("%q\n", strings.SplitAfterN("a,b,c", ",", 2)) // ["a,", "b,c"]

	// 完全分割
	fmt.Printf("%q\n", strings.SplitAfter("a,b,c", ",")) // ["a,","b,", "c"]
	fmt.Println("##########################################")
	// Trim 包含在cutset中的元素都会被去掉
	fmt.Print(strings.Trim("¡¡¡Hello, Gophers!!!", "!¡")) // Hello, Gophers

	// TrimFunc去掉满足条件的字符
	fmt.Print(strings.TrimFunc("¡¡¡Hello, Gophers!!!", func(r rune) bool {
		return !unicode.IsLetter(r) && !unicode.IsNumber(r)
	}))

	// TrimLeft 去掉左边满足包含在cutset中的元素，直到遇到不在cutset中的元素为止
	fmt.Print(strings.TrimLeft("¡¡¡Hello, Gophers!!!", "!¡")) // Hello, Gophers!!!

	// TrimLeftFunc 去掉左边属于函数返回值部分，直到遇到不在cutset中的元素为止
	fmt.Print(strings.TrimLeftFunc("¡¡¡Hello, Gophers!!!", func(r rune) bool {
		return !unicode.IsLetter(r) && !unicode.IsNumber(r)
	})) // Hello, Gophers!!!

	// TrimPrefix 去掉开头部分；TrimSuffix 去掉结尾部分
	var s1 = "¡¡¡Hello, Gophers!!!"
	s1 = strings.TrimPrefix(s1, "¡¡¡Hello, ")
	s1 = strings.TrimPrefix(s1, "¡¡¡Howdy, ")
	fmt.Println(s1)
	fmt.Println("##########################################")
	var b strings.Builder
	for i := 3; i >= 1; i-- {
		fmt.Fprintf(&b, "%d...", i)
		fmt.Fprint(&b, i)
	}
	b.WriteString("ignition")
	fmt.Println(b.String())
	fmt.Println("##########################################")
	r := strings.NewReader("中国")
	fmt.Println(r.Len()) // 输出14  初始时，未读长度等于字符串长度
	var buf []byte
	buf = make([]byte, 5)
	readLen, err := r.Read(buf)
	fmt.Println("读取到的长度:", readLen) //读取到的长度5
	if err != nil {
		fmt.Println("错误:", err)
	}
	fmt.Println(buf)      //adcde
	fmt.Println(r.Len())  //9   读取到了5个 剩余未读是14-5
	fmt.Println(r.Size()) //14   字符串的长度

}
