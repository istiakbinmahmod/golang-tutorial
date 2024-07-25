package main

import (
	"fmt"
	"strconv"
)

func main() {
	elems := []string{}
	loopCount := 500
	for i := 0; i < loopCount; i++ {
		fmt.Printf("Length: %5v and Capacity %5v\n", len(elems), cap(elems))
		elems = append(elems, strconv.Itoa(i))
	}
}
