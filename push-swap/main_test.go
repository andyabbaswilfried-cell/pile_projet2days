package main

import (
	"main/fonctions_pile"
	"reflect"
	"strconv"
	"testing"
)

func TestChiffre(t *testing.T) {
	tests := []struct {
		carte int
		tour  int
		want  int
	}{
		{carte: 6, tour: 0, want: 0},
		{carte: 6, tour: 1, want: 1},
		{carte: 6, tour: 2, want: 1},
		{carte: 7, tour: 2, want: 1},
		{carte: 8, tour: 3, want: 1},
	}

	for _, tt := range tests {
		if got := Chiffre(tt.carte, tt.tour); got != tt.want {
			t.Errorf("Chiffre(%d, %d) = %d, want %d", tt.carte, tt.tour, got, tt.want)
		}
	}
}

func TestIntegerToBinary(t *testing.T) {
	maxInt := int(^uint(0) >> 1)
	tests := []struct {
		name string
		nums []int
		want int
	}{
		{name: "empty", nums: nil, want: 0},
		{name: "zero", nums: []int{0}, want: 1},
		{name: "maximum", nums: []int{1, maxInt}, want: strconv.IntSize - 1},
		{name: "largest number determines bit count", nums: []int{1, 4, 8}, want: 4},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IntegerToBinary(tt.nums); got != tt.want {
				t.Errorf("IntegerToBinary(%v) = %d, want %d", tt.nums, got, tt.want)
			}
		})
	}
}

func TestAlgoTrier(t *testing.T) {
	tests := []struct {
		name string
		pile []int
		want []int
	}{
		{name: "empty", pile: nil, want: nil},
		{name: "single value", pile: []int{5}, want: []int{5}},
		{name: "already sorted", pile: []int{1, 2, 3}, want: []int{1, 2, 3}},
		{name: "reverse order", pile: []int{5, 4, 3, 2, 1}, want: []int{1, 2, 3, 4, 5}},
		{name: "mixed values", pile: []int{8, 3, 6, 1, 7, 2}, want: []int{1, 2, 3, 6, 7, 8}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			moves := algo_Trier(append([]int(nil), tt.pile...), nil)
			pileA := append([]int(nil), tt.pile...)
			var pileB []int
			for _, move := range moves {
				var ok bool
				pileA, pileB, ok = fonctions_pile.Instructions(move, pileA, pileB)
				if !ok {
					t.Fatalf("algo_Trier returned invalid instruction %q", move)
				}
			}

			if !reflect.DeepEqual(pileA, tt.want) {
				t.Errorf("applying algo_Trier moves left pile A as %v, want %v", pileA, tt.want)
			}
			if len(pileB) != 0 {
				t.Errorf("applying algo_Trier moves left pile B as %v, want it empty", pileB)
			}
		})
	}
}
