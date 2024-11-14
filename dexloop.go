package main

import (
	"bufio"
	"fmt"
	"os"
	"pokedex/internal/pokeapi"
	"strings"
)

type Dex struct {
	cmdList    []CLICommand
	caughtList map[string]pokeapi.PokemonData
	pokeAPI    *pokeapi.PokeAPI
}

func NewDex() *Dex {
	return &Dex{
		cmdList: []CLICommand{
			{"help", "Displays a help message"},
			{"exit", "Exit the pokedex"},
			{"map", "Display location data, incrementing with each call"},
			{"mapb", "Display previous map data"},
			{"explore", "Display info on pokemon in a specific area"},
			{"catch", "Catch a pokemon and add it to your pokedex"},
			{"pokedex", "List all caught pokemon"},
		},
		caughtList: make(map[string]pokeapi.PokemonData),
		pokeAPI:    pokeapi.NewPokeAPI(),
	}
}

func (dex *Dex) welcomeMessage() {
	fmt.Println("Welcome to the Pokedex! Please enter a command")
}

func (dex *Dex) cmdHelp() {
	fmt.Println("Usage: ")
	for _, cmd := range dex.cmdList {
		cmd.PrintCommand()
	}
}

func (dex *Dex) cmdExit() {
	fmt.Println("Goodbye!")
	os.Exit(0)
}

func (dex *Dex) cmdMap() {
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

func (dex *Dex) cmdMapb() {
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

func (dex *Dex) cmdExplore(area string) {
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

func (dex *Dex) cmdCatch(name string) {
	pokemonData, err := dex.pokeAPI.GetPokemonData(name)
	if err == nil {
		fmt.Printf("%s was caught! \n", name)
		// Add the caught mon to the dex
		dex.caughtList[name] = pokemonData
	} else {
		fmt.Printf("Error catching mon: %s\n", err)
	}
}

func (dex *Dex) cmdInspect(name string) {
	if mon, ok := dex.caughtList[name]; ok {
		mon.DisplayPokemonData()
	} else {
		fmt.Printf("You have not caught a %s\n", name)
	}
}

func (dex *Dex) cmdPokedex() {
	fmt.Println("Your pokedex: ")
	for name := range dex.caughtList {
		fmt.Printf(" - %s\n", name)
	}
}

func parseArgs(input string) (cmd, param *string) {
	args := strings.Split(input, " ")

	switch len(args) {
	case 1:
		cmd = &args[0]
		param = nil
	case 2:
		cmd = &args[0]
		param = &args[1]
	default:
		cmd = nil
		param = nil
	}

	return
}

func (dex *Dex) Repl() {
	dex.welcomeMessage()

	cli := bufio.NewScanner(os.Stdin)

	for {
		// Get a new command from the user
		// Parse a potential parameter as well
		cli.Scan()
		cmd, param := parseArgs(cli.Text())

		// Nil command indicates issue parsing args
		if cmd == nil {
			continue
		}

		switch *cmd {
		case "help":
			dex.cmdHelp()
		case "exit":
			dex.cmdExit()
		case "map":
			dex.cmdMap()
		case "mapb":
			dex.cmdMapb()
		case "explore":
			if param == nil {
				fmt.Println("Error: explore command needs an area!")
			} else {
				dex.cmdExplore(*param)
			}
		case "catch":
			if param == nil {
				fmt.Println("Error: catch command needs a name!")
			} else {
				dex.cmdCatch(*param)
			}
		case "inspect":
			if param == nil {
				fmt.Println("Error: inspect command needs a name!")
			} else {
				dex.cmdInspect(*param)
			}
		case "pokedex":
			dex.cmdPokedex()
		default:
			fmt.Println("Invalid command")
		}
	}
}
