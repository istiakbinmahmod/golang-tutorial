package main

import (
	"fmt"
	"strings"
)

// dump slice length, capacity, and contents
func dump(label string, slice []string) {
	fmt.Printf("%v: length %v, capacity %v %v\n", label, len(slice), cap(slice), slice)
}

func main() {
	moves := []int{3, 2, 1, 0, 4}
	ans := canReachLastIndex(moves)
	fmt.Println(ans)
	dwarfs := []string{"Ceres", "Pluto", "Haumea", "Makemake", "Eris"}
	dump("dwarfs", dwarfs)
	dump("dwarfs[1:2]", dwarfs[1:2])
	dump("dwarfs[2:4]", dwarfs[2:4])
	sliceFunctions()
	preAllocateSlices()
}

func preAllocateSlices() {
	sliceObj := make([]string, 0)
	fmt.Printf("emptySlice: %3v, length: %3v, capacity: %3v\n", sliceObj, len(sliceObj), cap(sliceObj))
	sliceObj = append(sliceObj, "Istiak")
	fmt.Printf("emptySlice: %3v, length: %3v, capacity: %3v\n", sliceObj, len(sliceObj), cap(sliceObj))
	sliceObj = append(sliceObj, "Bin", "Mahmod")
	fmt.Printf("emptySlice: %3v, length: %3v, capacity: %3v\n", sliceObj, len(sliceObj), cap(sliceObj))
}

func canReachLastIndex(moves []int) bool {
	canReachHighest := 0
	for i, move := range moves {
		if i > canReachHighest {
			return false
		}
		canReachHighest = max(canReachHighest, i+move)
		if canReachHighest >= len(moves)-1 {
			return true
		}
	}
	return true
}

func sliceFunctions() {
	name := []string{"Istiak", "Bin"}
	strName := strings.Join(name, " ")
	fmt.Printf("name: %v\n", strName)
	fmt.Printf("capacity %v length %v\n", cap(name), len(name))
	name = append(name, "Mahmod")
	strName = strings.Join(name, " ")
	fmt.Printf("name: %v\n", strName)
	fmt.Printf("capacity %v length %v\n", cap(name), len(name))
}
