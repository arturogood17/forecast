package main

import (
	"time"

	"github.com/arturogood17/forecast/internal/weatherapi"
)

func main() {
	c := &config{
		Client: weatherapi.NewClient(5*time.Second, 1*time.Minute),
	}
	StartRepl(c)
}
