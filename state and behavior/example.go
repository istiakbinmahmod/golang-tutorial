package main

import (
	"fmt"
	"math"
)

type World struct {
	radius float64
}

type Location struct {
	name      string
	lat, long float64
}

type GPS struct {
	current     Location
	destination Location
	world       World
}

type Rover struct {
	GPS
}

func rad(deg float64) float64 {
	return deg * math.Pi / 180
}

func (w World) distance(l1, l2 Location) float64 {
	s1, c1 := math.Sincos(rad(l1.lat))
	s2, c2 := math.Sincos(rad(l2.lat))
	clong := math.Cos(rad(l1.long - l2.long))
	return w.radius * math.Acos(s1*s2+c1*c2*clong)
}

func (l Location) description() string {
	description := "Location : " +
		l.name +
		" \n" +
		"latitude : " +
		fmt.Sprint(l.lat) +
		" and longitude: " +
		fmt.Sprint(l.long) +
		"\n"
	return description
}

func (g GPS) description() string {
	description := "Current location : \n" +
		g.current.description() +
		"Destination location : \n" +
		g.destination.description()
	return description
}

func (g GPS) distance() float64 {
	return g.world.distance(g.current, g.destination)
}

func (g GPS) message() string {
	return "Distance between " +
		g.current.name +
		" and " +
		g.destination.name +
		": " +
		fmt.Sprint(g.distance()) +
		"\n"
}

func main() {
	bradbury := Location{name: "Bradbury Landing", lat: -4.5895, long: 137.4417}
	elysium := Location{name: "Elysium Planitia", lat: 4.5, long: 135.9}
	world := World{radius: 5.0}
	gps := GPS{current: bradbury, destination: elysium, world: world}
	rover := Rover{GPS: gps}
	fmt.Println(rover.message())
}
