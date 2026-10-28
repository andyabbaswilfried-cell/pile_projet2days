package fonctions_pile

import "fmt"
import "strconv"

func Are_all_nbrs(pile []string) bool {
	res := true
	for _, r := range pile {
		_, err := strconv.Atoi(r)
		if err != nil {
			fmt.Println("cette pile n'est pas composé uniquement de chiffres", pile)
			res = false
		}
	}
	return res
}

func No_doublons(pile []string) bool {
	res := true
	for i := 0; i < len(pile); i++ {
		for j := i + 1; j < len(pile); j++ {
			if pile[i] == pile[j] {
				fmt.Println("Malheureusement il y a un doublon")
				res = false
			}
		}
	}
	return res
}

func Sorted(pileA []int, pileB []int) ( bool) {
	ok := true
	if len(pileB) == 0 {
		for i := 1; i < len(pileA); i++ {
			if pileA[i-1] == pileA[i] || pileA[i-1] > pileA[i] {
				ok = false
				return  ok
			}
		}
	} else {
		ok = false
		return ok
	}
	return ok
}
