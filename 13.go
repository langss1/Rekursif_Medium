//buatlagh fungsi saja

func banyakBilanganGanjil(n int) int {
	/*  fungsi mengembalikan  bilangan bulat yang menyatakan 
		banyaknya bilangan ganjil dari 1 hingga N */
		if n == 1{
			return 1
		}else if n%2 == 0{
			return banyakBilanganGanjil(n-1)
		}else{
			return 1 + banyakBilanganGanjil(n-1)
		}
	}