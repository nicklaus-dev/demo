package main

import (
	"encoding/json"
	"fmt"
	"time"
)

type Event struct {
	ID int
	time.Time
}

func main() {
	e := Event{
		ID:   1234,
		Time: time.Now(),
	}

	b, err := json.Marshal(e)
	if err != nil {
		fmt.Printf("err: %v\n", err)
	}

	fmt.Printf("e.Time: %v\n", e.Time)

	var e1 Event
	err = json.Unmarshal(b, &e1)
	if err != nil {
		fmt.Printf("err: %v\n", err)
	}

	fmt.Printf("e1.Time: %v\n", e1.Time)

}
