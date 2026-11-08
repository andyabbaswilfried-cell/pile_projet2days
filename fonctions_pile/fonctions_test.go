package fonctions_pile

import (
	"reflect"
	"testing"
)

func TestInstructionsExhaustive(t *testing.T) {
	tests := []struct {
		name        string
		instruction string
		pileA       []int
		pileB       []int
		wantA       []int
		wantB       []int
		valid       bool
	}{
		{"pa", "pa", []int{1, 2}, []int{3, 4}, []int{3, 1, 2}, []int{4}, true},
		{"pb", "pb", []int{1, 2}, []int{3, 4}, []int{2}, []int{1, 3, 4}, true},
		{"sa", "sa", []int{1, 2}, []int{3, 4}, []int{2, 1}, []int{3, 4}, true},
		{"sb", "sb", []int{1, 2}, []int{3, 4}, []int{1, 2}, []int{4, 3}, true},
		{"ss", "ss", []int{1, 2}, []int{3, 4}, []int{2, 1}, []int{4, 3}, true},
		{"ra", "ra", []int{1, 2, 3}, []int{4, 5, 6}, []int{2, 3, 1}, []int{4, 5, 6}, true},
		{"rb", "rb", []int{1, 2, 3}, []int{4, 5, 6}, []int{1, 2, 3}, []int{5, 6, 4}, true},
		{"rr", "rr", []int{1, 2, 3}, []int{4, 5, 6}, []int{2, 3, 1}, []int{5, 6, 4}, true},
		{"rra", "rra", []int{1, 2, 3}, []int{4, 5, 6}, []int{3, 1, 2}, []int{4, 5, 6}, true},
		{"rrb", "rrb", []int{1, 2, 3}, []int{4, 5, 6}, []int{1, 2, 3}, []int{6, 4, 5}, true},
		{"rrr", "rrr", []int{1, 2, 3}, []int{4, 5, 6}, []int{3, 1, 2}, []int{6, 4, 5}, true},
		{"unknown instruction", "invalid", []int{1, 2}, []int{3, 4}, []int{1, 2}, []int{3, 4}, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotA, gotB, valid := Instructions(tt.instruction, tt.pileA, tt.pileB)
			if valid != tt.valid {
				t.Fatalf("Instructions(%q) validity = %v, want %v", tt.instruction, valid, tt.valid)
			}
			if !reflect.DeepEqual(gotA, tt.wantA) || !reflect.DeepEqual(gotB, tt.wantB) {
				t.Errorf("Instructions(%q) = (%v, %v), want (%v, %v)",
					tt.instruction, gotA, gotB, tt.wantA, tt.wantB)
			}
		})
	}
}

func TestInstructionsOnShortStacks(t *testing.T) {
	instructions := []string{"pa", "pb", "sa", "sb", "ss", "ra", "rb", "rr", "rra", "rrb", "rrr"}
	stacks := [][]int{nil, {}, {1}}

	for _, instruction := range instructions {
		for _, pileA := range stacks {
			for _, pileB := range stacks {
				t.Run(instruction, func(t *testing.T) {
					gotA, gotB, valid := Instructions(instruction, pileA, pileB)
					if !valid {
						t.Fatalf("Instructions(%q) marked a valid operation invalid", instruction)
					}
					noOp := false
					switch instruction {
					case "pa":
						noOp = len(pileB) == 0
					case "pb":
						noOp = len(pileA) == 0
					case "sa", "ra", "rra":
						noOp = len(pileA) < 2
					case "sb", "rb", "rrb":
						noOp = len(pileB) < 2
					case "ss", "rr", "rrr":
						noOp = len(pileA) < 2 && len(pileB) < 2
					}
					if !noOp {
						return
					}
					if !reflect.DeepEqual(gotA, pileA) || !reflect.DeepEqual(gotB, pileB) {
						t.Errorf("Instructions(%q, %v, %v) = (%v, %v), want unchanged stacks",
							instruction, pileA, pileB, gotA, gotB)
					}
				})
			}
		}
	}
}

func TestValidationCases(t *testing.T) {
	t.Run("integer parsing", func(t *testing.T) {
		tests := []struct {
			input []string
			want  bool
		}{
			{nil, true},
			{[]string{"0", "-2", "+3"}, true},
			{[]string{"", "1"}, false},
			{[]string{"1.5"}, false},
			{[]string{"word"}, false},
			{[]string{"999999999999999999999999999999"}, false},
		}
		for _, tt := range tests {
			if got := AreAllNbrs(tt.input); got != tt.want {
				t.Errorf("AreAllNbrs(%q) = %v, want %v", tt.input, got, tt.want)
			}
		}
	})

	t.Run("duplicate values", func(t *testing.T) {
		tests := []struct {
			input []string
			want  bool
		}{
			{nil, true},
			{[]string{"1"}, true},
			{[]string{"1", "2"}, true},
			{[]string{"1", "1"}, false},
		}
		for _, tt := range tests {
			if got := NoDoublons(tt.input); got != tt.want {
				t.Errorf("NoDoublons(%q) = %v, want %v", tt.input, got, tt.want)
			}
		}
	})

	t.Run("sorted state", func(t *testing.T) {
		tests := []struct {
			pileA []int
			pileB []int
			want  bool
		}{
			{nil, nil, true},
			{[]int{1}, nil, true},
			{[]int{1, 2, 3}, nil, true},
			{[]int{3, 2, 1}, nil, false},
			{[]int{1, 1}, nil, false},
			{[]int{1, 2}, []int{3}, false},
		}
		for _, tt := range tests {
			if got := Sorted(tt.pileA, tt.pileB); got != tt.want {
				t.Errorf("Sorted(%v, %v) = %v, want %v", tt.pileA, tt.pileB, got, tt.want)
			}
		}
	})
}
