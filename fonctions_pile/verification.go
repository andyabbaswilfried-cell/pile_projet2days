package fonctions_pile

import "strconv"

// AreAllNbrs indique si chaque argument peut être converti en int.
func AreAllNbrs(pile []string) bool {
	res := true
	for _, r := range pile {
		_, err := strconv.Atoi(r)
		if err != nil {
			res = false
			return res
		}
	}
	return res
}

// NoDoublons indique si aucune valeur textuelle n'apparaît deux fois.
func NoDoublons(pile []string) bool {
	res := true
	for i := 0; i < len(pile); i++ {
		for j := i + 1; j < len(pile); j++ {
			if pile[i] == pile[j] {
				res = false
				return res
			}
		}
	}
	return res
}

// Sorted vérifie que A est strictement croissante et que B est vide.
func Sorted(pileA []int, pileB []int) bool {
	if len(pileB) == 0 {
		for i := 1; i < len(pileA); i++ {
			if pileA[i-1] >= pileA[i] {

				return false
			}
		}
	} else {

		return false
	}
	return true
}
