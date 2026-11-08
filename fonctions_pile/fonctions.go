package fonctions_pile

// Pa déplace le sommet de B au sommet de A, si B n'est pas vide.
func Pa(pileA []int, pileB []int) ([]int, []int) {
	if len(pileB) > 0 {
		haut := pileB[0]
		pileB = pileB[1:]
		pileA = append([]int{haut}, pileA...)
	}
	return pileA, pileB

}

// Pb déplace le sommet de A au sommet de B, si A n'est pas vide.
func Pb(pileA []int, pileB []int) ([]int, []int) {
	if len(pileA) > 0 {
		haut := pileA[0]
		pileA = pileA[1:]
		pileB = append([]int{haut}, pileB...)
	}
	return pileA, pileB
}

// Sa échange les deux premiers éléments de A si la pile en contient au moins deux.
func Sa(pileA []int) {
	if len(pileA) > 1 {
		pileA[0], pileA[1] = pileA[1], pileA[0]
	}
}

// Sb échange les deux premiers éléments de B si la pile en contient au moins deux.
func Sb(pileB []int) {
	if len(pileB) > 1 {
		pileB[0], pileB[1] = pileB[1], pileB[0]
	}
}

// Ss applique simultanément les échanges sa et sb.
func Ss(pileA []int, pileB []int) {
	Sa(pileA)
	Sb(pileB)
}

// Rra place le dernier élément de A au sommet.
func Rra(pileA []int) []int {
	if len(pileA) >= 2 {
		fond := pileA[len(pileA)-1]
		pileA = append([]int{fond}, pileA[:len(pileA)-1]...)
	}
	return pileA
}

// Rrb place le dernier élément de B au sommet.
func Rrb(pileB []int) []int {
	if len(pileB) >= 2 {
		fond := pileB[len(pileB)-1]
		pileB = append([]int{fond}, pileB[:len(pileB)-1]...)
	}
	return pileB
}

// Rrr applique simultanément les rotations inverses rra et rrb.
func Rrr(pileA []int, pileB []int) ([]int, []int) {
	pileB = Rrb(pileB)
	pileA = Rra(pileA)
	return pileA, pileB
}

// Ra place le sommet de A au fond de la pile.
func Ra(pileA []int) []int {
	if len(pileA) >= 2 {
		pileA = append(pileA[1:], pileA[0])
	}
	return pileA
}

// Rb place le sommet de B au fond de la pile.
func Rb(pileB []int) []int {
	if len(pileB) >= 2 {
		pileB = append(pileB[1:], pileB[0])
	}
	return pileB
}

// Rr applique simultanément les rotations ra et rb.
func Rr(pileA []int, pileB []int) ([]int, []int) {
	pileA = Ra(pileA)
	pileB = Rb(pileB)
	return pileA, pileB
}
