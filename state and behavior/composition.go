package main

import "fmt"

type celsius float64

type Temperature struct {
	High, Low celsius
}

type Location struct {
	Lat, Long float64
}

type Report struct {
	sol         int
	Temperature Temperature
	Location    // struct embedding
}

func (t Temperature) Average() celsius {
	return (t.High + t.Low) / 2
}

func (r Report) Average() celsius {
	return r.Temperature.Average()
}

func (l Location) PrintLocation() {
	fmt.Printf("Lat: %v Long: %v\n", l.Lat, l.Long)
}

func main() {
	bradbury := Location{Lat: -4.5, Long: -13.5}
	t := Temperature{High: 10.0, Low: 15}
	report := Report{sol: 15, Temperature: t, Location: bradbury}
	fmt.Println(report)
	fmt.Printf("%+v\n", report)
	fmt.Println(report.Average())
	report.PrintLocation()
}
