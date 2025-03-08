package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/arturogood17/forecast/internal/weatherapi"
)

type config struct {
	Client  weatherapi.Client
	Command map[string]Cli
}

type Cli struct {
	name     string
	descript string
	callback func(*config, ...string) error
}

func StartRepl(c *config) {
	for {
		fmt.Print("Weather app: ")
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		commands := scanner.Text()
		cl_comands := CleanCommands(commands)
		if len(cl_comands) == 1 {
			fmt.Println("add a city")
			continue
		}
		cmd_1 := cl_comands[0]
		args := cl_comands[1]
		v, exists := getCommands()[cmd_1]
		if !exists {
			fmt.Println("This command does not exist")
		} else {
			v.callback(c, args)
		}
	}
}

func getCommands() map[string]Cli {
	commands := map[string]Cli{
		"current": {
			name:     "current",
			descript: "current climate",
			callback: CurrentCommand,
		},
	}
	return commands
}

func CleanCommands(command string) []string {
	commands := strings.Fields(strings.ToLower(command))
	return commands
}
