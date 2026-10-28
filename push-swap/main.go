package main

import (
	"fmt"
	"main/fonctions_pile"
	"os"
	"strconv"
	"strings"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("il manque des arguments, la liste est vide")
		return
	}
	pile := strings.Fields(os.Args[1])
	if fonctions_pile.Are_all_nbrs(pile) && fonctions_pile.No_doublons(pile) {
		var pileA []int
		var pileB []int
		var moves []string
		for _, r := range strings.Fields(os.Args[1]) {

			value, err := strconv.Atoi(r)
			if err != nil {
				fmt.Println("la conversion échoue")
				return
			}
			pileA = append(pileA, value)
		}
		moves = algo_Trier(pileA, pileB)
		coups := strings.Join(moves, "\n")
		fmt.Print(coups)
		return
	}
}

func algo_Trier(pileA []int, pileB []int) []string {
	var instructions_à_faire []string
	if len(pileA) == 1 || len(pileA) == 0 {
		return nil
	} else if fonctions_pile.Sorted(pileA, pileB) {
		return nil
	} else {
		n := len(pileA)
		limit := IntegerToBinary(pileA)
		for tour := 0; tour < limit; tour++ { // 7 chiffres binaires suffisent jusqu'à 99 (1100011)
			for k := 0; k < n; k++ {
				// À toi : regarde le chiffre du tour de la carte du sommet de A.
				// S'il vaut 0, pose la carte sur B ; sinon, fais-la descendre au fond de A.
				if Chiffre(pileA[0], tour) == 0 {
					instructions_à_faire = append(instructions_à_faire, "pb")
					pileA, pileB = fonctions_pile.Pb(pileA, pileB)
				} else {
					instructions_à_faire = append(instructions_à_faire, "ra")
					pileA = fonctions_pile.Ra(pileA)
				}
			}
			for len(pileB) > 0 {
				instructions_à_faire = append(instructions_à_faire, "pa")
				pileA, pileB = fonctions_pile.Pa(pileA, pileB)
			}
		}
		// À toi : ramène toutes les cartes de B sur A.
	}
	return instructions_à_faire
}

// Chiffre renvoie le chiffre numéro tour de carte écrite en binaire, en partant
// de la droite. 6 s'écrit 110 : Chiffre(6, 0) vaut 0, Chiffre(6, 1) vaut 1.
func Chiffre(carte, tour int) int {
	return (carte >> tour) & 1
}
func IntegerToBinary(n []int) int {
	if len(n) == 0 {
		return 0
	}
	max := n[0]
	for _, r := range n {
		if r > max {
			max = r
		}
	}
	return len(strconv.FormatInt(int64(max), 2))
}
