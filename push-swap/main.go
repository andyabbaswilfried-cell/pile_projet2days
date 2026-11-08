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
		fmt.Fprintln(os.Stderr, "Error")
		return
	}
	pile := strings.Fields(os.Args[1])
	if !fonctions_pile.AreAllNbrs(pile) || !fonctions_pile.NoDoublons(pile) {
		fmt.Fprintln(os.Stderr, "Error")
		return
	}
	var pileA []int
	var pileB []int
	for _, r := range pile {
		value, _ := strconv.Atoi(r)
		pileA = append(pileA, value)
	}
	moves := algoTrier(pileA, pileB)
	coups := strings.Join(moves, "\n")
	fmt.Print(coups)
}

func algoTrier(pileA []int, pileB []int) []string {
	var instructionsToDo []string
	if len(pileA) == 1 || len(pileA) == 0 {
		return nil
	} else if fonctions_pile.Sorted(pileA, pileB) {
		return nil
	} else {
		min, positive := isMinimumPositive(pileA)
		if !positive {
			pileA = shiftAllNumbers(pileA, min)
		}
		n := len(pileA)
		limit := nbrOfBits(pileA)
		// Le tri radix traite les bits de poids faible au plus fort.
		for tour := 0; tour < limit; tour++ {
			for k := 0; k < n; k++ {
				// Les zéros passent sur B ; les uns restent dans le même ordre sur A.
				if chiffre(pileA[0], tour) == 0 {
					instructionsToDo = append(instructionsToDo, "pb")
					pileA, pileB = fonctions_pile.Pb(pileA, pileB)
				} else {
					instructionsToDo = append(instructionsToDo, "ra")
					pileA = fonctions_pile.Ra(pileA)
				}
			}
			for len(pileB) > 0 {
				instructionsToDo = append(instructionsToDo, "pa")
				pileA, pileB = fonctions_pile.Pa(pileA, pileB)
			}
		}
	}
	return instructionsToDo
}

// isMinimumPositive trouve le minimum et précise s'il est non négatif.
func isMinimumPositive(pileA []int) (int, bool) {
	min := pileA[0]
	for n := 1; n < len(pileA); n++ {
		if pileA[n] < min {
			min = pileA[n]
		}
	}
	if min < 0 {
		return min, false
	}
	return min, true
}

// shiftAllNumbers décale les valeurs afin que le minimum devienne zéro.
func shiftAllNumbers(pileA []int, min int) []int {
	var pile []int
	for _, r := range pileA {
		pile = append(pile, r+-1*min)
	}
	return pile
}

// chiffre renvoie le bit numéro tour, en partant de la droite.
func chiffre(carte, tour int) int {
	return (carte >> tour) & 1
}

// nbrOfBits renvoie le nombre de bits nécessaires pour représenter la plus grande valeur.
func nbrOfBits(n []int) int {
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
