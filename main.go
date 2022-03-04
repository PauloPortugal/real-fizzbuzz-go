package main

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
)

const (
	RangeOfNumbers = 20
	fizz           = "fizz"
	buzz           = "buzz"
	fizzbuzz       = "fizzbuzz"
	lucky          = "lucky"
	integer        = "integer"
)

func main() {
	slice := createSlice(RangeOfNumbers)
	fb := fizzBuzz(slice)

	freq, err := json.MarshalIndent(frequencies(fb), "", "  ")
	if err != nil {
		fmt.Println("marshall error:", err)
	}

	fmt.Println(fb)
	fmt.Println(string(freq))
}

func createSlice(n int) []int {
	s := make([]int, n+1)
	for i := range s {
		s[i] = i
	}
	return s[1:]
}

func fizzBuzz(slice []int) []string {
	ns := make([]string, RangeOfNumbers)
	for i, v := range slice {
		switch {
		case strings.Contains(strconv.Itoa(v), "3"):
			ns[i] = "lucky"
		case v%15 == 0:
			ns[i] = "fizzbuzz"
		case v%3 == 0:
			ns[i] = "fizz"
		case v%5 == 0:
			ns[i] = "buzz"
		default:
			ns[i] = strconv.Itoa(v)
		}
	}
	return ns
}

func frequencies(fb []string) map[string]int {
	freq := map[string]int{
		fizz:     0,
		buzz:     0,
		fizzbuzz: 0,
		lucky:    0,
		integer:  0,
	}

	for _, v := range fb {
		switch v {
		case fizz:
			freq[fizz]++
		case buzz:
			freq[buzz]++
		case fizzbuzz:
			freq[fizzbuzz]++
		case lucky:
			freq[lucky]++
		default:
			freq[integer]++
		}
	}
	return freq
}
