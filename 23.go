package main

import (
	"fmt"
)

func decimalToOktal(n int) string {
	if n < 8 {
		return fmt.Sprint(n)
	}
	return decimalToOktal(n/8) + fmt.Sprint(n%8)
}

func main() {
	var n int
	fmt.Print("Masukkan bilangan bulat desimal: ")
	fmt.Scan(&n)
	fmt.Println("Nilai oktal dari", n, "adalah", decimalToOktal(n))
}
