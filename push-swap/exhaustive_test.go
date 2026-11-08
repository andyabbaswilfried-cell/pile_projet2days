package main

import (
	"main/fonctions_pile"
	"reflect"
	"testing"
)

func TestAlgoTrierAllPermutationsUpToSeven(t *testing.T) {
	for size := 0; size <= 7; size++ {
		values := make([]int, size)
		for i := range values {
			values[i] = i
		}
		checkAllPermutations(t, values, 0)
	}
}

func checkAllPermutations(t *testing.T, values []int, position int) {
	t.Helper()
	if position == len(values) {
		original := append([]int(nil), values...)
		pileA := append([]int(nil), values...)
		var pileB []int
		for _, move := range algoTrier(pileA, pileB) {
			var valid bool
			pileA, pileB, valid = fonctions_pile.Instructions(move, pileA, pileB)
			if !valid {
				t.Fatalf("algoTrier(%v) emitted invalid instruction %q", original, move)
			}
		}

		want := append([]int(nil), values...)
		for i := 1; i < len(want); i++ {
			for j := i; j > 0 && want[j] < want[j-1]; j-- {
				want[j], want[j-1] = want[j-1], want[j]
			}
		}
		if !reflect.DeepEqual(pileA, want) || len(pileB) != 0 {
			t.Fatalf("algoTrier(%v) ended with A=%v, B=%v; want A=%v, B=[]",
				original, pileA, pileB, want)
		}
		return
	}

	for i := position; i < len(values); i++ {
		values[position], values[i] = values[i], values[position]
		checkAllPermutations(t, values, position+1)
		values[position], values[i] = values[i], values[position]
	}
}
