package pile

import (
	"fmt"
	"strconv"
	"strings"
)

func ConvertirArguments(arguments []string) ([]int, error) {

	nombres := make([]int, 0)

	doublons := make(map[int]bool)

	for _, argument := range arguments {
		mots := strings.Fields(argument)
		for _, mot := range mots {
			nombre, err := strconv.Atoi(mot)
			if err != nil {
				return nil, fmt.Errorf("argument invalide : %s", mot)
			}

			if doublons[nombre] {
				return nil, fmt.Errorf("nombre en double : %d", nombre)
			}

			doublons[nombre] = true
			nombres = append(nombres, nombre)
		}
	}

	if len(nombres) == 0 {
		return nil, fmt.Errorf("aucun nombre fourni")
	}

	return nombres, nil
}

func InitialiserPiles(nombres []int) Piles {
	
	a := nombres

	b := make([]int, 0)
	
	return Piles{
		A: a,
		B: b,
	}
}

