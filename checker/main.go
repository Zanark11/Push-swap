package main

import (
	"bufio"
	"fmt"
	"io"
	"os"

	"push-swap/pile"
)

// verifier contrôle les instructions et vérifie le résultat du tri
func verifier(arguments []string, entree io.Reader, sortie io.Writer, erreur io.Writer) {

	// On vérifie si des arguments ont été fournis
	if len(arguments) == 0 {
		return
	}

	// On convertit et vérifie les arguments
	nombres, err := pile.ConvertirArguments(arguments)

	if err != nil {
		fmt.Fprintln(erreur, "Error")
		return
	}

	// On initialise les deux piles
	piles := pile.InitialiserPiles(nombres)

	// On prépare la lecture des instructions
	scanner := bufio.NewScanner(entree)

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
			fmt.Fprintln(erreur, "Error")
			return
		}
	}

	// On vérifie si la lecture s'est terminée correctement
	if scanner.Err() != nil {
		fmt.Fprintln(erreur, "Error")
		return
	}

	// On vérifie si la pile A est triée et si la pile B est vide
	if pile.EstTriee(piles) {
		fmt.Fprintln(sortie, "OK")
	} else {
		fmt.Fprintln(sortie, "KO")
	}
}

// main lance le checker avec les entrées et sorties du terminal
func main() {
	verifier(os.Args[1:], os.Stdin, os.Stdout, os.Stderr)
}
