package main

import (
	"fmt"
	"time"
)

func main() {
	idx := 0

	for {
		fmt.Printf("%d, Hello World!\n", idx)
		idx ++
		time.Sleep(time.Second)
	}
}
