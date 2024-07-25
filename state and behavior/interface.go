package main

import (
	"fmt"
	"strings"
)

/*
	Used together, composition and interfaces make a very powerful tool.
*/

type talker interface {
	talk() string
}

type martian struct{}

func (m martian) talk() string {
	return "peck peck"
}

func (m martian) siuu() string {
	return "martian siuuuti"
}

type laser int

func (l laser) talk() string {
	return strings.Repeat("siuuuti ", int(l))
}

type istiak struct {
	laser
}

func shout(t talker) {
	louder := strings.ToUpper(t.talk())
	fmt.Printf("louder: %v\n", louder)
}

func main() {
	var t talker
	t = martian{}
	fmt.Println(t.talk())
	t = laser(3)
	fmt.Println(t.talk())
	shout(t)
	shout(martian{})
	shout(laser(4))
	shout(istiak{laser: laser(5)})
}
