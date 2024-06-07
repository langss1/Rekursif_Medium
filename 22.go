package main

import (
	"fmt"
)

func decimalToBiner(n int) int {
	if n == 0 {
		return 0
	} else {
		return (n % 2) + decimalToBiner(n/2)*10
	}
}

func main() {
	var n int
	fmt.Print("Masukkan bilangan bulat desimal: ")
	fmt.Scan(&n)
	fmt.Println("Nilai oktal dari", n, "adalah", decimalToOktal(n))
}
