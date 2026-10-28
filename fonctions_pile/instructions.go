package fonctions_pile

func Instructions(nom string, pileA []int, pileB []int) ([]int, []int,bool) {
	possible:= true
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
		possible= false
	}
	return pileA, pileB,possible
}