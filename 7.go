func jumlahtiapdigit(n int)int{
    /*{Terdefinisi sebuah bilangan bulat n.
   Fungsi mengembalikan hasil penjumlahan dari tiap digit n. }
   
   hint :  gunakan modulo dan penjumlahan dengan fungsi rekursif*/
   if n < 10{
       return n
   }else{
       return n%10 + jumlahtiapdigit(n/10)
   }
}