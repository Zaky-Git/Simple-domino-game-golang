//nama : Muhammad Zaky Fathurahim
//Kelas : SE-45-02
//Nim : 1302213067
//TP ALPRO MOD 12

package main

import (
	"fmt"
)

const nmax = 100000

type Trec struct {
	v1 int
	vx struct {
		v2, v3 int
	}
	v4 int
}

func BanyakNilai(rec *Trec) { //1.
	var arr_rec = [nmax]int{rec.v1, rec.vx.v2, rec.vx.v3, rec.v4}
	size := 4
	rec.v1 = NilaiMinimum(arr_rec, size)
	fmt.Println("Nilai Minimum dari Field: ", rec.v1)
	rec.vx.v2 = 0
	for _, i := range arr_rec {
		rec.vx.v2 = rec.vx.v2 + i
	}
	fmt.Println("Nilai total dari seluruh field semula: ", rec.vx.v2)
	rec.vx.v3 = rec.vx.v2 / size
	fmt.Println("Nilai rata rata dari field : ", rec.vx.v3)
	for i := 0; i < size; i++ {
		if arr_rec[i] > rec.v4 {
			rec.v4 = arr_rec[i]
		}
	}
	fmt.Printf("Nilai Terbesar dari semua Field : %v\n", rec.v4) //nilai max
}

func TambahData(tab *[nmax]int, n *int) { //2. input data
	for true {
		fmt.Scan(&tab[*n])
		if tab[*n] == 9999 {
			tab[*n] = 0
			break
		}
		*n = *n + 1
	}
}

func CariSekuensial(tab [nmax]int, v int) int { //3. sekuensial search
	for a, i := range tab {
		if i == v {
			return a
		}
	}
	return -1
}

func NilaiMinimum(tab [nmax]int, n int) int { //4. mencari nilai terkecil
	min := tab[0]
	for i := 0; i < n; i++ {
		if tab[i] < min {
			min = tab[i]
		}
	}
	return min
}

func NilaiRerata(tab [nmax]int, n int) int { //5. mencari nilai rata-rata
	nilai := 0
	for _, i := range tab {
		nilai = nilai + i
	}
	return nilai / n
}

func TerurutA(tab *[nmax]int, n int) { //6. mencari nilai terurut keatas (ascending) (selection sort)
	for i := 0; i < n; i++ {
		for j := i; j < n; j++ {
			if tab[i] > tab[j] {
				tab[i], tab[j] = tab[j], tab[i]
			}
		}
	}
}

// func TerurutB(tab *[nmax]int, n int) { //7. mencari nilai terurut kebawah (descending) (selection sort)
// 	for i := 0; i < n; i++ {
// 		for j := i; j < n; j++ {
// 			if tab[i] < tab[j] {
// 				tab[i], tab[j] = tab[j], tab[i]
// 			}
// 		}
// 	}
// }

func TerurutB(tab *[nmax]int, n int) { //kebawah tapi pakai insertion sort
	for i := 0; i < n-1; i++ {
		for j := i + 1; j > 0; j-- {
			if tab[j] > tab[j-1] {
				tab[j], tab[j-1] = tab[j-1], tab[j]
			}
		}
	}
}

func CariCepat(tab [nmax]int, n, v int) int { //8. biner
	TerurutB(&tab, n)
	low := 0
	up := n - 1

	for low <= up {
		med := (low + up) / 2
		if tab[med] > v {
			low = med + 1
		} else {
			up = med - 1
		}
	}

	if low != len(tab) && tab[low] == v {
		return low
	} else {
		return -1
	}
}

func Shaggy(tab [nmax]Trec, n int) { //9. procedure mencari shaggy (penjelasan cara kerja ada dibawah)
	var i int                      //deklarasi variable i tipe data integer
	var found1, found2, found bool // del;arasi variable bool
	found = false                  //inisialisasi found = false
	i = 2                          // Perulangan dibawah ini akan terus berulang selama found = false dan i kurang dari n
	for i < n && found == false {  // Operasi perulangan: mengecek found1 dan found2  dan i (i sebagai index) nya dimulai dari 2
		found1 = tab[i-1].v1 == tab[i].vx.v2 // Variable found1 akan bernilai true jika nilai property tab[i-1].v1 sama dengan tab[i].vx.v2 (property v2 didalam struct vx)
		found2 = tab[i].vx.v3 == tab[i].v4   // Variable found2 akan bernilai true jika nilai property tab[i].vx.v3 (property v3 didalam struct vx) sama dengan tab[i].v4
		found = found1 && found2             // Found akan bernilai true jika found1 dan found2 true.
		i++                                  //increment i akan bertambah 1 selama masih berulang (found masih false)
	}
	if found == true { //operasi pengkondisian (jika found = true)
		fmt.Println("Ada Shaggy disana. yaitu ...?") //output akan dikeluarkan
	} //pengkondisian selesai
}

// Jadi pada intinya procedure shaggy ini akan mencari shaggy selama found1 dan found2 nya false, menggunakan perulangan (i + 1)
// untuk mencari index yang nilainya sama (untuk membuat found1 dan found2 true dan found = true juga lalu perulangan selesai.)

func main() { //hanya untuk cek
	var n int
	var v int
	var tab [nmax]int
	var rec Trec
	fmt.Println("Masukkan nilai ke rec v1,v2,v3,v4: ")
	fmt.Scan(&rec.v1, &rec.vx.v2, &rec.vx.v3, &rec.v4)
	BanyakNilai(&rec)
	fmt.Println("Tambah data (masukkan 9999 jika ingin berhenti): ")
	TambahData(&tab, &n)
	fmt.Print("Masukkan angka yang ingin dicari: ")
	fmt.Scan(&v)
	fmt.Printf("Angka %v berada pada index (sequential): %v \n", v, CariSekuensial(tab, v))
	fmt.Printf("Nilai Minimum: %v\n", NilaiMinimum(tab, n))
	fmt.Printf("Nilai Rerata: %v\n", NilaiRerata(tab, n))

	fmt.Print("Ascending (Keatas): ")
	TerurutA(&tab, n)
	for _, j := range tab {
		if j == 0 {
			break
		}
		fmt.Print(j, " ")
	}

	TerurutB(&tab, n)
	fmt.Print("\nDescending (Kebawah): ")
	for _, k := range tab {
		if k == 0 {
			break
		}
		fmt.Print(k, " ")
	}
	fmt.Printf("\nAngka %v berada pada index (biner (cek descending)): %v", v, CariCepat(tab, n, v))
}
