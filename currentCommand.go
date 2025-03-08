package main

import (
	"errors"
	"fmt"
)

func CurrentCommand(c *config, city ...string) error {
	if len(city) != 1 {
		return errors.New("provide a valid city name")
	}
	ct := city[0]
	climate, err := c.Client.Climate(ct)
	if err != nil {
		return err
	}
	fmt.Printf("Region: %v\n", climate.Location.Name)
	fmt.Printf("Country: %v\n", climate.Location.Country)
	fmt.Printf("Temp: %v\n", climate.Current.TempC)
	fmt.Printf("Wind: %v\n", climate.Current.WindKph)
	fmt.Printf("Humidity: %v\n", climate.Current.Humidity)
	fmt.Printf("FeelsLike: %v\n", climate.Current.FeelslikeC)
	return nil
}
