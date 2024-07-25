package main

import (
	"fmt"
	"math/rand"
)

type WumpusWorld [][]int
type DIRECTION int

type Position struct {
	x, y int
}

const (
	LENGTH int = 4
	WIDTH  int = 4
)

const (
	DIRECTION_LEFT  DIRECTION = 0
	DIRECTION_UP    DIRECTION = 1
	DIRECTION_RIGHT DIRECTION = 2
	DIRECTION_BELOW DIRECTION = 3
)

const (
	WUMPUS int = 1
	PIT    int = 2
	GOLD   int = 3
)

func RandomCoordinate() Position {
	rand1, rand2 := rand.Intn(LENGTH), rand.Intn(WIDTH)
	return Position{x: rand1, y: rand2}
}

func (p Position) Equal(p1 Position) bool {
	return p.x == p1.x && p.y == p1.y
}

func NewWumpusWorld() WumpusWorld {
	w := make(WumpusWorld, LENGTH)
	for i := range w {
		w[i] = make([]int, WIDTH)
	}
	// initialize Wumpus, Pit, and Gold
	var wumpusPosition, pitPosition, goldPosition Position
	wumpusPosition = RandomCoordinate()
	for true {
		pitPosition = RandomCoordinate()
		if !pitPosition.Equal(wumpusPosition) {
			break
		}
	}
	for true {
		goldPosition = RandomCoordinate()
		if !goldPosition.Equal(wumpusPosition) && !goldPosition.Equal(pitPosition) {
			break
		}
	}
	w[wumpusPosition.x][wumpusPosition.y] = WUMPUS
	w[pitPosition.x][pitPosition.y] = PIT
	w[goldPosition.x][goldPosition.y] = GOLD
	return w
}

func (d DIRECTION) shiftRight() DIRECTION {
	return (d + 1) % 4
}

func (d DIRECTION) shiftLeft() DIRECTION {
	return (d + 3) % 4
}

func (w WumpusWorld) shootArrow(p Position, d DIRECTION) {
	switch d {
	case DIRECTION_LEFT:
		return
	case DIRECTION_UP:
		return
	case DIRECTION_RIGHT:
		return
	case DIRECTION_BELOW:
		return
	}
}

func main() {
	fmt.Println("Hello")
}
