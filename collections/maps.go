package main

import (
	"fmt"
	"strings"
)

func main() {
	// mapExample1()
	mapExample2()

}

func mapExample2() {
	names := []string{
		"Istiak", "Messi", "Penaldo", "MoneyBappe", "Siuu", "Istiak", "Penaldo",
	}
	nameMap := make(map[string]int)
	namesByFirstAlphabet := make(map[uint8][]string)
	for _, name := range names {
		nameMap[name]++
		namesByFirstAlphabet[name[0]] = append(namesByFirstAlphabet[name[0]], name)
	}
	for t, count := range nameMap {
		formattedName := "`" + t + "`"
		fmt.Printf("Name: %-10v ocurred %2v times\n", formattedName, count)
	}
	for t, _names := range namesByFirstAlphabet {
		fmt.Printf("Key: %v and Value: %-10v\n", string(t), strings.Join(_names, ","))
	}
	fmt.Printf("Length of namesByFirstAlphabet: %v\n", len(namesByFirstAlphabet))
}

func mapExample1() {
	citiesByCountry := map[string]string{
		"United States": "Washington DC",
		"Bangladesh":    "Dhaka",
		"Pakistan":      "Islamabad",
	}
	delete(citiesByCountry, "Bangladesh")
	country := "United States"
	if cap, ok := citiesByCountry[country]; ok {
		fmt.Printf("capital: %v\n", cap)
	} else {
		fmt.Printf("capital for %v couldn't be found\n", country)
	}
	// citiesByCountry["Bangladesh"] = "Dhaka"
	country = "Bangladesh"
	if cap, ok := citiesByCountry[country]; ok {
		fmt.Printf("capital: %v\n", cap)
	} else {
		fmt.Printf("capital for %v couldn't be found\n", country)
	}
	country = "India"
	if cap, ok := citiesByCountry[country]; ok {
		fmt.Printf("capital: %v\n", cap)
	} else {
		fmt.Printf("capital for %v couldn't be found\n", country)
	}
}
