package main

import (
	"fmt"
	"os"
)

func ExitCommand(c *config, args ...string) error {
	fmt.Println("Exiting the program!")
	os.Exit(0)
	return nil
}
