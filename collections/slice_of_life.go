package main

import (
	"fmt"
	"math/rand"
	"time"
)

const (
	width  = 80
	height = 25
)

type Universe [][]bool

func NewUniverse() Universe {
	u := make(Universe, height)
	for i := range u {
		u[i] = make([]bool, width)
	}
	return u
}

func (u Universe) Show() {
	for _, row := range u {
		for _, cell := range row {
			if cell == true {
				fmt.Printf("%v ", "*")
			} else {
				fmt.Printf("%v ", " ")
			}
		}
		fmt.Println()
	}
}

func (u Universe) Seed() {
	for i, row := range u {
		for j := range row {
			// Set approximately 25% cells to alive
			randInt := rand.Intn(4)
			if randInt == 0 {
				u[i][j] = true
			} else {
				u[i][j] = false
			}
		}
	}
}

func (u Universe) Alive(x, y int) bool {
	return u[(y+height)%height][(x+width)%width]
}

func (u Universe) LiveNeighbors(x, y int) int {
	aliveCount := 0
	comb := []int{-1, 0, 1}
	for _, i := range comb {
		for _, j := range comb {
			// skip own cell i.e. (x,y)
			if i != 0 || j != 0 {
				if u.Alive(x+i, y+j) {
					aliveCount++
				}
			}
		}
	}
	return aliveCount
}

func (u Universe) Next(x, y int) bool {
	aliveNeighbors := u.LiveNeighbors(x, y)
	if u.Alive(x, y) {
		return aliveNeighbors == 2 || aliveNeighbors == 3
	} else {
		return aliveNeighbors == 3
	}
}

func (u Universe) Step(next Universe) {
	for i, row := range u {
		for j := range row {
			next[i][j] = u.Next(i, j)
		}
	}
}

/*
Function to check whether universe is converged or not

if each alive cell is alive in next iteration, and dead cell dead in next iteration,
it will recursively be the same in all future iterations, hence become converged

Return true if converges, else return false
*/
func (u Universe) ConvergeCheck() bool {
	for i, row := range u {
		for j, cell := range row {
			if cell != u.Next(i, j) {
				return false
			}
		}
	}
	return true
}

func main() {
	myUniverse := NewUniverse()
	nextUniverse := NewUniverse()
	myUniverse.Seed()
	for true {
		fmt.Print("\033[H")
		myUniverse.Show()
		time.Sleep(100 * time.Millisecond) // refresh screen after every 100 millisecond
		myUniverse.Step(nextUniverse)
		// Converges after some time. It's a beautiful thing!
		if myUniverse.ConvergeCheck() {
			fmt.Println("Converged to this^ shape!!!")
			break
		}
		myUniverse = nextUniverse
	}
}
