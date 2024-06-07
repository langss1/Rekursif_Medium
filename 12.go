//buatlah fungsi saja

func jumlahDeretGanjil(x int) int {
	/*  fungsi mengembalikan hasil penjumlahan dari seluruh
		bilangan ganjil positif yang kurang dari atau sama dengan N */
		if x == 1 {
			return 1
		} else if x%2 == 0 {
			return jumlahDeretGanjil(x-1) // masukkan fungsi rekursif pada baris ini 
		}
		return x + jumlahDeretGanjil(x-2) // masukkan fungsi rekursif pada baris ini 
	}