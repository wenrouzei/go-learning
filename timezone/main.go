package main

import (
	"fmt"
	"time"
)

func main() {
	eight := time.FixedZone("CST", 8*3600)
	night := time.FixedZone("CST", 9*3600)
	fmt.Println(time.Now().In(eight).Format("2006-01-02 15:04:05"), time.Now().In(night).Format("2006-01-02 15:04:05"))
}
