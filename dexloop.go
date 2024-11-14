package main

import (
	"bufio"
	"fmt"
	"os"
	"pokedex/internal/pokeapi"
	"strings"
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
			CLICommand{"explore", "Display info on pokemon in a specific area"},
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

func (dex *Dex) Explore(area string) {
	areaData, err := dex.pokeAPI.GetAreaData(area)
	if err == nil {
		fmt.Printf("Exploring %s\n Found Pokemon: \n\n", area)
		for _, encounter := range areaData.PokemonEncounters {
			fmt.Printf(" - %s\n", encounter.Pokemon.Name)
		}
	} else {
		fmt.Printf("Error retrieving area data: %s\n", err)
	}
}

func (dex *Dex) Repl() {
	dex.WelcomeMessage()

	cli := bufio.NewScanner(os.Stdin)

	for {
		// Get a new command from the user
		// Parse a potential parameter as well
		cli.Scan()
		input := strings.Split((cli.Text()), " ")
		var cmd string
		var param string

		// TODO this is a bit janky
		switch len(input) {
		case 1:
			cmd = input[0]
		case 2:
			cmd = input[0]
			param = input[1]
		default:
			fmt.Println("Error: too many params")
			continue
		}

		switch cmd {
		case "help":
			dex.Help()
		case "exit":
			dex.Exit()
		case "map":
			dex.Map()
		case "mapb":
			dex.Mapb()
		case "explore":
			if param == "" {
				fmt.Println("Error: explore command needs an area!")
			} else {
				dex.Explore(param)
			}
		default:
			fmt.Println("Invalid command")
		}
	}
}
