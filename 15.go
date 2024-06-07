package main

import "fmt"

func jumlahDeretGenap(N int) int {
	if N <= 0 {
		return 0
	}
	if N%2 == 0 {
		return N + jumlahDeretGenap(N-2)
	}
	return jumlahDeretGenap(N - 1)
}

func main() {
	var N int
	fmt.Scan(&N)
	fmt.Println(jumlahDeretGenap(N))
}
