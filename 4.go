package main

import (
	"fmt"
)

// Prosedur rekursif untuk mencetak pola bintang
func pola(n int, s string) {
	/* I.S. Terdefinisi nilai bilangan bulat n dimana n>0 */
	/* F.S. prosedur akan menampilkan pola dengan string (*) */
	if n <= 0 {
		return
	}

	// Mencetak baris bintang saat ini
	fmt.Println(s)

	// Panggilan rekursif untuk baris berikutnya dengan menambahkan satu bintang
	pola(n-1, s+"*")

	// Mencetak baris bintang saat kembali dari rekursi
	if n > 1 {
		fmt.Println(s)
	}
}

func main() {
	var n int
	fmt.Scan(&n)
	pola(n, "*")
}
