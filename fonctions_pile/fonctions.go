package fonctions_pile

func Pa(pileA []int, pileB []int) ([]int, []int) {
	if len(pileB) > 0 {
		haut := pileB[0]
		pileB = pileB[1:]
		pileA = append([]int{haut}, pileA...)
	}
	return pileA, pileB

}
func Pb(pileA []int, pileB []int) ([]int, []int) {
	if len(pileA) > 0 {
		haut := pileA[0]
		pileA = pileA[1:]  
		pileB = append([]int{haut}, pileB...)
	}
	return pileA, pileB
}

func Sa(pileA []int) {
	if len(pileA) > 1 {
		pileA[0], pileA[1] = pileA[1], pileA[0]
	}
}

func Sb(pileB []int) {
	if len(pileB) > 1 {
		pileB[0], pileB[1] = pileB[1], pileB[0]
	}
}

func Ss(pileA []int, pileB []int) {
	Sa(pileA)
	Sb(pileB)
}

func Rra(pileA []int) []int {
	if len(pileA) >= 2 {
		fond := pileA[len(pileA)-1]
		pileA = append([]int{fond}, pileA[:len(pileA)-1]...)
	}
	return pileA
}
func Rrb(pileB []int) []int {
	if len(pileB) >= 2 {
		fond := pileB[len(pileB)-1]
		pileB = append([]int{fond}, pileB[:len(pileB)-1]...)
	}
	return pileB
}
func Rrr(pileA []int, pileB []int) ([]int, []int) {
	pileB = Rrb(pileB)
	pileA = Rra(pileA)
	return pileA, pileB
}

func Ra(pileA []int) []int {
	if len(pileA) >= 2 {
		pileA = append(pileA[1:], pileA[0])
	}
	return pileA
}

func Rb(pileB []int) []int {
	if len(pileB) >= 2 {
		pileB = append(pileB[1:], pileB[0])
	}
	return pileB
}

func Rr(pileA []int, pileB []int) ([]int, []int) {
	pileA = Ra(pileA)
	pileB = Rb(pileB)
	return pileA, pileB
}


