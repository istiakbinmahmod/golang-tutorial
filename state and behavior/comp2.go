package main

import (
	"encoding/json"
	"fmt"
	"os"
	"time"
)

type Event struct {
	ID int
	time.Time
}

func main() {
	event := Event{
		ID:   6735,
		Time: time.Now(),
	}
	second := event.Second()
	fmt.Printf("second: %d\n", second)
	b, err := json.Marshal(event)
	if err != nil {
		os.Exit(1)
	}
	fmt.Printf("json: %s\n", string(b))
}
