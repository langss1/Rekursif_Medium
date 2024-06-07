func maksimum(max int) int {
    var x int
		fmt.Scan(&x)
		// gunakan percabangan untuk memanggil fungsi rekursif
		if x == 0 {
			return max
		}
		if x > max {
			max = x
		}
		return maksimum(max)
	}