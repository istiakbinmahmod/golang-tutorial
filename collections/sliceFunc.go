package main

import "fmt"

type StringSlice []string

func (planets StringSlice) terraform(prefix string) {
	for i := range planets {
		planets[i] = prefix + planets[i]
	}
}

func main() {
	planets := StringSlice{"Mars", "Uranus", "Neptune"}
	planets.terraform("New ")
	for i, planet := range planets {
		fmt.Printf("planet %v: %v\n", i+1, planet)
	}
}
