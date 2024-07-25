package main

import (
	"encoding/json"
	"fmt"
	"os"
)

func main() {
	structExample()
}

func exitOnError(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func structExample() {
	type location struct {
		Latitude  float64 `json:"lat"`
		Longitude float64 `json:"long"`
	}
	curiosity := location{-100.0, 50.0}
	bytes, err := json.Marshal(curiosity)
	exitOnError(err)
	fmt.Println(string(bytes))
}
