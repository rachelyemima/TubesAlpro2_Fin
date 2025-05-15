package main

import (
	"fmt"
	"strings"
)

type Kegiatan struct {
	ID          int
	NamaProyek  string
	Klien       string
	Status      string // "Selesai" atau "Belum"
	DurasiHari  int
}

var dataKegiatan [100]Kegiatan
var jumlahData int

// ===== Fungsi Tambahan =====
func inputString(prompt string) string {
	var input string
	fmt.Print(prompt)
	fmt.Scanln(&input)
	return strings.TrimSpace(input)
}

func inputInt(prompt string) int {
	var input int
	fmt.Print(prompt)
	fmt.Scanln(&input)
	return input
}

func tambahKegiatan(k Kegiatan) {
	if jumlahData < len(dataKegiatan) {
		dataKegiatan[jumlahData] = k
		jumlahData++
		fmt.Println("✅ Kegiatan berhasil ditambahkan.")
	} else {
		fmt.Println("❌ Data penuh. Tidak bisa menambahkan kegiatan lagi.")
	}
}

func cariKegiatanByNama(nama string) int {
	for i := 0; i < jumlahData; i++ {
		if strings.ToLower(dataKegiatan[i].NamaProyek) == strings.ToLower(nama) {
			return i
		}
	}
	return -1
}

func binarySearchByID(id int) int {
	kiri := 0
	kanan := jumlahData - 1

	for kiri <= kanan {
		tengah := (kiri + kanan) / 2
		if dataKegiatan[tengah].ID == id {
			return tengah
		} else if id < dataKegiatan[tengah].ID {
			kanan = tengah - 1
		} else {
			kiri = tengah + 1
		}
	}
	return -1
}

func selectionSortDurasiAsc() {
	for i := 0; i < jumlahData-1; i++ {
		min := i
		for j := i + 1; j < jumlahData; j++ {
			if dataKegiatan[j].DurasiHari < dataKegiatan[min].DurasiHari {
				min = j
			}
		}
		dataKegiatan[i], dataKegiatan[min] = dataKegiatan[min], dataKegiatan[i]
	}
}

func insertionSortNamaDesc() {
	for i := 1; i < jumlahData; i++ {
		temp := dataKegiatan[i]
		j := i
		for j > 0 && strings.ToLower(dataKegiatan[j-1].NamaProyek) < strings.ToLower(temp.NamaProyek) {
			dataKegiatan[j] = dataKegiatan[j-1]
			j--
		}
		dataKegiatan[j] = temp
	}
}

func tampilkanKegiatan() {
	fmt.Println("\n--- Daftar Kegiatan ---")
	for i := 0; i < jumlahData; i++ {
		k := dataKegiatan[i]
		fmt.Printf("ID: %d | Proyek: %s | Klien: %s | Status: %s | Durasi: %d hari\n",
			k.ID, k.NamaProyek, k.Klien, k.Status, k.DurasiHari)
	}
}

func updateStatusProyek() {
	id := inputInt("Masukkan ID proyek yang ingin di-update statusnya: ")
	idx := binarySearchByID(id)
	if idx != -1 {
		fmt.Printf("Status saat ini: %s\n", dataKegiatan[idx].Status)
		newStatus := inputString("Masukkan status baru (Selesai/Belum): ")
		dataKegiatan[idx].Status = newStatus
		fmt.Println("✅ Status berhasil diperbarui.")
	} else {
		fmt.Println("❌ Proyek dengan ID tersebut tidak ditemukan. Pastikan data sudah diurutkan untuk binary search.")
	}
}

func hapusProyek() {
	id := inputInt("Masukkan ID proyek yang ingin dihapus: ")
	idx := binarySearchByID(id)
	if idx != -1 {
		for i := idx; i < jumlahData-1; i++ {
			dataKegiatan[i] = dataKegiatan[i+1]
		}
		jumlahData--
		fmt.Println("✅ Proyek berhasil dihapus.")
	} else {
		fmt.Println("❌ Proyek dengan ID tersebut tidak ditemukan. Pastikan data sudah diurutkan untuk binary search.")
	}
}


func main() {
	for {
		fmt.Println("-----------------------------------------------")
		fmt.Println("\n=== Aplikasi Manajemen Kegiatan Freelance ===")
		fmt.Println("\n-----------------------------------------------")
		fmt.Println("1. Tambah Kegiatan ➕")
		fmt.Println("2. Tampilkan Semua Kegiatan 🧾")
		fmt.Println("3. Cari Kegiatan (Nama Proyek) 🔍")
		fmt.Println("4. Urutkan (Durasi Asc / Nama Desc) 🧮")
		fmt.Println("5. Cari Kegiatan (Binary Search by ID) 🔍🧾")
		fmt.Println("6. Update Status Proyek ✏️")
		fmt.Println("7. Hapus Proyek 🗑️")
		fmt.Println("0. Keluar 🚪")
		pilihan := inputInt(" Pilih menu: ")

		if pilihan == 1 {
			var k Kegiatan
			k.ID = inputInt("ID: ")
			k.NamaProyek = inputString("Nama Proyek: ")
			k.Klien = inputString("Nama Klien: ")
			k.Status = inputString("Status (Selesai/Belum): ")
			k.DurasiHari = inputInt("Durasi (hari): ")
			tambahKegiatan(k)
		} else if pilihan == 2 {
			tampilkanKegiatan()
		} else if pilihan == 3 {
			nama := inputString("Masukkan nama proyek: ")
			idx := cariKegiatanByNama(nama)
			if idx != -1 {
				k := dataKegiatan[idx]
				fmt.Printf("Ditemukan: ID: %d | Proyek: %s | Klien: %s | Status: %s | Durasi: %d hari\n",
					k.ID, k.NamaProyek, k.Klien, k.Status, k.DurasiHari)
			} else {
				fmt.Println("❌ Data tidak ditemukan.")
			}
		} else if pilihan == 4 {
			fmt.Println("1. Urutkan Durasi (Asc)")
			fmt.Println("2. Urutkan Nama Proyek (Desc)")
			sub := inputInt("Pilih: ")
			if sub == 1 {
				selectionSortDurasiAsc()
			} else if sub == 2 {
				insertionSortNamaDesc()
			}
		} else if pilihan == 5 {
			id := inputInt("Masukkan ID: ")
			idx := binarySearchByID(id)
			if idx != -1 {
				k := dataKegiatan[idx]
				fmt.Printf("Ditemukan: ID: %d | Proyek: %s | Klien: %s | Status: %s | Durasi: %d hari\n",
					k.ID, k.NamaProyek, k.Klien, k.Status, k.DurasiHari)
			} else {
				fmt.Println("❌ ID tidak ditemukan (pastikan data sudah terurut).")
			}
		} else if pilihan == 0 {
			fmt.Println("Terima kasih!")
			return
		} else if pilihan == 6 {
			updateStatusProyek()
		} else if pilihan == 7 {
			hapusProyek()
		}
	}
}
