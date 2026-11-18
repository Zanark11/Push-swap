package main

import (
	"push-swap/pile"
	"sort"
)

// Trier vérifie la pile et lance le tri si nécessaire
func Trier(piles *pile.Piles) []string {

	// On prépare une liste pour stocker les instructions
	instructions := make([]string, 0)

	// Si la pile est déjà triée, on s'arrête
	if pile.EstTriee(*piles) {
		return instructions
	}

	// Si la pile contient deux nombres
	if len(piles.A) == 2 {
		if piles.A[0] > piles.A[1] {
			pile.Sa(piles)
			instructions = append(instructions, "sa")
		}
	}

	// Si la pile contient trois nombres
	if len(piles.A) == 3 {
		TrierTroisNombres(piles, &instructions)
	}

	// Si la pile contient quatre nombres
	if len(piles.A) == 4 {

		// On place le plus petit nombre en première position
		PlacerMinimum(piles, &instructions)

		// On envoie le plus petit nombre dans B
		pile.Pb(piles)
		instructions = append(instructions, "pb")

		// On trie les trois nombres restants
		TrierTroisNombres(piles, &instructions)

		// On récupère le plus petit nombre dans A
		pile.Pa(piles)
		instructions = append(instructions, "pa")
	}

	// Si la pile contient cinq nombres
	if len(piles.A) == 5 {

		// On place le plus petit nombre en première position
		PlacerMinimum(piles, &instructions)

		// On envoie le plus petit nombre dans B
		pile.Pb(piles)
		instructions = append(instructions, "pb")

		// On cherche le nouveau plus petit nombre
		PlacerMinimum(piles, &instructions)

		// On envoie le deuxième plus petit nombre dans B
		pile.Pb(piles)
		instructions = append(instructions, "pb")

		// On trie les trois nombres restants
		TrierTroisNombres(piles, &instructions)

		// On récupère les deux nombres de B
		pile.Pa(piles)
		instructions = append(instructions, "pa")

		pile.Pa(piles)
		instructions = append(instructions, "pa")
	}

	// Si la pile contient plus de cinq nombres
	if len(piles.A) > 5 {
		TrierPlusDeCinq(piles, &instructions)
	}
	return instructions
}

// TrouverMinimum cherche la position du plus petit nombre dans A
func TrouverMinimum(piles *pile.Piles) int {
	position := 0

	for i := 1; i < len(piles.A); i++ {
		if piles.A[i] < piles.A[position] {
			position = i
		}
	}

	return position
}

// TrierNombres prépare une copie triée des nombres
func TrierNombres(piles *pile.Piles) []int {

	// On crée une nouvelle liste de même taille que A
	nombresTries := make([]int, len(piles.A))

	// On copie les nombres de A dans cette nouvelle liste
	copy(nombresTries, piles.A)

	// On trie la copie dans l'ordre croissant
	sort.Ints(nombresTries)

	return nombresTries
}

// TrouverRang cherche le rang d'un nombre dans la liste triée
func TrouverRang(nombre int, nombresTries []int) int {

	// On parcourt la liste triée
	for i := 0; i < len(nombresTries); i++ {

		// Si le nombre correspond à celui recherché
		if nombresTries[i] == nombre {
			return i
		}
	}

	// Si le nombre n'est pas trouvé
	return -1
}

// RangSommet retourne le rang du nombre situé en haut de A
func RangSommet(piles *pile.Piles, nombresTries []int) int {

	// On récupère le nombre situé en première position de A
	nombre := piles.A[0]

	// On recherche son rang dans la liste triée
	rang := TrouverRang(nombre, nombresTries)

	return rang
}

// PlacerMinimum met le plus petit nombre en première position
func PlacerMinimum(piles *pile.Piles, instructions *[]string) {
	position := TrouverMinimum(piles)
	taille := len(piles.A)

	if position <= taille/2 {
		for i := 0; i < position; i++ {
			pile.Ra(piles)
			*instructions = append(*instructions, "ra")
		}
	} else {
		for i := 0; i < taille-position; i++ {
			pile.Rra(piles)
			*instructions = append(*instructions, "rra")
		}
	}
}

// TrierTroisNombres trie une pile contenant pile poil trois nombres
func TrierTroisNombres(piles *pile.Piles, instructions *[]string) {
	a := piles.A[0]
	b := piles.A[1]
	c := piles.A[2]

	// Cas 1 : [1 3 2]
	if a < b && b > c && a < c {
		pile.Sa(piles)
		*instructions = append(*instructions, "sa")
		pile.Ra(piles)
		*instructions = append(*instructions, "ra")
	}

	// Cas 2 : [2 1 3]
	if a > b && a < c {
		pile.Sa(piles)
		*instructions = append(*instructions, "sa")
	}

	// Cas 3 : [2 3 1]
	if a < b && a > c {
		pile.Rra(piles)
		*instructions = append(*instructions, "rra")
	}

	// Cas 4 : [3 1 2]
	if a > b && a > c && b < c {
		pile.Ra(piles)
		*instructions = append(*instructions, "ra")
	}

	// Cas 5 : [3 2 1]
	if a > b && b > c {
		pile.Sa(piles)
		*instructions = append(*instructions, "sa")
		pile.Rra(piles)
		*instructions = append(*instructions, "rra")
	}
}

// TrierPlusDeCinq trie une pile contenant plus de cinq nombres
func TrierPlusDeCinq(piles *pile.Piles, instructions *[]string) {

	// On crée une copie triée de la pile A
	nombresTries := TrierNombres(piles)

	// On récupère le nombre total d'éléments
	taille := len(piles.A)

	// On choisit environ 5 groupes
	tailleGroupe := (taille + 4) / 5

	// On traite les groupes les uns après les autres
	for debut := 0; debut < taille; debut += tailleGroupe {

		// On définit la limite du groupe actuel
		limite := debut + tailleGroupe

		if limite > taille {
			limite = taille
		}

		// On mémorise le nombre d'éléments à examiner
		examines := len(piles.A)

		// On examine chaque élément de A
		for i := 0; i < examines; i++ {

			if len(piles.A) == 0 {
				break
			}

			// On récupère le rang du nombre en haut de A
			rang := RangSommet(piles, nombresTries)

			// Si le rang appartient au groupe actuel
			if rang >= debut && rang < limite {

				// On envoie le nombre dans B
				pile.Pb(piles)
				*instructions = append(*instructions, "pb")

			} else {

				// Sinon, on place le nombre en bas de A
				pile.Ra(piles)
				*instructions = append(*instructions, "ra")
			}
		}
	}

	// On récupère les nombres de B du plus grand au plus petit
	for len(piles.B) > 0 {

		position := TrouverMaximum(piles)

		// On choisit le chemin le plus court pour atteindre le maximum
		if position <= len(piles.B)/2 {

			for i := 0; i < position; i++ {
				pile.Rb(piles)
				*instructions = append(*instructions, "rb")
			}

		} else {

			for i := 0; i < len(piles.B)-position; i++ {
				pile.Rrb(piles)
				*instructions = append(*instructions, "rrb")
			}
		}

		// On replace le plus grand nombre dans A
		pile.Pa(piles)
		*instructions = append(*instructions, "pa")
	}
}

// TrouverMaximum cherche la position du plus grand nombre dans B
func TrouverMaximum(piles *pile.Piles) int {

	position := 0

	for i := 1; i < len(piles.B); i++ {

		if piles.B[i] > piles.B[position] {
			position = i
		}
	}

	return position
}
