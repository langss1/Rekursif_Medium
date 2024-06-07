func hitungBakteri(x, i, bakteri int) int {
    /*I.S. Terdefinisi nilai bilangan bulat x. 
    /*F.S. fungsu mengembalikan jumlah banyaknya bakteri saat x*/    
	if i == x+ 1{
        return bakteri
	}

	return hitungBakteri(x, i+1, bakteri*2)

}