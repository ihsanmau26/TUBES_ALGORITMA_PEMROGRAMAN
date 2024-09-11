package main
import "fmt"

const NMAX int = 1024

type sparepart struct{
	nama string
	id_sparepart, harga, jumlah int
}

type pelanggan struct{
	nama, no_telp string
	id_pelanggan, service int
}

type transaksi struct{
	tgl, nama string
	id_transaksi, jenis, total int
	sparepart [NMAX-1]sparepart
}

type jenisSparepart [NMAX-1]sparepart

type customer [NMAX-1]pelanggan

type histori [NMAX-1]transaksi

func header(){
	fmt.Println()
	fmt.Println("\t\t\t   BENGKEL SERVICE MOTOR HELLO WORLD")
	fmt.Println("=======================================================================================")
}

func menu(){
	var A jenisSparepart
	var B customer
	var C histori
	var nSparepart, nHistori, nPelanggan int
	var x, y int

	header()
	fmt.Println("\tMenu:")
	fmt.Println("1. Data Sparepart")
	fmt.Println("2. Data Pelanggan")
	fmt.Println("3. Histori Transaksi")
	fmt.Println("0. Exit")
	fmt.Println()
	fmt.Print("Silahkan pilih menu \t: ")
	fmt.Scan(&x)
	
	for x == 1 || x == 2 || x == 3{
		for x == 1{
			fmt.Println()
			fmt.Println("Sparepart:")
			fmt.Println("1. Input Data")
			fmt.Println("2. Edit Data")
			fmt.Println("3. Hapus Data")
			fmt.Println("4. Lihat Data")
			fmt.Println("5. Cari Data")
			fmt.Println("0. Exit")
			fmt.Print("Silahkan pilih opsi \t: ")
			fmt.Scan(&y)
			fmt.Println()
			if y == 1{
				inputSparepart(&A, &nSparepart)
			}else if y == 2{
				fmt.Println("_______________________________________________________________________________________")
				printSparepart(A, nSparepart)
				fmt.Println("_______________________________________________________________________________________")
				editSparepart(&A, &C, nSparepart, nHistori)
			}else if y == 3{
				fmt.Println("_______________________________________________________________________________________")
				printSparepart(A, nSparepart)
				fmt.Println("_______________________________________________________________________________________")
				hapusSparepart(&A, &nSparepart)
			}else if y == 4{
				fmt.Println("Tampilkan berdasarkan \t:")
				fmt.Println("1. Paling sedikit dibeli")
				fmt.Println("2. Paling banyak dibeli")
				fmt.Println("3. Default")
				fmt.Println("_______________________________________________________________________________________")
				sparepartSort(A, nSparepart)
				fmt.Println("_______________________________________________________________________________________")
			}else if y == 5{
				fmt.Println("Cari Berdasarkan \t:")
				fmt.Println("1. ID")
				fmt.Println("2. Nama")
				fmt.Println("3. Harga")
				fmt.Println("4. Jumlah pembelian")
				fmt.Println("_______________________________________________________________________________________")
				cariSparepart(A, nSparepart)
				fmt.Println("_______________________________________________________________________________________")
			}else if y == 0{
				x = 0
			}
		}

		for x == 2{
			fmt.Println()
			fmt.Println("Data Pelanggan:")
			fmt.Println("1. Input Data")
			fmt.Println("2. Edit Data")
			fmt.Println("3. Hapus Data")
			fmt.Println("4. Lihat Data")
			fmt.Println("5. Cari Data")
			fmt.Println("0. Exit")
			fmt.Print("Silahkan pilih opsi \t: ")
			fmt.Scan(&y)
			fmt.Println()
			if y == 1{
				inputPelanggan(&B, &nPelanggan)
			}else if y == 2{
				fmt.Println("_______________________________________________________________________________________")
				printPelanggan(B, nPelanggan)
				fmt.Println("_______________________________________________________________________________________")
				editPelanggan(&B, nPelanggan)
			}else if y == 3{
				fmt.Println("_______________________________________________________________________________________")
				printPelanggan(B, nPelanggan)
				fmt.Println("_______________________________________________________________________________________")
				hapusPelanggan(&B, &nPelanggan)
			}else if y == 4{
				fmt.Println("_______________________________________________________________________________________")
				printPelanggan(B, nPelanggan)
				fmt.Println("_______________________________________________________________________________________")
			}else if y == 5{
				fmt.Println("Cari Berdasarkan \t:")
				fmt.Println("1. ID")
				fmt.Println("2. Nama")
				fmt.Println("3. Nomor Telepon")
				fmt.Println("4. Banyak Service")
				fmt.Println("_______________________________________________________________________________________")
				cariPelanggan(B, nPelanggan)
				fmt.Println("_______________________________________________________________________________________")
			}else if y == 0{
				x = 0
			}
		}

		for x == 3{
			fmt.Println()
			fmt.Println("Histori Transaksi:")
			fmt.Println("1. Input Data")
			fmt.Println("2. Edit Data")
			fmt.Println("3. Hapus Data")
			fmt.Println("4. Lihat Data")
			fmt.Println("5. Cari Data")
			fmt.Println("0. Exit")
			fmt.Print("Silahkan pilih opsi \t: ")
			fmt.Scan(&y)
			fmt.Println()
			if y == 1{
				inputHistori(&A, &B, &C, &nSparepart, &nPelanggan, &nHistori)
			}else if y == 2{
				fmt.Println("_______________________________________________________________________________________")
				printHistori(C, nHistori)
				fmt.Println("_______________________________________________________________________________________")
				editHistori(&A, &B, &C, &nSparepart, &nPelanggan, nHistori)
			}else if y == 3{
				fmt.Println("_______________________________________________________________________________________")
				printHistori(C, nHistori)
				fmt.Println("_______________________________________________________________________________________")
				hapusHistori(&A, &B, &C, nSparepart, nPelanggan, &nHistori)
			}else if y == 4{
				fmt.Println("_______________________________________________________________________________________")
				printHistori(C, nHistori)
				fmt.Println("_______________________________________________________________________________________")
			}else if y == 5{
				fmt.Println("Cari Berdasarkan \t:")
				fmt.Println("1. ID")
				fmt.Println("2. Tanggal")
				fmt.Println("3. Nama Pelanggan")
				fmt.Println("4. Nama Sparepart")
				fmt.Println("5. Banyak Jenis Sparepart")
				fmt.Println("6. Total Biaya")
				fmt.Println("_______________________________________________________________________________________")
				cariHistori(C, nHistori)
				fmt.Println("_______________________________________________________________________________________")
			}else if y == 0{
				x = 0
			}
		}

		fmt.Println()
		fmt.Println("Menu:")
		fmt.Println("1. Data Sparepart")
		fmt.Println("2. Data Pelanggan")
		fmt.Println("3. Histori Transaksi")
		fmt.Println("0. Exit")
		fmt.Println()
		fmt.Print("Silahkan pilih menu \t: ")
		fmt.Scan(&x)
	}

	fmt.Println()
	fmt.Println("============================== ANDA KELUAR DARI APLIKASI ==============================")
}

func inputSparepart(A *jenisSparepart, nSparepart *int){
	var input sparepart

	fmt.Println("\t\t\tKetik 'SAVE', untuk menyimpan data!")
	fmt.Println("_______________________________________________________________________________________")
	fmt.Print("Nama Sparepart \t \t: ")
	fmt.Scan(&input.nama)

	for input.nama != "SAVE"{
		fmt.Print("Harga Sparepart \t: ")
		fmt.Scan(&input.harga)
		fmt.Println("_______________________________________________________________________________________")
		A[*nSparepart] = input
		A[*nSparepart].id_sparepart = *nSparepart + 1
		*nSparepart++
		fmt.Print("Nama Sparepart \t \t: ")
		fmt.Scan(&input.nama)
	}
}

func inputPelanggan(B *customer, nPelanggan *int){
	var input pelanggan

	fmt.Println("\t\t\tKetik 'SAVE', untuk menyimpan data!")
	fmt.Println("_______________________________________________________________________________________")
	fmt.Print("Nama Pelanggan \t \t: ")
	fmt.Scan(&input.nama)

	for input.nama != "SAVE"{
		fmt.Print("No. Telepon Pelanggan \t: ")
		fmt.Scan(&input.no_telp)
		fmt.Println("_______________________________________________________________________________________")
		B[*nPelanggan] = input
		B[*nPelanggan].id_pelanggan = *nPelanggan + 1
		*nPelanggan++
		fmt.Print("Nama Pelanggan \t \t: ")
		fmt.Scan(&input.nama)
	}
}

func inputHistori(A* jenisSparepart, B *customer, C *histori, nSparepart, nPelanggan, nHistori *int){
	var input transaksi
	var found1, found2 bool = false, false

	fmt.Println("\t\t\tKetik 'SAVE', untuk menyimpan data!")
	fmt.Println("_______________________________________________________________________________________")
	fmt.Print("Tanggal Service \t: ")
	fmt.Scan(&input.tgl)

	for input.tgl != "SAVE"{
		fmt.Print("Nama Pelanggan \t \t: ")
		fmt.Scan(&input.nama)
		fmt.Print("Banyak Jenis Sparepart \t: ")
		fmt.Scan(&input.jenis)
		input.total = 0
		for i:=0; i < input.jenis; i++{
			fmt.Println()
			fmt.Print("Nama Sparepart \t \t: ")
			fmt.Scan(&input.sparepart[i].nama)
			fmt.Print("Jumlah Sparepart \t: ")
			fmt.Scan(&input.sparepart[i].jumlah)
			for j:=0; j < *nSparepart; j ++{
				if input.sparepart[i].nama == A[j].nama{
					input.sparepart[i].harga = A[j].harga
					A[j].jumlah = A[j].jumlah + input.sparepart[i].jumlah
					input.total = hitungTotal(input, i)
					found1 = true
				}
			}

			if found1 == false{
				A[*nSparepart].nama = input.sparepart[i].nama
				A[*nSparepart].jumlah = input.sparepart[i].jumlah

				fmt.Print("Harga Sparepart \t: ")
				fmt.Scan(&input.sparepart[i].harga)
				A[*nSparepart].harga = input.sparepart[i].harga
				
				A[*nSparepart].id_sparepart = *nSparepart + 1
				*nSparepart++
				input.total = hitungTotal(input, i)
			}

			found1 = false
		}

		fmt.Println("_______________________________________________________________________________________")
		C[*nHistori] = input
		C[*nHistori].id_transaksi = *nHistori + 1

		for k:=0; k < *nPelanggan; k++{
			if C[*nHistori].nama == B[k].nama{
				B[k].service++
				found2 = true
			}
		}

		if found2 == false{
			B[*nPelanggan].id_pelanggan = *nPelanggan + 1
			B[*nPelanggan].nama = C[*nHistori].nama
			B[*nPelanggan].no_telp = "NONE"
			B[*nPelanggan].service++
			*nPelanggan++
		}

		*nHistori++

		fmt.Println()
		fmt.Print("Tanggal Service \t: ")
		fmt.Scan(&input.tgl)
		found2 = false
	}
}

func hitungTotal(C transaksi, i int)int{
	return C.total + (C.sparepart[i].harga * C.sparepart[i].jumlah)
}

func sparepartSort(A jenisSparepart, nSparepart int){
	var urut, i, j, idx int
	var t sparepart

	i = 1

	fmt.Println()
	fmt.Print("Silahkan pilih opsi \t: ")
	fmt.Scan(&urut)

	if urut == 1{
		//selection sort
		for i <= nSparepart-1{
			idx = i-1
			j = i
			for j < nSparepart{
				if A[idx].jumlah > A[j].jumlah{
					idx = j
				}
				j = j+1
			}
			t = A[idx]
			A[idx] = A[i-1]
			A[i-1] = t
			i = i+1
		}
	}else if urut == 2{
		//insertion sort
		for i <= nSparepart-1{
			j = i
			idx = i-1
			t = A[j]
			for j > 0 && t.jumlah > A[j-1].jumlah{
				A[j] = A[j-1]
				j = j-1
			}
			A[j] = t
			i = i+1
		}
	}


	printSparepart(A, nSparepart)
}

func indeksSparepart(A jenisSparepart, nSparepart, id int)int{
	var found bool = false
	var i int = 0

	//sekuensial search
	for i < nSparepart && !found{
		found = A[i].id_sparepart == id
		i = i + 1
	}

	return i
}

func indeksPelanggan(B customer, nPelanggan, id int)int{
	var found bool = false
	var i int = 0

	//sekuensial search
	for i < nPelanggan && !found{
		found = B[i].id_pelanggan == id
		i = i + 1
	}

	return i
}

func indeksHistori(C histori, nHistori, id int)int{
	var kr, kn, med, found int 

	kr, kn, found = 0, nHistori-1, -1

	//binary search
	for kr <= kn && found == -1{
		med = (kr+kn)/2
		if id < C[med].id_transaksi{
			kn = med-1
		}else if id > C[med].id_transaksi{
			kr = med+1
		}else{
			found = med+1
		}
	}

	return found
}

func editSparepart(A *jenisSparepart, C *histori, nSparepart, nHistori int){
	var id, i int
	var new sparepart

	fmt.Print("ID Sparepart \t \t: ")
	fmt.Scan(&id)
	fmt.Println()

	if id > 0 && id <= nSparepart{
		i = indeksSparepart(*A, nSparepart, id)
		fmt.Print("Nama Sparepart \t \t: ")
		fmt.Scan(&new.nama)
		fmt.Print("Harga Sparepart \t: ")
		fmt.Scan(&new.harga)
		A[i-1].nama = new.nama
		A[i-1].harga = new.harga

		for j:=0; j < nHistori; j++{
			for k:=0; k < C[j].jenis; k++{
				if C[j].sparepart[k].nama == A[i-1].nama{
					C[j].total = C[j].total - (C[j].sparepart[k].harga * C[j].sparepart[k].jumlah)
					C[j].sparepart[k].harga = A[i-1].harga
					C[j].total = hitungTotal(C[j], k)
				}
			}
		}

		fmt.Println("_______________________________________________________________________________________")
	}else{
		fmt.Println()
		fmt.Println("Maaf, sparepart dengan ID", id, "tidak ada!")
		fmt.Println("_______________________________________________________________________________________")
	}
	
}

func editPelanggan(B *customer, C *histori, nPelanggan, nHistori int){
	var id, i int
	var new pelanggan

	fmt.Print("ID Pelanggan \t \t: ")
	fmt.Scan(&id)
	fmt.Println()

	if id > 0 && id <= nPelanggan{
		i = indeksPelanggan(*B, nPelanggan, id)
		fmt.Print("Nama Pelanggan \t \t: ")
		fmt.Scan(&new.nama)
		fmt.Print("No Telepon Pelanggan \t: ")
		fmt.Scan(&new.no_telp)
		for j:=0; j < nHistori: j++{
			if new.nama == C[j].nama{
				C[j].nama = new.nama
			}
		}
		B[i-1].nama = new.nama
		B[i-1].no_telp = new.no_telp

		fmt.Println("_______________________________________________________________________________________")
	}else{
		fmt.Println()
		fmt.Println("Maaf, pelanggan dengan ID", id, "tidak ada!")
		fmt.Println("_______________________________________________________________________________________")
	}
	
}

func editHistori(A *jenisSparepart, B *customer, C *histori, nSparepart, nPelanggan *int, nHistori int){
	var id, i int
	var new transaksi
	var found1, found2 bool = false, false

	fmt.Print("ID Transaksi \t \t: ")
	fmt.Scan(&id)
	fmt.Println()

	if id > 0 && id <= nHistori{
		i = indeksHistori(*C, nHistori, id)

		for y:=0; y < *nSparepart; y++{
			for z:=0; z < C[i-1].jenis; z++{
				if C[i-1].sparepart[z].nama == A[y].nama{
					A[y].jumlah = A[y].jumlah - C[i-1].sparepart[z].jumlah
				}
			}
		}

		for x:=0; x < *nPelanggan; x++ {
			if C[i-1].nama == B[x].nama{
				B[x].service--
			}
		}

		fmt.Print("Tanggal Service \t: ")
		fmt.Scan(&new.tgl)
		fmt.Print("Nama Pelanggan \t \t: ")
		fmt.Scan(&new.nama)
		fmt.Print("Banyak Jenis Sparepart \t: ")
		fmt.Scan(&new.jenis)

		for i:=0; i < new.jenis; i++{
			fmt.Println()
			fmt.Print("Nama Sparepart \t \t: ")
			fmt.Scan(&new.sparepart[i].nama)
			fmt.Print("Jumlah Sparepart \t: ")
			fmt.Scan(&new.sparepart[i].jumlah)
			for j:=0; j < *nSparepart; j ++{
					if new.sparepart[i].nama == A[j].nama{
						new.sparepart[i].harga = A[j].harga
						A[j].jumlah = A[j].jumlah + new.sparepart[i].jumlah
						new.total = hitungTotal(new, i)
						found1 = true
					}
				}

				if found1 == false{
					A[*nSparepart].nama = new.sparepart[i].nama
					A[*nSparepart].jumlah = new.sparepart[i].jumlah

					fmt.Print("Harga Sparepart \t: ")
					fmt.Scan(&new.sparepart[i].harga)
					A[*nSparepart].harga = new.sparepart[i].harga
					
					A[*nSparepart].id_sparepart = *nSparepart + 1
					*nSparepart++
					new.total = hitungTotal(new, i)
				}

				found1 = false

		}

		new.id_transaksi = C[i-1].id_transaksi

		C[i-1] = new

		for m:=0; m < nHistori; m++{
			if new.nama == B[m].nama{
				B[m].service++
				found2 = true
			}
		}

		if found2 == false{
			B[*nPelanggan].id_pelanggan = *nPelanggan + 1
			B[*nPelanggan].nama = C[i-1].nama
			B[*nPelanggan].no_telp = "NONE"
			B[*nPelanggan].service++
			*nPelanggan++
		}

		fmt.Println("_______________________________________________________________________________________")
	}else{
		fmt.Println()
		fmt.Println("Maaf, transaksi dengan ID", id, "tidak ada!")
		fmt.Println("_______________________________________________________________________________________")
	}

}

func hapusSparepart(A *jenisSparepart, nSparepart *int){
	var id, i int

	fmt.Print("ID Sparepart \t \t: ")
	fmt.Scan(&id)

	if id > 0 && id <= *nSparepart{
		for i = indeksSparepart(*A, *nSparepart, id)-1; i < *nSparepart; i++{
			A[i] = A[i+1]
			A[i].id_sparepart--
		}

		*nSparepart--
		fmt.Println("_______________________________________________________________________________________")
	}else{
		fmt.Println()
		fmt.Println("Maaf, sparepart dengan ID", id, "tidak ada!")
		fmt.Println("_______________________________________________________________________________________")
	}
}

func hapusPelanggan(B *customer, nPelanggan *int){
	var id, i int

	fmt.Print("ID Sparepart \t \t: ")
	fmt.Scan(&id)

	if id > 0 && id <= *nPelanggan{
		for i = indeksPelanggan(*B, *nPelanggan, id)-1; i < *nPelanggan; i++{
			B[i] = B[i+1]
			B[i].id_pelanggan--
		}

		*nPelanggan--
		fmt.Println("_______________________________________________________________________________________")
	}else{
		fmt.Println()
		fmt.Println("Maaf, pelanggan dengan ID", id, "tidak ada!")
		fmt.Println("_______________________________________________________________________________________")
	}
}

func hapusHistori(A *jenisSparepart, B *customer, C *histori, nSparepart, nPelanggan int, nHistori *int){
	var id, i int
	
	fmt.Print("ID Transaksi \t \t: ")
	fmt.Scan(&id)

	if id > 0 && id <= *nHistori{
		i = indeksHistori(*C, *nHistori, id)

		for j:=0; j < nSparepart; j++{
			for k:=0; k < C[i-1].jenis; k++{
				if C[i-1].sparepart[k].nama == A[j].nama{
					A[j].jumlah = A[j].jumlah - C[i-1].sparepart[k].jumlah
				}
			}
		}

		for l:=0; l < nPelanggan; l++{
			if C[i-1].nama == B[l].nama{
				B[l].service--
			}
		}

		for i = indeksHistori(*C, *nHistori, id)-1; i < *nHistori; i++{
			C[i] = C[i+1]
			C[i].id_transaksi--
		}
		*nHistori--

		fmt.Println("_______________________________________________________________________________________")
	}else{
		fmt.Println()
		fmt.Println("Maaf, transaksi dengan ID", id, "tidak ada!")
		fmt.Println("_______________________________________________________________________________________")
	}
}

func printSparepart(A jenisSparepart, nSparepart int){
	for i:=0; i < nSparepart; i++{
		fmt.Println(A[i].id_sparepart, A[i].nama, A[i].harga, A[i].jumlah)
	}
}

func printPelanggan(B customer, nPelanggan int){
	for i:=0; i < nPelanggan; i++{
		fmt.Println(B[i].id_pelanggan, B[i].nama, B[i].no_telp, B[i].service)
	}
}

func printHistori(C histori, nHistori int){
	for i:=0; i < nHistori; i++{
		fmt.Println(C[i].id_transaksi, C[i].tgl, C[i].nama, C[i].jenis)
		for j:=0; j < C[i].jenis; j++{
			fmt.Println(C[i].sparepart[j].nama, C[i].sparepart[j].harga, C[i].sparepart[j].jumlah)
		}
		fmt.Println(C[i].total)
		fmt.Println()
	}
}

func cariSparepart(A jenisSparepart, nSparepart int){
	var kategori, id, harga, jumlah int
	var nama string
	var found bool = false

	fmt.Print("Silahkan pilih opsi \t: ")
	fmt.Scan(&kategori)
	fmt.Println()

	if kategori == 1{
		fmt.Print("ID Sparepart \t\t: ")
		fmt.Scan(&id)
		for i:=0; i < nSparepart; i++{
			if id == A[i].id_sparepart{
				fmt.Println()
				fmt.Println(A[i].id_sparepart, A[i].nama, A[i].harga, A[i].jumlah)
				found = true
			}
		}

		if found == false{
			fmt.Println()
			fmt.Println("Maaf, sparepart dengan ID", id, "tidak ada!")
		}
	}else if kategori == 2{
		fmt.Print("Nama Sparepart \t\t: ")
		fmt.Scan(&nama)
		for i:=0; i < nSparepart; i++{
			if nama == A[i].nama{
				fmt.Println()
				fmt.Println(A[i].id_sparepart, A[i].nama, A[i].harga, A[i].jumlah)
				found = true
			}
		}

		if found == false{
			fmt.Println()
			fmt.Println("Maaf, sparepart dengan nama", nama, "tidak ada!")
		}
	}else if kategori == 3{
		fmt.Print("Harga Sparepart \t: ")
		fmt.Scan(&harga)
		for i:=0; i < nSparepart; i++{
			if harga == A[i].harga{
				fmt.Println()
				fmt.Println(A[i].id_sparepart, A[i].nama, A[i].harga, A[i].jumlah)
				found = true
			}
		}

		if found == false{
			fmt.Println()
			fmt.Println("Maaf, sparepart dengan harga sebesar", harga, "tidak ada!")
		}
	}else if kategori == 4{
		fmt.Print("Jumlah Sparepart \t: ")
		fmt.Scan(&jumlah)
		for i:=0; i < nSparepart; i++{
			if jumlah == A[i].jumlah{
				fmt.Println()
				fmt.Println(A[i].id_sparepart, A[i].nama, A[i].harga, A[i].jumlah)
				found = true
			}
		}

		if found == false{
			fmt.Println()
			fmt.Println("Maaf, sparepart dengan pembelian sebanyak", jumlah, "tidak ada!")
		}
	}
}

func cariPelanggan(B customer, nPelanggan int){
	var kategori, id, service int
	var nama, no_telp string
	var found bool = false

	fmt.Print("Silahkan pilih opsi \t: ")
	fmt.Scan(&kategori)
	fmt.Println()

	if kategori == 1{
		fmt.Print("ID Pelanggan \t\t: ")
		fmt.Scan(&id)
		for i:=0; i < nPelanggan; i++{
			if id == B[i].id_pelanggan{
				fmt.Println()
				fmt.Println(B[i].id_pelanggan, B[i].nama, B[i].no_telp, B[i].service)
				found = true
			}
		}

		if found == false{
			fmt.Println()
			fmt.Println("Maaf, pelanggan dengan ID", id, "tidak ada!")
		}
	}else if kategori == 2{
		fmt.Print("Nama Pelanggan \t\t: ")
		fmt.Scan(&nama)
		for i:=0; i < nPelanggan; i++{
			if nama == B[i].nama{
				fmt.Println()
				fmt.Println(B[i].id_pelanggan, B[i].nama, B[i].no_telp, B[i].service)
				found = true
			}
		}

		if found == false{
			fmt.Println()
			fmt.Println("Maaf, pelanggan dengan nama", nama, "tidak ada!")
		}
	}else if kategori == 3{
		fmt.Print("No. telepon Pelanggan \t: ")
		fmt.Scan(&no_telp)
		for i:=0; i < nPelanggan; i++{
			if no_telp == B[i].no_telp{
				fmt.Println()
				fmt.Println(B[i].id_pelanggan, B[i].nama, B[i].no_telp, B[i].service)
				found = true
			}
		}

		if found == false{
			fmt.Println()
			fmt.Println("Maaf, pelanggan dengan nomor telepon", no_telp, "tidak ada!")
		}
	}else if kategori == 4{
		fmt.Print("Jumlah Service \t\t: ")
		fmt.Scan(&service)
		for i:=0; i < nPelanggan; i++{
			if service == B[i].service{
				fmt.Println()
				fmt.Println(B[i].id_pelanggan, B[i].nama, B[i].no_telp, B[i].service)
				found = true
			}
		}

		if found == false{
			fmt.Println()
			fmt.Println("Maaf, pelanggan dengan service sebanyak", service, "tidak ada!")
		}
	}
}

func cariHistori(C histori, nHistori int){
	var kategori, id, jenis, total int
	var tgl, nama, nama_sparepart string
	var found bool = false

	fmt.Print("Silahkan pilih opsi \t: ")
	fmt.Scan(&kategori)
	fmt.Println()

	if kategori == 1{
		fmt.Print("ID Transaksi \t\t: ")
		fmt.Scan(&id)
		for i:=0; i < nHistori; i++{
			if id == C[i].id_transaksi{
				fmt.Println()
				fmt.Println(C[i].id_transaksi, C[i].tgl, C[i].nama, C[i].jenis)
				for j:=0; j < C[i].jenis; j++{
					fmt.Println(C[i].sparepart[j].nama, C[i].sparepart[j].harga, C[i].sparepart[j].jumlah)
				}
				fmt.Println(C[i].total)
				found = true
			}
		}

		if found == false{
			fmt.Println()
			fmt.Println("Maaf, transaksi dengan ID", id, "tidak ada!")
		}
	}else if kategori == 2{
		fmt.Print("Tanggal Transaksi \t\t: ")
		fmt.Scan(&tgl)
		for i:=0; i < nHistori; i++{
			if tgl == C[i].tgl{
				fmt.Println()
				fmt.Println(C[i].id_transaksi, C[i].tgl, C[i].nama, C[i].jenis)
				for j:=0; j < C[i].jenis; j++{
					fmt.Println(C[i].sparepart[j].nama, C[i].sparepart[j].harga, C[i].sparepart[j].jumlah)
				}
				fmt.Println(C[i].total)
				found = true
			}
		}

		if found == false{
			fmt.Println()
			fmt.Println("Maaf, transaksi dengan tanggal", tgl, "tidak ada!")
		}
	}else if kategori == 3{
		fmt.Print("Nama Pelanggan \t\t: ")
		fmt.Scan(&nama)
		for i:=0; i < nHistori; i++{
			if nama == C[i].nama{
				fmt.Println()
				fmt.Println(C[i].id_transaksi, C[i].tgl, C[i].nama, C[i].jenis)
				for j:=0; j < C[i].jenis; j++{
					fmt.Println(C[i].sparepart[j].nama, C[i].sparepart[j].harga, C[i].sparepart[j].jumlah)
				}
				fmt.Println(C[i].total)
				found = true
			}
		}

		if found == false{
			fmt.Println()
			fmt.Println("Maaf, transaksi dengan nama pelanggan", nama, "tidak ada!")
		}
	}else if kategori == 4{
		fmt.Print("Nama Sparepart \t\t: ")
		fmt.Scan(&nama_sparepart)
		for i:=0; i < nHistori; i++{
			for k:=0; k < C[i].jenis; k++{
				if nama_sparepart == C[i].sparepart[k].nama{
					fmt.Println()
					fmt.Println(C[i].id_transaksi, C[i].tgl, C[i].nama, C[i].jenis)
					for j:=0; j < C[i].jenis; j++{
						fmt.Println(C[i].sparepart[j].nama, C[i].sparepart[j].harga, C[i].sparepart[j].jumlah)
					}
					fmt.Println(C[i].total)
					found = true
				}
			}
		}

		if found == false{
			fmt.Println()
			fmt.Println("Maaf, transaksi dengan nama sparepart", nama_sparepart, "tidak ada!")
		}
	}else if kategori == 5{
		fmt.Print("Banyak jenis sparepart \t: ")
		fmt.Scan(&jenis)
		for i:=0; i < nHistori; i++{
			if jenis == C[i].jenis{
				fmt.Println()
				fmt.Println(C[i].id_transaksi, C[i].tgl, C[i].nama, C[i].jenis)
				for j:=0; j < C[i].jenis; j++{
					fmt.Println(C[i].sparepart[j].nama, C[i].sparepart[j].harga, C[i].sparepart[j].jumlah)
				}
				fmt.Println(C[i].total)
				found = true
			}
		}

		if found == false{
			fmt.Println()
			fmt.Println("Maaf, transaksi dengan banyak jenis sparepart", jenis, "tidak ada!")
		}
	}else if kategori == 6{
		fmt.Print("Total biaya \t\t: ")
		fmt.Scan(&total)
		for i:=0; i < nHistori; i++{
			if total == C[i].total{
				fmt.Println()
				fmt.Println(C[i].id_transaksi, C[i].tgl, C[i].nama, C[i].jenis)
				for j:=0; j < C[i].jenis; j++{
					fmt.Println(C[i].sparepart[j].nama, C[i].sparepart[j].harga, C[i].sparepart[j].jumlah)
				}
				fmt.Println(C[i].total)
				found = true
			}
		}

		if found == false{
			fmt.Println()
			fmt.Println("Maaf, transaksi dengan total biaya sebanyak", total, "tidak ada!")
		}
	}
}

func main(){
	menu()
}