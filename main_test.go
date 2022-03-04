package main

import (
	"reflect"
	"testing"
)

func Test_createSlice(t *testing.T) {
	tests := []struct {
		name string
		n    int
		want []int
	}{
		{
			name: "createSlice will return a list of 1 to 20 integers",
			n:    20,
			want: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := createSlice(tt.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("createSlice() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_fizzBuzz(t *testing.T) {
	tests := []struct {
		name  string
		slice []int
		want  []string
	}{
		{
			name:  "fizzBuzz() will return the correct output for a 1-20 integers list",
			slice: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20},
			want:  []string{"1", "2", "lucky", "4", "buzz", "fizz", "7", "8", "fizz", "buzz", "11", "fizz", "lucky", "14", "fizzbuzz", "16", "17", "fizz", "19", "buzz"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := fizzBuzz(tt.slice); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("fizzBuzz() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_frequencies(t *testing.T) {
	tests := []struct {
		name string
		fb   []string
		want map[string]int
	}{
		{
			name: "will return the correct frequencies",
			fb:   []string{"1", "2", "lucky", "4", "buzz", "fizz", "7", "8", "fizz", "buzz", "11", "fizz", "lucky", "14", "fizzbuzz", "16", "17", "fizz", "19", "buzz"},
			want: map[string]int{
				fizz:     4,
				buzz:     3,
				fizzbuzz: 1,
				lucky:    2,
				integer:  10,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := frequencies(tt.fb); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("frequencies() = %v, want %v", got, tt.want)
			}
		})
	}
}
