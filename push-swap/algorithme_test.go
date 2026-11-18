package main

import (
	"reflect"
	"testing"

	"push-swap/pile"
)

func TestTrierDeuxNombres(t *testing.T) {

	// Cas 1 : les nombres sont déjà triés
	piles := pile.InitialiserPiles([]int{1, 2})

	instructions := Trier(&piles)

	if !reflect.DeepEqual(instructions, []string{}) {
		t.Errorf("Résultat attendu : [], obtenu : %v", instructions)
	}

	// Cas 2 : les nombres ne sont pas triés
	piles = pile.InitialiserPiles([]int{2, 1})

	instructions = Trier(&piles)

	if !reflect.DeepEqual(instructions, []string{"sa"}) {
		t.Errorf("Résultat attendu : [sa], obtenu : %v", instructions)
	}

	// On vérifie que la pile est maintenant triée
	if !pile.EstTriee(piles) {
		t.Errorf("La pile devrait être triée, obtenu : %v", piles.A)
	}
}

func TestTrierTroisNombres(t *testing.T) {

	// Les six configurations possibles avec trois nombres différents
	cas := [][]int{
		{1, 2, 3},
		{1, 3, 2},
		{2, 1, 3},
		{2, 3, 1},
		{3, 1, 2},
		{3, 2, 1},
	}

	// On teste chaque configuration
	for _, nombres := range cas {

		// On initialise les piles
		piles := pile.InitialiserPiles(nombres)

		// On récupère les instructions
		instructions := Trier(&piles)

		// On vérifie que la pile est triée
		if !pile.EstTriee(piles) {
			t.Errorf("La pile %v n'est pas triée. Instructions : %v. Résultat : %v",
				nombres, instructions, piles.A)
		}
	}
}

func TestTrouverMinimum(t *testing.T) {

	// On crée une pile avec plusieurs nombres
	piles := pile.InitialiserPiles([]int{5, 3, 8, 1, 6})

	// On recherche la position du minimum
	position := TrouverMinimum(&piles)

	// Le nombre 1 se trouve à l'index 3
	if position != 3 {
		t.Errorf("Position attendue : 3, position obtenue : %d", position)
	}
}

func TestPlacerMinimum(t *testing.T) {

	// Cas 1 : le minimum est dans la deuxième moitié
	piles := pile.InitialiserPiles([]int{5, 3, 8, 1, 6})
	instructions := make([]string, 0)

	PlacerMinimum(&piles, &instructions)

	// Le minimum doit être placé en première position
	if piles.A[0] != 1 {
		t.Errorf("Le minimum devrait être en première position, obtenu : %v", piles.A)
	}

	// On vérifie que deux rotations rra ont été utilisées
	if len(instructions) != 2 || instructions[0] != "rra" || instructions[1] != "rra" {
		t.Errorf("Instructions attendues : [rra rra], obtenues : %v", instructions)
	}

	// Cas 2 : le minimum est dans la première moitié
	piles = pile.InitialiserPiles([]int{3, 1, 4, 5, 6})
	instructions = make([]string, 0)

	PlacerMinimum(&piles, &instructions)

	if piles.A[0] != 1 {
		t.Errorf("Le minimum devrait être en première position, obtenu : %v", piles.A)
	}

	// On vérifie que ra a été utilisé une seule fois
	if len(instructions) != 1 || instructions[0] != "ra" {
		t.Errorf("Instructions attendues : [ra], obtenues : %v", instructions)
	}
}

func TestTrierQuatreNombres(t *testing.T) {

	// On prépare une pile contenant quatre nombre
	piles := pile.InitialiserPiles([]int{4, 2, 1, 3})
	instructions := Trier(&piles)

	// On vérifie que la pile est correctement triée
	if !pile.EstTriee(piles) {
		t.Errorf("La pile n'est pas correctement triée : %v", piles.A)
	}

	// On vérifie que des instructions ont été générées
	if len(instructions) == 0 {
		t.Error("Aucune instruction n'a été générée")
	}
}

func TestTrierCinqNombres(t *testing.T) {

	// On prépare une pile contenant cinq nombres
	piles := pile.InitialiserPiles([]int{5, 2, 4, 1, 3})
	instructions := Trier(&piles)

	// On vérifie que la pile est correctement triée
	if !pile.EstTriee(piles) {
		t.Errorf("La pile n'est pas correctement triée : %v", piles.A)
	}

	// On vérifie que des instructions ont été générées
	if len(instructions) == 0 {
		t.Error("Aucune instruction n'a été générée")
	}
}

func TestTrierDixNombres(t *testing.T) {

	// On prépare une pile contenant dix nombres
	piles := pile.InitialiserPiles([]int{8, 3, 10, 1, 6, 2, 9, 4, 7, 5})
	instructions := Trier(&piles)

	// On vérifie que la pile est correctement triée
	if !pile.EstTriee(piles) {
		t.Errorf("La pile n'est pas correctement triée : %v", piles.A)
	}

	// On vérifie que des instructions ont été générées
	if len(instructions) == 0 {
		t.Error("Aucune instruction n'a été générée")
	}
}

func TestTrierCentNombres(t *testing.T) {

	// On prépare une pile contenant 100 nombres dans le désordre
	nombres := make([]int, 100)

	for i := 0; i < 100; i++ {
		nombres[i] = 100 - i
	}

	piles := pile.InitialiserPiles(nombres)
	instructions := Trier(&piles)

	// On vérifie que la pile est correctement triée
	if !pile.EstTriee(piles) {
		t.Errorf("La pile n'est pas correctement triée : %v", piles.A)
	}

	// On affiche le nombre d'instructions utilisées
	t.Logf("Nombre d'instructions utilisées : %d", len(instructions))
}

func TestTrierCentNombresMelanges(t *testing.T) {

	// On prépare 100 nombres dans un ordre différent
	nombres := make([]int, 100)

	for i := 0; i < 100; i++ {
		nombres[i] = (i * 37) % 100
	}

	// On initialise les piles
	piles := pile.InitialiserPiles(nombres)
	instructions := Trier(&piles)

	// On vérifie que la pile est correctement triée
	if !pile.EstTriee(piles) {
		t.Errorf("La pile n'est pas correctement triée : %v", piles.A)
	}

	// On affiche le nombre d'instructions utilisées
	t.Logf("Nombre d'instructions utilisées : %d", len(instructions))
}

func TestTrierCinqCentsNombres(t *testing.T) {

	// On prépare 500 nombres dans un ordre différent
	nombres := make([]int, 500)

	for i := 0; i < 500; i++ {
		nombres[i] = (i * 137) % 500
	}

	// On initialise les piles
	piles := pile.InitialiserPiles(nombres)
	instructions := Trier(&piles)

	// On vérifie que la pile est correctement triée
	if !pile.EstTriee(piles) {
		t.Errorf("La pile n'est pas correctement triée")
	}

	// On affiche le nombre d'instructions utilisées
	t.Logf("Nombre d'instructions utilisées : %d", len(instructions))
}

func TestTrierNombresNegatifs(t *testing.T) {

	// On prépare une pile avec des nombres positifs, négatifs et zéro
	nombres := []int{5, -10, 0, 8, -3, 12, -1, 4}

	// On initialise les piles
	piles := pile.InitialiserPiles(nombres)
	instructions := Trier(&piles)

	// On vérifie que la pile est correctement triée
	if !pile.EstTriee(piles) {
		t.Errorf("La pile n'est pas correctement triée : %v", piles.A)
	}

	// On affiche le nombre d'instructions utilisées
	t.Logf("Nombre d'instructions utilisées : %d", len(instructions))
}

func TestTrierPileDejaTriee(t *testing.T) {

	// On prépare une pile déjà triée
	piles := pile.InitialiserPiles([]int{1, 2, 3, 4, 5})

	// On lance le tri
	instructions := Trier(&piles)

	// On vérifie que la pile reste triée
	if !pile.EstTriee(piles) {
		t.Errorf("La pile devrait rester triée : %v", piles.A)
	}

	// On vérifie qu'aucune instruction n'a été générée
	if len(instructions) != 0 {
		t.Errorf("Aucune instruction ne devrait être générée, obtenu : %v", instructions)
	}
}

func TestTrierPileVide(t *testing.T) {

	// On prépare une pile vide
	piles := pile.InitialiserPiles([]int{})

	// On lance le tri
	instructions := Trier(&piles)

	// On vérifie que la pile est considérée comme triée
	if !pile.EstTriee(piles) {
		t.Errorf("Une pile vide devrait être considérée comme triée")
	}

	// On vérifie qu'aucune instruction n'a été générée
	if len(instructions) != 0 {
		t.Errorf("Aucune instruction ne devrait être générée")
	}
}

func TestTrouverRang(t *testing.T) {
	nombresTries := []int{-5, 12, 42, 100}

	rang := TrouverRang(100, nombresTries)

	if rang != 3 {
		t.Errorf("Rang attendu : 3, rang obtenu : %d", rang)
	}
}
