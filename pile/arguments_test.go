package pile

import "testing"

func TestInitialiserPiles(t *testing.T) {

	// On prépare une liste de nombres
	nombres := []int{4, 2, 7, 1}

	// On initialise les deux piles
	piles := InitialiserPiles(nombres)

	// On vérifie que la pile A contient 4 éléments
	if len(piles.A) != 4 {
		t.Errorf("La pile A devrait contenir 4 éléments")
	}

	// On vérifie que les nombres sont dans le bon ordre
	for i := 0; i < len(nombres); i++ {
		if piles.A[i] != nombres[i] {
			t.Errorf("Erreur dans la pile A à la position %d", i)
		}
	}

	// On vérifie que la pile B est vide
	if len(piles.B) != 0 {
		t.Errorf("La pile B devrait être vide")
	}
}

func TestConvertirArguments(t *testing.T) {

	// Cas 1 : des nombres valides
	nombres, err := ConvertirArguments([]string{"4", "2", "7", "1"})

	if err != nil {
		t.Errorf("Une erreur inattendue est apparue : %v", err)
	}

	if len(nombres) != 4 {
		t.Errorf("On attendait 4 nombres, mais on en a reçu %d", len(nombres))
	}

	// Cas 2 : des nombres dans une seule chaîne
	nombres, err = ConvertirArguments([]string{"4 2 7 1"})

	if err != nil {
		t.Errorf("Une erreur inattendue est apparue : %v", err)
	}

	if len(nombres) != 4 {
		t.Errorf("On attendait 4 nombres, mais on en a reçu %d", len(nombres))
	}

	// Cas 3 : un texte qui n'est pas un nombre
	_, err = ConvertirArguments([]string{"bonjour"})

	if err == nil {
		t.Errorf("Une erreur était attendue pour un texte invalide")
	}

	// Cas 4 : deux nombres identiques
	_, err = ConvertirArguments([]string{"4", "2", "4"})

	if err == nil {
		t.Errorf("Une erreur était attendue pour un doublon")
	}

	// Cas 5 : aucun nombre
	_, err = ConvertirArguments([]string{})

	if err == nil {
		t.Errorf("Une erreur était attendue pour une liste vide")
	}
}

func TestSa(t *testing.T) {

	// On prépare les deux piles
	piles := InitialiserPiles([]int{4, 2, 7, 1})

	// On exécute l'instruction sa
	Sa(&piles)

	// On vérifie que les deux premiers nombres ont été échangés
	if piles.A[0] != 2 || piles.A[1] != 4 {
		t.Errorf("Sa devrait donner [2 4 7 1], mais donne %v", piles.A)
	}
}

func TestSb(t *testing.T) {

	// On initialise les piles
	piles := InitialiserPiles([]int{4, 2, 7, 1})

	// On ajoute des nombres dans la pile B
	piles.B = []int{8, 3, 5, 1}

	// On exécute l'instruction sb
	Sb(&piles)

	// On vérifie que les deux premiers nombres ont été échangés
	if piles.B[0] != 3 || piles.B[1] != 8 {
		t.Errorf("Sb devrait donner [3 8 5 1], mais donne %v", piles.B)
	}
}

func TestSs(t *testing.T) {

	// On initialise les piles
	piles := InitialiserPiles([]int{4, 2, 7, 1})

	// On prépare la pile B
	piles.B = []int{8, 3, 5, 1}

	// On exécute l'instruction ss
	Ss(&piles)

	// On vérifie la pile A
	if piles.A[0] != 2 || piles.A[1] != 4 {
		t.Errorf("Ss devrait modifier A en [2 4 7 1], mais donne %v", piles.A)
	}

	// On vérifie la pile B
	if piles.B[0] != 3 || piles.B[1] != 8 {
		t.Errorf("Ss devrait modifier B en [3 8 5 1], mais donne %v", piles.B)
	}
}

func TestPb(t *testing.T) {

	// On initialise les piles
	piles := InitialiserPiles([]int{4, 2, 7})

	// On déplace le premier élément de A vers B
	Pb(&piles)

	// On vérifie les deux piles
	if piles.A[0] != 2 || piles.B[0] != 4 {
		t.Errorf("Pb incorrect : A = %v, B = %v", piles.A, piles.B)
	}
}

func TestPa(t *testing.T) {

	// On initialise les piles
	piles := InitialiserPiles([]int{4, 2, 7})
	piles.B = []int{8, 3}

	// On déplace le premier élément de B vers A
	Pa(&piles)

	// On vérifie les deux piles
	if piles.A[0] != 8 || piles.B[0] != 3 {
		t.Errorf("Pa incorrect : A = %v, B = %v", piles.A, piles.B)
	}
}

func TestRa(t *testing.T) {

	piles := InitialiserPiles([]int{4, 2, 7, 1})

	Ra(&piles)

	if piles.A[0] != 2 || piles.A[3] != 4 {
		t.Errorf("Ra incorrect : A = %v", piles.A)
	}
}

func TestRb(t *testing.T) {

	piles := InitialiserPiles([]int{4, 2, 7, 1})
	piles.B = []int{8, 3, 5, 1}

	Rb(&piles)

	if piles.B[0] != 3 || piles.B[3] != 8 {
		t.Errorf("Rb incorrect : B = %v", piles.B)
	}
}

func TestRr(t *testing.T) {

	// On initialise les piles
	piles := InitialiserPiles([]int{4, 2, 7, 1})
	piles.B = []int{8, 3, 5, 1}

	// On exécute rr
	Rr(&piles)

	// On vérifie la pile A
	if piles.A[0] != 2 || piles.A[3] != 4 {
		t.Errorf("Rr incorrect pour A : %v", piles.A)
	}

	// On vérifie la pile B
	if piles.B[0] != 3 || piles.B[3] != 8 {
		t.Errorf("Rr incorrect pour B : %v", piles.B)
	}
}

func TestRra(t *testing.T) {

	piles := InitialiserPiles([]int{4, 2, 7, 1})

	Rra(&piles)

	if piles.A[0] != 1 || piles.A[3] != 7 {
		t.Errorf("Rra incorrect : A = %v", piles.A)
	}
}

func TestRrb(t *testing.T) {

	piles := InitialiserPiles([]int{4, 2, 7, 1})
	piles.B = []int{8, 3, 5, 1}

	Rrb(&piles)

	if piles.B[0] != 1 || piles.B[3] != 5 {
		t.Errorf("Rrb incorrect : B = %v", piles.B)
	}
}

func TestRrr(t *testing.T) {

	// On initialise les piles
	piles := InitialiserPiles([]int{4, 2, 7, 1})
	piles.B = []int{8, 3, 5, 1}

	// On exécute rrr
	Rrr(&piles)

	// On vérifie la pile A
	if piles.A[0] != 1 || piles.A[3] != 7 {
		t.Errorf("Rrr incorrect pour A : %v", piles.A)
	}

	// On vérifie la pile B
	if piles.B[0] != 1 || piles.B[3] != 5 {
		t.Errorf("Rrr incorrect pour B : %v", piles.B)
	}
}

func TestInstructionsAvecPilesCourtes(t *testing.T) {
	// On prépare une pile A avec un seul nombre
	piles := InitialiserPiles([]int{5})

	// Ces instructions ne doivent rien modifier
	Sa(&piles)
	Sb(&piles)
	Ss(&piles)
	Ra(&piles)
	Rb(&piles)
	Rr(&piles)
	Rra(&piles)
	Rrb(&piles)
	Rrr(&piles)

	// On vérifie que les piles n'ont pas changé
	if len(piles.A) != 1 || piles.A[0] != 5 {
		t.Errorf("La pile A ne devrait pas avoir changé : %v", piles.A)
	}

	if len(piles.B) != 0 {
		t.Errorf("La pile B devrait rester vide : %v", piles.B)
	}

	// Pa ne doit rien faire si B est vide
	Pa(&piles)

	// Pb déplace le nombre de A vers B
	Pb(&piles)

	// Un deuxième Pb ne doit rien faire car A est vide
	Pb(&piles)

	if len(piles.A) != 0 || len(piles.B) != 1 || piles.B[0] != 5 {
		t.Errorf("Résultat incorrect après Pb : A = %v, B = %v", piles.A, piles.B)
	}

	// Pa remet le nombre dans A
	Pa(&piles)

	// Un deuxième Pa ne doit rien faire car B est vide
	Pa(&piles)

	if len(piles.A) != 1 || piles.A[0] != 5 || len(piles.B) != 0 {
		t.Errorf("Résultat incorrect après Pa : A = %v, B = %v", piles.A, piles.B)
	}
}
