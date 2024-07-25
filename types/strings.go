package main

import (
	"fmt"
	"strings"
	"unicode/utf8"
)

func main() {
	// codePoint()
	// stringOps()
	stringEncoding()
	// caesarDecipher()
}

func caesarDecipher() {
	str := "L fdph, L vdz, L frqtxhuhg."
	for _, c := range str {
		if c >= 'a' && c <= 'z' {
			c = c - 3
			if c < 'a' {
				c += 'a'
			}
		} else if c >= 'A' && c <= 'Z' {
			c = c - 3
			if c < 'A' {
				c += 'A'
			}
		}
		fmt.Printf("%c", c)
	}
}

func stringEncoding() {
	// siuuti := "tumi ar kedo na"
	// for i, c := range siuuti {
	// 	fmt.Printf("%v %c %[2]v\n", i, c)
	// }

	question := "¿Cómo estás?"
	fmt.Println(len(question), "bytes")

	fmt.Println(utf8.RuneCountInString(question), "runes")

	c, size := utf8.DecodeRuneInString(question)
	fmt.Printf("First rune: %c, %v bytes", c, size)
}

func stringOps() {
	peace := "was never an option"
	ind6 := peace[6]
	fmt.Printf("%c\n", ind6)
	// below code gives an error since strings in go are immutable
	// peace[6] = 'c'
	peace = "can be an option"
	fmt.Println(peace)
	peace = "peace " + peace
	peace += " too"
	fmt.Println(peace)
	splittedStrings := strings.Split(peace, " ")
	fmt.Println(splittedStrings)
	for i := 0; i < len(splittedStrings)-1; i++ {
		fmt.Println(splittedStrings[i])
	}
}

func codePoint() {
	var codePoint rune = 65
	fmt.Printf("%c %[1]T\n", codePoint)
	siuu := 'π'
	fmt.Printf("%c %[1]v %[1]T\n", siuu)
	var star byte = '*'
	fmt.Printf("%c %[1]v %[1]T\n", star)
	g := 'g'
	g = g - 'a' + 'A'
	fmt.Printf("g: %c\n", g)
}
