package day2

import (
	"reflect"
	"testing"
)

func Test_calculate(t *testing.T) {
	tests := []struct {
		name string
		in   []int
		out  []int
	}{
		{"In: 1,0,0,0,99", []int{1, 0, 0, 0, 99}, []int{2, 0, 0, 0, 99}},
		{"In: 2,3,0,3,99", []int{2, 3, 0, 3, 99}, []int{2, 3, 0, 6, 99}},
		{"In: 2,4,4,5,99,0", []int{2, 4, 4, 5, 99, 0}, []int{2, 4, 4, 5, 99, 9801}},
		{"In: 1,1,1,4,99,5,6,0,99", []int{1, 1, 1, 4, 99, 5, 6, 0, 99}, []int{30, 1, 1, 4, 2, 5, 6, 0, 99}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotOut := calculate(tt.in); !reflect.DeepEqual(gotOut, tt.out) {
				t.Errorf("calculate() = %v, want %v", gotOut, tt.out)
			}
		})
	}
}

func Test_reverse(t *testing.T) {
	type args struct {
		in     []int
		result int
	}
	tests := []struct {
		name     string
		args     args
		wantNoun int
		wantVerb int
	}{
		{
			name: "Test for 9706670",
			args: args{
				in:     puzzleInput,
				result: 9706670,
			},
			wantNoun: 12,
			wantVerb: 2,
		},
		{
			name: "Test for 19690720",
			args: args{
				in:     puzzleInput,
				result: 19690720,
			},
			wantNoun: 25,
			wantVerb: 52,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotNoun, gotVerb := reverse(tt.args.in, tt.args.result)
			if gotNoun != tt.wantNoun {
				t.Errorf("reverse() gotNoun = %v, want %v", gotNoun, tt.wantNoun)
			}
			if gotVerb != tt.wantVerb {
				t.Errorf("reverse() gotVerb = %v, want %v", gotVerb, tt.wantVerb)
			}
		})
	}
}
