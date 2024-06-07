package main

import "fmt"

func banyakBilanganGenap(N int) int {
	if N <= 0 {
		return 0
	}
	if N%2 == 0 {
		return 1 + banyakBilanganGenap(N-2)
	}
	return banyakBilanganGenap(N - 1)
}

func main() {
	var N int
	fmt.Scan(&N)
	fmt.Println(banyakBilanganGenap(N))
}
