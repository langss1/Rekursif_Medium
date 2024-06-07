func reverse(str string, n int) {
    /*{ I.S. Terdefinisi str sebagai input string
	F.S. menampilkan string yang terbalik menggunakan fungsi rekursif*/
    if n > 0 {
        temp := string(str[n-1]) // hint :  gunakan var temporary
        fmt.Print(temp)
        reverse(str, n-1)
	}
}