package main

import "fmt"

func perkalianBilanganGanjil(N int) int {
	if N <= 0 {
		return 1
	}
	if N%2 == 1 {
		return N * perkalianBilanganGanjil(N-2)
	}
	return perkalianBilanganGanjil(N - 1)
}

func main() {
	var N int
	fmt.Scan(&N)
	fmt.Println(perkalianBilanganGanjil(N))
}
