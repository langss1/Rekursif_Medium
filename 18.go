package main

import "fmt"

func perkalianBilanganGenap(N int) int {
	if N <= 0 {
		return 1
	}
	if N%2 == 1 {
		return N * perkalianBilanganGenap(N-2)
	}
	return perkalianBilanganGenap(N - 1)
}

func main() {
	var N int
	fmt.Scan(&N)
	fmt.Println(perkalianBilanganGenap(N))
}
