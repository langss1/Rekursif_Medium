package main

import (
	"fmt"
)

func decimalToQuarternary(n int) int {
	if n == 0 {
		return 0
	} else {
		return (n % 4) + decimalToQuarternary(n/4)*10
	}
}

func main() {
	var n int
	fmt.Print("Masukkan bilangan bulat desimal: ")
	fmt.Scan(&n)
	fmt.Println("Nilai quarternary dari", n, "adalah", decimalToQuarternary(n))
}
