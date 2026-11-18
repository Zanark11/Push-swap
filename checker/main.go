package main

import (
	"bufio"
	"fmt"
	"os"

	"push-swap/pile"
)

func main() {

	// On vérifie si des arguments ont été fournis
	if len(os.Args) < 2 {
		return
	}

	// On récupère les arguments
	arguments := os.Args[1:]

	// On convertit et vérifie les arguments
	nombres, err := pile.ConvertirArguments(arguments)

	if err != nil {
		fmt.Fprintln(os.Stderr, "Error")
		return
	}

	// On initialise les deux piles
	piles := pile.InitialiserPiles(nombres)

	// On prépare la lecture des instructions
	scanner := bufio.NewScanner(os.Stdin)

	// On lit les instructions une par une
	for scanner.Scan() {

		instruction := scanner.Text()

		// On exécute l'instruction correspondante
		switch instruction {
		case "sa":
			pile.Sa(&piles)
		case "sb":
			pile.Sb(&piles)
		case "ss":
			pile.Ss(&piles)
		case "pa":
			pile.Pa(&piles)
		case "pb":
			pile.Pb(&piles)
		case "ra":
			pile.Ra(&piles)
		case "rb":
			pile.Rb(&piles)
		case "rr":
			pile.Rr(&piles)
		case "rra":
			pile.Rra(&piles)
		case "rrb":
			pile.Rrb(&piles)
		case "rrr":
			pile.Rrr(&piles)
		default:
			fmt.Fprintln(os.Stderr, "Error")
			return
		}
	}

	// On vérifie si la lecture s'est terminée correctement
	if scanner.Err() != nil {
		fmt.Fprintln(os.Stderr, "Error")
		return
	}

	// On vérifie si la pile A est triée et si la pile B est vide
	if pile.EstTriee(piles) {
    	fmt.Println("OK")
	} else {
    	fmt.Println("KO")
	}
}
