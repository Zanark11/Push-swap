package main

import (
	"fmt"
	"os"

	"push-swap/pile"
)

func main() {

	// On récupère les arguments du terminal
	arguments := os.Args[1:]

	// Si aucun argument n'est fourni, on arrête le programme
	if len(arguments) == 0 {
		return
	}

	// On vérifie et convertit les arguments en nombres
	nombres, err := pile.ConvertirArguments(arguments)

	// Si les arguments sont invalides, on affiche Error
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error")
		return
	}

	// On initialise les deux piles
	piles := pile.InitialiserPiles(nombres)

	// On lance l'algorithme de tri
	instructions := Trier(&piles)

	// On affiche chaque instruction sur une ligne
	for _, instruction := range instructions {
		fmt.Println(instruction)
	}
}
