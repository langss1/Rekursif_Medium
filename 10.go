func tabungan(t, n int) int {
    /*I.S. Terdefinisi nilai bilangan bulat n dan t*/
    /*F.S. fungsi ini menampilkan total tabungan kakak setelah n minggu. */
    if  n ==1{
        return t
    }else{
        return t + tabungan (t + 2500, n-1)
    }
}