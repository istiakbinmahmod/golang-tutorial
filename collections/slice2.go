package main

import (
	"fmt"
	"sort"
	"strings"
)

type StringSlice []string

func hyperspace(worlds StringSlice) {
	for i := range worlds {
		worlds[i] = strings.TrimSpace(worlds[i])
	}
}

func main() {
	planets := StringSlice{
		"Venus", " ", "Earth", " ", "Mars",
	}
	hyperspace(planets)
	fmt.Println(strings.Join(planets, ""))
	sort.StringSlice(planets).Sort()
	fmt.Println(strings.Join(planets, ","))
	planetsConcatenated := strings.Join(planets, "")
	fmt.Printf("planetsConcatenated: %v\n", planetsConcatenated)
	sort.Strings(planets)
}
