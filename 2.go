func kali(m, n int) int {
    /*{ I.S. Terdefinisi m dan n  sebagai bilangan bulat
	F.S. function mengembalikan hasil dari m*n */    
	if n == 0 {
		return 0
	} else if n == 1{
		return m
	}else {
	    return m + kali(m, n-1)
	}
}
