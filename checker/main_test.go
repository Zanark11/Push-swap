package main

import (
	"bytes"
	"strings"
	"testing"
)

// TestVerifierPileTriee vérifie que le checker affiche OK lorsque la pile est correctement triée
func TestVerifierPileTriee(t *testing.T) {

	// On prépare les arguments
	arguments := []string{"1", "2", "3"}

	// On prépare une entrée vide car aucune instruction n'est nécessaire
	entree := strings.NewReader("")

	// On prépare deux zones pour récupérer les résultats
	var sortie bytes.Buffer
	var erreur bytes.Buffer

	// On lance le checker
	verifier(arguments, entree, &sortie, &erreur)

	// On vérifie que le résultat est OK
	if sortie.String() != "OK\n" {
		t.Errorf("Résultat attendu : OK, résultat obtenu : %q", sortie.String())
	}

	// On vérifie qu'aucune erreur n'a été affichée
	if erreur.String() != "" {
		t.Errorf("Aucune erreur attendue, obtenu : %q", erreur.String())
	}
}

// TestVerifierPileNonTriee vérifie que le checker affiche KO lorsque la pile n'est pas triée

func TestVerifierPileNonTriee(t *testing.T) {

	// On prépare les arguments avec une pile non triée
	arguments := []string{"2", "1", "3"}

	// On ne fournit aucune instruction
	entree := strings.NewReader("")

	// On prépare les zones pour récupérer les résultats
	var sortie bytes.Buffer
	var erreur bytes.Buffer

	// On lance le checker
	verifier(arguments, entree, &sortie, &erreur)

	// On vérifie que le résultat est KO
	if sortie.String() != "KO\n" {
		t.Errorf("Résultat attendu : KO, résultat obtenu : %q", sortie.String())
	}

	// On vérifie qu'aucune erreur n'a été affichée
	if erreur.String() != "" {
		t.Errorf("Aucune erreur attendue, obtenu : %q", erreur.String())
	}
}

// TestVerifierInstructionInconnue vérifie que le checker refuse une instruction qui n'existe pas
func TestVerifierInstructionInconnue(t *testing.T) {

	// On prépare une pile valide
	arguments := []string{"2", "1", "3"}

	// On envoie une instruction inconnue
	entree := strings.NewReader("bonjour\n")

	// On prépare les zones pour récupérer les résultats
	var sortie bytes.Buffer
	var erreur bytes.Buffer

	// On lance le checker
	verifier(arguments, entree, &sortie, &erreur)

	// On vérifie que rien n'est affiché dans la sortie normale
	if sortie.String() != "" {
		t.Errorf("Aucune sortie attendue, obtenu : %q", sortie.String())
	}

	// On vérifie que le checker affiche Error
	if erreur.String() != "Error\n" {
		t.Errorf("Erreur attendue : Error, résultat obtenu : %q", erreur.String())
	}
}

// TestVerifierArgumentInvalide vérifie que le checker
// refuse un argument qui n'est pas un nombre.
func TestVerifierArgumentInvalide(t *testing.T) {

	// On prépare des arguments contenant du texte.
	arguments := []string{"2", "bonjour", "3"}

	// On prépare une entrée vide.
	entree := strings.NewReader("")

	var sortie bytes.Buffer
	var erreur bytes.Buffer

	verifier(arguments, entree, &sortie, &erreur)

	// On vérifie que le checker affiche Error.
	if erreur.String() != "Error\n" {
		t.Errorf("Erreur attendue : Error, résultat obtenu : %q", erreur.String())
	}

	// On vérifie qu'aucun résultat normal n'est affiché.
	if sortie.String() != "" {
		t.Errorf("Aucune sortie attendue, obtenu : %q", sortie.String())
	}
}

// TestVerifierNombreEnDouble vérifie que le checker refuse les nombres présents plusieurs fois
func TestVerifierNombreEnDouble(t *testing.T) {

	// On prépare des arguments avec un doublon
	arguments := []string{"2", "2", "3"}

	entree := strings.NewReader("")

	var sortie bytes.Buffer
	var erreur bytes.Buffer

	verifier(arguments, entree, &sortie, &erreur)

	// On vérifie que le checker affiche Error
	if erreur.String() != "Error\n" {
		t.Errorf("Erreur attendue : Error, résultat obtenu : %q", erreur.String())
	}

	// On vérifie qu'aucun résultat normal n'est affiché
	if sortie.String() != "" {
		t.Errorf("Aucune sortie attendue, obtenu : %q", sortie.String())
	}
}

// TestVerifierSansArguments vérifie que le checker n'affiche rien lorsqu'aucun argument n'est fourni
func TestVerifierSansArguments(t *testing.T) {

	// On prépare une liste d'arguments vide
	arguments := []string{}

	entree := strings.NewReader("")

	var sortie bytes.Buffer
	var erreur bytes.Buffer

	verifier(arguments, entree, &sortie, &erreur)

	// On vérifie que rien n'est affiché
	if sortie.String() != "" {
		t.Errorf("Aucune sortie attendue, obtenu : %q", sortie.String())
	}

	// On vérifie qu'aucune erreur n'est affichée
	if erreur.String() != "" {
		t.Errorf("Aucune erreur attendue, obtenu : %q", erreur.String())
	}
}

// TestVerifierInstructionValide vérifie que le checker exécute correctement une instruction valide
func TestVerifierInstructionValide(t *testing.T) {

	// On prépare une pile qui n'est pas triée
	arguments := []string{"2", "1", "3"}

	// On envoie l'instruction sa.
	entree := strings.NewReader("sa\n")

	var sortie bytes.Buffer
	var erreur bytes.Buffer

	// On lance le checker
	verifier(arguments, entree, &sortie, &erreur)

	// On vérifie que le résultat est OK
	if sortie.String() != "OK\n" {
		t.Errorf("Résultat attendu : OK, résultat obtenu : %q", sortie.String())
	}

	// On vérifie qu'aucune erreur n'a été affichée
	if erreur.String() != "" {
		t.Errorf("Aucune erreur attendue, obtenu : %q", erreur.String())
	}
}
