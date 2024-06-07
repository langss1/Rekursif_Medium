func jumlahdigit(n int)int{
	/*{ I.S. Terdefinisi sebuah bilangan bulat n.
		F.S. Fungsi mengembalikan jumlah digit n. }*/
		if n < 10{
			return 1
		}else{
			return 1 + jumlahdigit(n/10)
		}
	}
	