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
		fmt.Println("il manque des arguments, la liste est vide")
		return
	}

	pile := strings.Fields(os.Args[1])
	if fonctions_pile.Are_all_nbrs(pile) && fonctions_pile.No_doublons(pile) {

		var pileA []int
		var pileB []int
		var geste []string
		var ok bool
		scanner := bufio.NewScanner(os.Stdin)
		for scanner.Scan() {
			// 3. Extract the line as a string
			line := scanner.Text()
			geste = append(geste, line)
		}
		if err := scanner.Err(); err != nil {
			fmt.Fprintln(os.Stderr, "Error", err)
			return
		}
		for _, r := range strings.Fields(os.Args[1]) {
			value, err := strconv.Atoi(r)
			if err != nil {
				fmt.Println("la conversion échoue")
				return
			}
			pileA = append(pileA, value)
		}
		for _, r := range geste {
			pileA, pileB, ok = fonctions_pile.Instructions(r, pileA, pileB)
			if ok == false {
				fmt.Fprintln(os.Stderr, "Error")
				return
			}
				
				
			
		
		}
		if sorted(pileA, pileB) == false{
					fmt.Fprintln(os.Stderr, "KO")
					fmt.Fprintln(os.Stdout, pileA)
					return
				}else{
					fmt.Fprintln(os.Stdout, pileA)
					fmt.Fprintln(os.Stdout, "OK")
					fmt.Print(pileA)
				}
	}else{
		fmt.Fprintln(os.Stdout, "Il y a un problème, vérifie tes entrés")
	}
}

func sorted(pileA []int, pileB []int) ( bool) {
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
