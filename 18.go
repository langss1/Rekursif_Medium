func perkalianBilanganGenap(N int) int {
/*  mengembalikan sebuah bilangan bulat yang menyatakan 
    hasil perkalian bilangan genap dari 2 hingga N */
	if N <= 0 {
		return 1
	}
	if N%2 == 0 {
		return N * perkalianBilanganGenap(N-2)
	}
	return perkalianBilanganGenap(N - 1)
}
