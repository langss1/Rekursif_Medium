func pangkat(m, n int) int {
    /*{Terdefinisi sebuah bilangan bulat positif m dan n.
    Fungsi mengembalikan hasil m^n. }*/
    if n == 0  {
		return 1 
	} else if n == 1 {
		return  m
	}else {
	    return m * pangkat(m, n-1)
	}
	   
}