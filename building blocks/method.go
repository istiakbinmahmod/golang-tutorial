package main

import "fmt"

type celsius float64

type songType string

func main() {
	var song songType = "siuuuti"
	fmt.Println(song.completeSong())
}

func (s songType) completeSong() songType {
	return s + songType(" tumi ar kedo na")
}

func temp() {
	var temperature celsius = 20
	fmt.Printf("temperature: %v\n", temperature)
	warmup := 10
	temperature += celsius(warmup)
	fmt.Printf("new temperature: %v\n", temperature)
}
