package main

import (
	"bufio"
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
	var geste []string
	var ok bool
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		// Chaque ligne est une instruction qui sera appliquée dans l'ordre.
		line := scanner.Text()
		geste = append(geste, line)
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "Error")
		return
	}
	for _, r := range pile {
		value, _ := strconv.Atoi(r)
		pileA = append(pileA, value)
	}
	// Une instruction inconnue invalide toute la séquence.
	for _, r := range geste {
		pileA, pileB, ok = fonctions_pile.Instructions(r, pileA, pileB)
		if !ok {
			fmt.Fprintln(os.Stderr, "Error")
			return
		}
	}
	if !fonctions_pile.Sorted(pileA, pileB) {
		fmt.Fprintln(os.Stdout, "KO")
		return
	}
	fmt.Fprintln(os.Stdout, "OK")
}
