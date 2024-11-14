package main

import (
	"bufio"
	"fmt"
	"os"
	"pokedex/internal/pokeapi"
)

type Dex struct {
	cmdList []CLICommand
	pokeAPI *pokeapi.PokeAPI
}

func NewDex() *Dex {
	return &Dex{
		cmdList: []CLICommand{
			CLICommand{"help", "Displays a help message"},
			CLICommand{"exit", "Exit the pokedex"},
			CLICommand{"map", "Display location data, incrementing with each call"},
			CLICommand{"mapb", "Display previous map data"},
		},
		pokeAPI: pokeapi.NewPokeAPI(),
	}
}

func (dex *Dex) WelcomeMessage() {
	fmt.Println("Welcome to the Pokedex! Please enter a command")
}

func (dex *Dex) Help() {
	fmt.Println("Usage: \n")
	for _, cmd := range dex.cmdList {
		cmd.PrintCommand()
	}
}

func (dex *Dex) Exit() {
	fmt.Println("Goodbye!")
	os.Exit(0)
}

func (dex *Dex) Map() {
	mapdata, err := dex.pokeAPI.GetNextMapData()
	if err == nil {
		fmt.Println("Location data retrieved: ")
		for _, loc := range mapdata.Results {
			fmt.Println(loc.Name)
		}
	} else {
		fmt.Printf("Error retrieving map data: %s\n", err)
	}
}

func (dex *Dex) Mapb() {
	mapdata, err := dex.pokeAPI.GetPrevMapData()
	if err == nil {
		fmt.Println("Previous location data retrieved: ")
		for _, loc := range mapdata.Results {
			fmt.Println(loc.Name)
		}
	} else {
		fmt.Printf("Error retrieving map data: %s\n", err)
	}
}

func (dex *Dex) Repl() {
	dex.WelcomeMessage()

	cli := bufio.NewScanner(os.Stdin)

	for {
		// Get a new command from the user
		cli.Scan()
		cmd := cli.Text()

		switch cmd {
		case "help":
			dex.Help()
		case "exit":
			dex.Exit()
		case "map":
			dex.Map()
		case "mapb":
			dex.Mapb()
		default:
			fmt.Println("Invalid command")
		}
	}
}
