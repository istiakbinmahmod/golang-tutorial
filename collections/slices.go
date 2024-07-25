package main

import "fmt"

type siuu rune

func main() {
	names := [...]string{
		"Istiak",
		"Bin",
		"Mahmod",
		"Lionel",
		"Messi",
		"Cristiano",
		"Ronaldo",
	}

	firstName := names[0:3]
	secondName := names[3:5]
	thirdName := names[5:7]

	fmt.Printf("firstName: %v\n", firstName)
	fmt.Printf("secondName: %v\n", secondName)
	fmt.Printf("thirdName: %v\n", thirdName)

	thirdName[1] = "Penaldo"
	pr7 := names[5:7]
	fmt.Printf("pr7: %v\n", pr7)

	fmt.Printf("Type of names : %T\n", names)
	fmt.Printf("Type of pr7 : %T\n", pr7)

	lm10 := []string{
		"10",
		"LaLiga",
		"8",
		"Ballon D'or",
		"None",
		"Stolen",
	}
	fmt.Printf("lm10: %v\n", lm10)
	fmt.Printf("lm10 type: %T\n", lm10)

	nameStr := "Istiak Bin Mahmod"
	lastName := nameStr[11:17]
	fmt.Printf("lastName: %v\n", lastName)

	nameStr = "¿Cómo estás?"
	fmt.Printf("nameStr[:5]: %v\n", nameStr[:5])
}
