package fonctions_pile

// Instructions applique une instruction reconnue et indique si son nom est valide.
// Les opérations impossibles sur des piles trop courtes restent des instructions valides
// et ne modifient pas ces piles.
func Instructions(nom string, pileA []int, pileB []int) ([]int, []int, bool) {
	possible := true
	switch nom {
	case "pb":
		pileA, pileB = Pb(pileA, pileB)
	case "pa":
		pileA, pileB = Pa(pileA, pileB)
	case "sa":
		Sa(pileA)
	case "sb":
		Sb(pileB)
	case "ss":
		Ss(pileA, pileB)
	case "ra":
		pileA = Ra(pileA)

	case "rb":
		pileB = Rb(pileB)
	case "rr":
		pileA, pileB = Rr(pileA, pileB)
	case "rra":
		pileA = Rra(pileA)
	case "rrb":
		pileB = Rrb(pileB)
	case "rrr":
		pileA, pileB = Rrr(pileA, pileB)

	default:
		possible = false
	}
	return pileA, pileB, possible
}
