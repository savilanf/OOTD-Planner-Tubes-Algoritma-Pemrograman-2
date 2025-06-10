package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	var koleksi Koleksi
	var fav Favorit
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Println("\n=== MENU UTAMA ===")
		fmt.Println("1. Lihat Koleksi")
		fmt.Println("2. Tambah Pakaian")
		fmt.Println("3. Hapus Pakaian")
		fmt.Println("4. Edit Pakaian")
		fmt.Println("5. Kombinasi Outfit")
		fmt.Println("6. Cari Pakaian by warna")
		fmt.Println("7. Urutkan Pakaian")
		fmt.Println("8. Rekomendasi Outfit")
		fmt.Println("0. Exit")
		fmt.Print("Pilih menu: ")
		var choice int
		fmt.Scan(&choice)
		clearBuffer()

		switch choice {
		case 1:
			fmt.Println("\n=== MENU LIHAT KOLEKSI ===")
			lihatKoleksi(&koleksi, &fav)
		case 2:
			fmt.Println("\n=== MENU TAMBAH PAKAIAN ===")
			tambahPakaian(&koleksi, reader)
		case 3:
			fmt.Println("\n=== MENU HAPUS PAKAIAN ===")
			hapusPakaian(&koleksi, reader)
		case 4:
			fmt.Println("\n=== MENU EDIT PAKAIAN ===")
			editPakaian(&koleksi, reader)
		case 5:
			fmt.Println("\n=== MENU KOMBINASI OUTFIT ===")
			kombinasiOutfit(&koleksi, &fav, reader)
		case 6:
			fmt.Println("\n=== MENU CARI PAKAIAN BY WARNA ===")
			cariPakaian(&koleksi, reader)
		case 7:
			fmt.Println("\n=== MENU URUTKAN PAKAIAN ===")
			urutkanPakaian(&koleksi)
		case 8:
			fmt.Println("\n=== MENU REKOMENDASI OUTFIT ===")
			rekomendasiOutfit(&koleksi, reader)
		case 0:
			fmt.Println("Terima kasih telah menggunakan OOTD Planner!")
			return
		default:
			fmt.Println("Oops, pilihan kamu tidak valid 😢❌")
		}
	}
}

func clearBuffer() {
	var temp string
	fmt.Scanln(&temp)
}

func readInput(pesan string, reader *bufio.Reader) string {
	fmt.Print(pesan)
	input, _ := reader.ReadString('\n')
	return strings.TrimSpace(input)
}

func compareStrings(a, b string) bool {
	return strings.ToLower(a) == strings.ToLower(b)
}

type Pakaian struct {
	ID         int
	Nama       string
	Jenis      string
	Warna      string
	Kategori   string
	Formalitas string
	Cuaca      string
	Acara      string
	Date       time.Time
}

type Koleksi struct {
	count   int
	Pakaian [20]Pakaian
}

type Kombinasi struct {
	Date     time.Time
	Atasan   Pakaian
	Bawahan  Pakaian
	Footwear Pakaian
}

type Favorit struct {
	count      int
	Kombinasi [20]Kombinasi
}

func cekValidID(ID int) bool {
	return ID >= 1 && ID <= 20
}

func cekUnikID(ID int, k *Koleksi) bool {
    for i := 0; i < k.count; i++ {
        if k.Pakaian[i].ID == ID {
            return true
        }
    }
    return false
}

func validJenis(jenis string) bool {
	jenis = strings.ToLower(jenis)
	return jenis == "atasan" || jenis == "bawahan" || jenis == "outer" || jenis == "dress" || jenis == "footwear" || jenis == "aksesoris"
}

func validFormalitas(formalitas string) bool {
    return formalitas == "1" || formalitas == "2" || formalitas == "3"
}

func validKategori(kategori string) bool {
	kategori = strings.ToLower(kategori)
	return kategori == "casual" || kategori == "formal" || kategori == "sporty"
}

func tambahPakaian(k *Koleksi, reader *bufio.Reader) {
	var p Pakaian
	if k.count >= 20 {
		fmt.Println("😿 Yahh, koleksi kamu udah penuh! tidak muat lagi... coba hapus yang lama dulu yaa 🧺✨")
		return
	}
	validID := false
	for !validID {
		fmt.Print("Masukkan ID (1 - 20): ")
		fmt.Scan(&p.ID)
		clearBuffer()
		if cekValidID(p.ID) {
			if !cekUnikID(p.ID, k) {
				validID = true
			} else {
				fmt.Println("ID udah digunain, masukkin ID lain yaa")
			}
		} else {
			fmt.Println("ID tidak valid, harus dari 1 sampai 20")
		}
	}

	p.Nama = readInput("Nama: ", reader)

	jenisValid := false
	for !jenisValid {
		p.Jenis = readInput("Jenis (atasan/bawahan/outer/dress/footwear/aksesoris): ", reader)
		if validJenis(p.Jenis) {
			jenisValid = true
		} else {
			fmt.Println("Jenis tidak valid, coba lagi yaa")
		}
	}

	p.Warna = readInput("Warna: ", reader)

	formalValid := false
	for !formalValid {
		p.Formalitas = readInput("Tingkat formalitas, angkanya saja (tinggi = 3 / sedang = 2 / rendah = 1): ", reader)
		if validFormalitas(p.Formalitas) {
			formalValid = true
		} else {
			fmt.Println("Formalitas tidak valid, coba lagi yaa")
		}
	}

	kategoriValid := false
	for !kategoriValid {
		p.Kategori = readInput("Kategori (formal/casual/sporty): ", reader)
		if validKategori(p.Kategori) {
			kategoriValid = true
		} else {
			fmt.Println("Kategori tidak valid, coba lagi yaa")
		}
	}

	p.Cuaca = readInput("Cocok untuk cuaca: ", reader)
	p.Acara = readInput("Cocok untuk acara: ", reader)
	p.Date = time.Now()
	k.Pakaian[k.count] = p
	k.count++
	fmt.Println("Yeay, pakaian berhasil ditambahkan ke koleksi kamu! 🎉😊")
}

func lihatKoleksi(k *Koleksi, f *Favorit) {
	fmt.Println("1. Lihat koleksi pakaian")
	fmt.Println("2. Lihat histori kombinasi outfit")
	fmt.Print("Pilih: ")
	var choice int
	fmt.Scan(&choice)
	clearBuffer()

	if choice == 1 {
		if k.count == 0 {
			fmt.Println("Wah, koleksi pakaian kamu masih kosong. Yuk tambah koleksi biar makin kece! 😎✨")
			return
		}
		fmt.Println("\nKoleksi:")
		layoutFormat := "02-01-2006 15:04:05"
		for i := 0; i < k.count; i++ {
			fmt.Printf("- ID: %v\n", k.Pakaian[i].ID)
			fmt.Printf("- Nama: %v\n", k.Pakaian[i].Nama)
			fmt.Printf("- Jenis: %v\n", k.Pakaian[i].Jenis)
			fmt.Printf("- Warna: %v\n", k.Pakaian[i].Warna)
			fmt.Printf("- Tingkat formalitas: %v\n", k.Pakaian[i].Formalitas)
			fmt.Printf("- Kategori: %v\n", k.Pakaian[i].Kategori)
			fmt.Printf("- Cocok untuk acara: %v\n", k.Pakaian[i].Acara)
			fmt.Printf("- Cocok untuk cuaca: %v\n", k.Pakaian[i].Cuaca)
			fmt.Printf("- Tanggal ditambahkan: %v\n", k.Pakaian[i].Date.Format(layoutFormat))
			fmt.Println("----------------------------------------------")
		}
	} else if choice == 2 {
		if f.count == 0 {
			fmt.Println("Wah, kombinasi outfit kamu masih kosong. Yuk buat outfit biar makin kece! 😎✨")
			return
		}
		fmt.Println("\nKombinasi:")
		layoutFormat := "02-01-2006 15:04:05"
		for i := 0; i < f.count; i++ {
			date := f.Kombinasi[i].Date
			atasan := f.Kombinasi[i].Atasan
			bawahan := f.Kombinasi[i].Bawahan
			footwear := f.Kombinasi[i].Footwear
			fmt.Printf("- Tanggal dibuat: %v\n", date.Format(layoutFormat))
			fmt.Printf("- Atasan: %v (%v)\n", atasan.Nama, atasan.Warna)
			fmt.Printf("- Bawahan: %v (%v)\n", bawahan.Nama, bawahan.Warna)
			fmt.Printf("- Footwear: %v (%v)\n", footwear.Nama, footwear.Warna)
			fmt.Println("----------------------------------------------")
		}
	} else {
		fmt.Println("Oops, pilihan kamu tidak valid 😢❌")
	}
}

func editPakaian(k *Koleksi, reader *bufio.Reader) {
	if k.count == 0 {
		fmt.Println("Wah, koleksi pakaian kamu masih kosong. Yuk tambah koleksi biar makin kece! 😎✨")
		return
	}
	var ID int
	fmt.Print("ID pakaian yang ingin di edit: ")
	fmt.Scan(&ID)
	clearBuffer()

	found := false
	i := 0
	for i < k.count && !found {
		if k.Pakaian[i].ID == ID {
			found = true
			fmt.Printf("Nama: %v | Warna: %v | Jenis: %v | Formalitas: %v | Kategori : %v | Cocok untuk cuaca: %v | Cocok untuk acara: %v\n", 
			k.Pakaian[i].Nama, k.Pakaian[i].Warna, k.Pakaian[i].Jenis, k.Pakaian[i].Formalitas, k.Pakaian[i].Kategori, k.Pakaian[i].Cuaca, k.Pakaian[i].Acara)
			var choice int
			fmt.Println("1. Nama")
			fmt.Println("2. Jenis")
			fmt.Println("3. Warna")
			fmt.Println("4. Formalitas")
			fmt.Println("5. Kategori")
			fmt.Println("6. Cuaca")
			fmt.Println("7. Acara")
			fmt.Print("Pilih mana yang mau diedit: ")
			fmt.Scan(&choice)
			clearBuffer()

			if choice == 1 {
				k.Pakaian[i].Nama = readInput("Nama baru: ", reader)
			} else if choice == 2 {
				jenisValid := false
				for !jenisValid {
					k.Pakaian[i].Jenis = readInput("Jenis baru (atasan/bawahan/outer/dress/footwear/aksesoris): ", reader)
					if validJenis(k.Pakaian[i].Jenis) {
						jenisValid = true
					} else {
						fmt.Println("Jenis tidak valid, coba lagi yaa")
					}
				}
			} else if choice == 3 {
				k.Pakaian[i].Warna = readInput("Warna baru: ", reader)
			} else if choice == 4 {
				formalitasValid := false
				for !formalitasValid {
					k.Pakaian[i].Formalitas = readInput("Formalitas baru, angkanya saja (tinggi = 3 / sedang = 2 / rendah = 1): ", reader)
					if validFormalitas(k.Pakaian[i].Formalitas) {
						formalitasValid = true
					} else {
						fmt.Println("Formalitas tidak valid, coba lagi yaa")
					}
				}
			} else if choice == 5 {
				kategoriValid := false
				for !kategoriValid {
					k.Pakaian[i].Kategori = readInput("Kategori baru (formal/casual/sporty): ", reader)
					if validKategori(k.Pakaian[i].Kategori) {
						kategoriValid = true
					} else {
						fmt.Println("Kategori tidak valid, coba lagi yaa")
					}
				}
			} else if choice == 6 {
				k.Pakaian[i].Cuaca = readInput("Cuaca baru: ", reader)
			} else if choice == 7 {
				k.Pakaian[i].Acara = readInput("Acara baru: ", reader)
			} else {
				fmt.Println("Oops, pilihan kamu tidak valid 😢❌")
			}
		}
		i++
	}
	if found {
		fmt.Println("Yeay, pakaian berhasil diedit sesuai keinginan kamu! ✨😄")
	} else {
		fmt.Println("Maaf, ID yang kamu cari tidak ketemu 😢🔍") 
	}
}

func hapusPakaian(k *Koleksi, reader *bufio.Reader) {
	if k.count == 0 {
		fmt.Println("Wah, koleksi pakaian kamu masih kosong. Yuk tambah koleksi biar makin kece! 😎✨")
		return
	}
	var ID int
	fmt.Print("ID pakaian yang ingin dihapus: ")
	fmt.Scan(&ID)
	clearBuffer()

	found := false
	idx := 0
	i := 0
	for i < k.count && !found {
		if k.Pakaian[i].ID == ID {
			found = true
			fmt.Printf("Nama: %v | Warna: %v | Jenis: %v | Formalitas: %v | Kategori : %v | Cocok untuk cuaca: %v | Cocok untuk acara: %v\n", 
			k.Pakaian[i].Nama, k.Pakaian[i].Warna, k.Pakaian[i].Jenis, k.Pakaian[i].Formalitas, k.Pakaian[i].Kategori, k.Pakaian[i].Cuaca, k.Pakaian[i].Acara)
			idx = i
		}
		i++
	}
	if found {
		for i := idx; i < k.count-1; i++ {
			k.Pakaian[i] = k.Pakaian[i+1]
		}
		k.count--
		k.Pakaian[k.count] = Pakaian{}
		fmt.Println("Yeay, pakaian berhasil dihapus dari koleksi kamu! 👍😃")
	} else {
		fmt.Println("Maaf, ID yang kamu cari tidak ketemu 😢🔍") 
	}
}

func kombinasiOutfit(k *Koleksi, f *Favorit, reader *bufio.Reader) {
	if k.count == 0 {
		fmt.Println("Wah, koleksi pakaian kamu masih kosong. Yuk tambah koleksi biar makin kece! 😎✨")
		return
	}

	if k.count < 3 {
		fmt.Println("Belum bisa mengkombinasikan outfit, minimal jenis atasan, bawahan, dan footwear ada di koleksi 👍😃")
		return
	}

	if f.count >= 20 {
		fmt.Println("Sudah tidak bisa mengkombinasikan outfit 😥💔")
		return
	}

	for i := 0; i < k.count; i++ {
		fmt.Printf("ID: %v | Nama: %v | Jenis: %v\n", k.Pakaian[i].ID, k.Pakaian[i].Nama, k.Pakaian[i].Jenis)
	}
	var idAtasan, idBawahan, idFootwear int
	fmt.Print("Masukkan ID atasan: ")
	fmt.Scan(&idAtasan)
	fmt.Print("Masukkan ID bawahan: ")
	fmt.Scan(&idBawahan)
	fmt.Print("Masukkan ID footwear: ")
	fmt.Scan(&idFootwear)
	clearBuffer()

	var atasan, bawahan, footwear Pakaian
	foundAtasan, foundBawahan, foundFootwear := false, false, false
	for i := 0; i < k.count; i++ {
		p := k.Pakaian[i]
		if p.ID == idAtasan && compareStrings(p.Jenis, "atasan") {
			atasan = p
			foundAtasan = true
		}
		if p.ID == idBawahan && compareStrings(p.Jenis, "bawahan") {
			bawahan = p
			foundBawahan = true
		}
		if p.ID == idFootwear && compareStrings(p.Jenis, "footwear") {
			footwear = p
			foundFootwear = true
		}
	}
	date := time.Now()
	if foundAtasan && foundBawahan && foundFootwear {
		fmt.Println("Outfit minimalis berhasil dibuat dan ditambahkan! 🎉😊")

		f.Kombinasi[f.count] = Kombinasi{
			Date:     date,
			Atasan:   atasan,
			Bawahan:  bawahan,
			Footwear: footwear,
		}
		f.count++
	} else {
		fmt.Println("Tidak bisa mengkombinasikan outfit, ID tidak ditemukan atau jenis tidak cocok 😿")
	}
}

func cariPakaian(k *Koleksi, reader *bufio.Reader) {
	if k.count == 0 {
		fmt.Println("Wah, koleksi pakaian kamu masih kosong. Yuk tambah koleksi biar makin kece! 😎✨")
		return
	}
	warna := readInput("Masukkan warna: ", reader) // sequential search
	found := false
	for i := 0; i < k.count; i++ {
		if compareStrings(k.Pakaian[i].Warna, warna) {
			if !found {
				fmt.Println("\nHasil pencarian: ")
				found = true
			}
			fmt.Printf("- ID: %v\n", k.Pakaian[i].ID)
				fmt.Printf("- Nama: %v (%v) | Jenis: %v\n", k.Pakaian[i].Nama, k.Pakaian[i].Warna, k.Pakaian[i].Jenis)
			fmt.Println("----------------------------------------------")
		}
	}
	if !found {
		fmt.Printf("Tidak ada pakaian warna %v 😿\n", warna)
	}
}

func urutkanPakaian(k *Koleksi) {
	if k.count == 0 {
		fmt.Println("Wah, koleksi pakaian kamu masih kosong. Yuk tambah koleksi biar makin kece! 😎✨")
		return
	}

	fmt.Println("1. Urutkan berdasarkan tingkat formalitas (tinggi ke rendah)") // selection sort
	fmt.Println("2. Urutkan berdasarkan terakhir ditambahkan (terbaru ke terlama)") // insertion sort
	fmt.Print("Pilih: ")
	var choice int
	fmt.Scan(&choice)
	clearBuffer()

	if choice == 1 {
		for i := 0; i < k.count-1; i++ {
			idx := i
			for j := i + 1; j < k.count; j++ {
				if k.Pakaian[j].Formalitas > k.Pakaian[idx].Formalitas {
					idx = j
				}
			}
			k.Pakaian[i], k.Pakaian[idx] = k.Pakaian[idx], k.Pakaian[i]
		}
		fmt.Println("Yey, pakaian berhasil diurutkan berdasarkan formalitas 🎉 silakan lihat di menu 'Lihat Koleksi' → 'Lihat koleksi pakaian'")
	} else if choice == 2 {
		for i := 1; i < k.count; i++ {
			temp := k.Pakaian[i]
			j := i - 1
			for j >= 0 && temp.Date.After(k.Pakaian[j].Date) {
				k.Pakaian[j+1] = k.Pakaian[j]
				j--
			}
			k.Pakaian[j+1] = temp
		}
		fmt.Println("Yey, pakaian berhasil diurutkan berdasarkan terakhir ditambah 🎉 silakan lihat di menu 'Lihat Koleksi' → 'Lihat koleksi pakaian'")
	} else {
		fmt.Println("Oops, pilihan kamu tidak valid 😢❌")
	}
}

func rekomendasiOutfit(k *Koleksi, reader *bufio.Reader) {
	fmt.Println("1. Rekomendasi dari kami berdasarkan cuaca")
	fmt.Println("2. Rekomendasi dari kami berdasarkan acara")
	fmt.Println("3. Rekomendasi dari koleksi anda berdasarkan acara")
	fmt.Println("4. Rekomendasi dari koleksi anda berdasarkan cuaca")
	fmt.Print("Pilih: ")
	var choice int
	fmt.Scan(&choice)
	clearBuffer()

	if choice == 1 {
		cuaca := strings.ToLower(readInput("Masukkan cuaca (cerah/panas/dingin/hujan/mendung): ", reader))
		fmt.Println("\nRekomendasi outfit:")
		if cuaca == "panas" || cuaca == "cerah" {
			fmt.Println(" - kaos\n - kemeja tipis\n - bawahan panjang/pendek\n - sandal kodok")
		} else if cuaca == "hujan" || cuaca == "dingin" || cuaca == "mendung" {
			fmt.Println(" - jaket\n - hoodie\n - sweater\n - bawahan panjang\n - sepatu boots")
		} else {
			fmt.Printf("Tidak ada rekomendasi outfit untuk cuaca %v 😞\n", cuaca) 
		}
	} else if choice == 2 {
		acara := strings.ToLower(readInput("Masukkan acara (rapat/formal/hangout/pesta/olahraga/santai): ", reader))
		fmt.Println("\nRekomendasi outfit:")
		if acara == "formal" || acara == "rapat" {
			fmt.Println(" - kemeja/jas formal\n - bawahan formal\n - sepatu pantofel\n - batik")
		} else if acara == "hangout" || acara == "santai" {
			fmt.Println(" - kaos\n - bawahan pendek\n - sneakers")
		} else if acara == "olahraga" {
			fmt.Println(" - kaos olahraga\n - celana training\n - sepatu olahraga")
		} else if acara == "pesta" {
			fmt.Println("Perempuan:\n - gaun\n - sepatu hak\n - batik")
			fmt.Println("Laki laki:\n - jas\n - celana formal\n - batik")
		} else {
			fmt.Printf("Tidak ada rekomendasi outfit untuk acara %v 😞\n", acara) 
		}
	} else if choice == 3 {
		if k.count == 0 {
			fmt.Println("Wah, koleksi pakaian kamu masih kosong. Yuk tambah koleksi biar makin kece! 😎✨")
			return
		}
		acara := readInput("Masukkan acara: ", reader)
		fmt.Println("\nRekomendasi outfit: ")
		ada := false
		for i := 0; i < k.count; i++ {
			if compareStrings(k.Pakaian[i].Acara, acara) {
				if !ada {
					ada = true
				}
				fmt.Printf("- ID: %v\n", k.Pakaian[i].ID)
				fmt.Printf("- Nama: %v %v (%v) | Jenis: %v\n", k.Pakaian[i].Nama, k.Pakaian[i].Warna, k.Pakaian[i].Kategori, k.Pakaian[i].Jenis)
				fmt.Println("----------------------------------------------")
			}
		}
		if !ada {
			fmt.Printf("Tidak ada rekomendasi outfit untuk acara %v 😞\n", acara)
		}
	} else if choice == 4 {
		if k.count == 0 {
			fmt.Println("Wah, koleksi pakaian kamu masih kosong. Yuk tambah koleksi biar makin kece! 😎✨")
			return
		}
		cuaca := readInput("Masukkan cuaca: ", reader)
		fmt.Println("\nRekomendasi outfit: ")
		ada := false
		for i := 0; i < k.count; i++ {
			if compareStrings(k.Pakaian[i].Cuaca, cuaca) {
				if !ada {
					ada = true
				}
				fmt.Printf("- ID: %v\n", k.Pakaian[i].ID)
				fmt.Printf("- Nama: %v %v (%v) | Jenis: %v\n", k.Pakaian[i].Nama, k.Pakaian[i].Warna, k.Pakaian[i].Kategori, k.Pakaian[i].Jenis)
				fmt.Println("----------------------------------------------")
			}
		}
		if !ada {
			fmt.Printf("Tidak ada rekomendasi outfit untuk cuaca %v 😞\n", cuaca)
		}
	} else {
		fmt.Println("Oops, pilihan kamu tidak valid 😢❌")
	}
}