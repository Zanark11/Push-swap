package pile

func Sa(piles *Piles) {
	if len(piles.A) < 2 {
		return
	}
	piles.A[0], piles.A[1] = piles.A[1], piles.A[0]
}

func Sb(piles *Piles) {
	if len(piles.B) < 2 {
		return
	}
	piles.B[0], piles.B[1] = piles.B[1], piles.B[0]
}

func Ss(piles *Piles) {
	Sa(piles)
	Sb(piles)
}

func Pb(piles *Piles) {
	if len(piles.A) == 0 {
		return
	}
	nombre := piles.A[0]
	piles.B = append([]int{nombre}, piles.B...)
	piles.A = piles.A[1:]
}

func Pa(piles *Piles) {
	if len(piles.B) == 0 {
		return
	}
	nombre := piles.B[0]
	piles.A = append([]int{nombre}, piles.A...)
	piles.B = piles.B[1:]
}

func Ra(piles *Piles) {
	if len(piles.A) < 2 {
		return
	}
	nombre := piles.A[0]
	piles.A = piles.A[1:]
	piles.A = append(piles.A, nombre)
}

func Rb(piles *Piles) {
	if len(piles.B) < 2 {
		return
	}
	nombre := piles.B[0]
	piles.B = piles.B[1:]
	piles.B = append(piles.B, nombre)
}

func Rr(piles *Piles) {
	Ra(piles)
	Rb(piles)
}

func Rra(piles *Piles) {
	if len(piles.A) < 2 {
		return
	}
	nombre := piles.A[len(piles.A)-1]
	piles.A = piles.A[:len(piles.A)-1]
	piles.A = append([]int{nombre}, piles.A...)
}

func Rrb(piles *Piles) {
	if len(piles.B) < 2 {
		return
	}
	nombre := piles.B[len(piles.B)-1]
	piles.B = piles.B[:len(piles.B)-1]
	piles.B = append([]int{nombre}, piles.B...)
}

func Rrr(piles *Piles) {
	Rra(piles)
	Rrb(piles)
}




