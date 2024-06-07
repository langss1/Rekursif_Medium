package main

import "fmt"

func cetakGoLang(i, n int) {
	if i > n {
		return
	}

	if i%2 == 0 && i%3 == 0 {
		fmt.Println("GoLang")
	} else if i%2 == 0 {
		fmt.Println("Go")
	} else if i%3 == 0 {
		fmt.Println("Lang")
	} else {
		fmt.Println(i)
	}

	cetakGoLang(i+1, n)
}

func main() {
	var n int
	fmt.Print("Masukkan bilangan bulat positif N: ")
	fmt.Scan(&n)
	cetakGoLang(1, n)
}
