package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Proyek struct {
	Nama      string
	Klien     string
	Prioritas string
	Deadline  string
	Status    string
}

var daftarProyek []Proyek
var scanner = bufio.NewScanner(os.Stdin)

func main() {
	for {
		fmt.Println("\n=== Tugasin: Aplikasi Manajemen Proyek Freelance CLI ===")
		fmt.Println("1. Tambah Proyek")
		fmt.Println("2. Tampilkan Semua Proyek")
		fmt.Println("3. Sorting Proyek")
		fmt.Println("4. Cari Proyek/Klien")
		fmt.Println("5. Update Status Proyek")
		fmt.Println("6. Hapus Proyek")
		fmt.Println("7. Keluar")
		fmt.Print("Pilih menu: ")

		var pilihan int
		fmt.Scanln(&pilihan)

		switch pilihan {
		case 1:
			tambahProyek()
		case 2:
			tampilkanProyek()
		case 3:
			menuSorting()
		case 4:
			menuSearching()
		case 5:
			updateStatus()
		case 6:
			hapusProyek()
		case 7:
			fmt.Println("Terima kasih telah menggunakan Tugasin.")
			return
		default:
			fmt.Println("Pilihan tidak valid.")
		}
	}
}

func tambahProyek() {
	fmt.Print("Nama proyek: ")
	scanner.Scan()
	nama := scanner.Text()

	fmt.Print("Nama klien: ")
	scanner.Scan()
	klien := scanner.Text()

	fmt.Print("Prioritas (tinggi/sedang/rendah): ")
	scanner.Scan()
	prioritas := scanner.Text()

	fmt.Print("Deadline (contoh: 2025-06-30): ")
	scanner.Scan()
	deadline := scanner.Text()

	fmt.Print("Status (belum mulai/dikerjakan/selesai): ")
	scanner.Scan()
	status := scanner.Text()

	proyek := Proyek{nama, klien, prioritas, deadline, status}
	daftarProyek = append(daftarProyek, proyek)
	fmt.Println("Proyek berhasil ditambahkan!")
}

func tampilkanProyek() {
	fmt.Println("\n--- Daftar Proyek ---")
	for i, p := range daftarProyek {
		fmt.Printf("[%d] %s | Klien: %s | Prioritas: %s | Deadline: %s | Status: %s\n",
			i+1, p.Nama, p.Klien, p.Prioritas, p.Deadline, p.Status)
	}
	if len(daftarProyek) == 0 {
		fmt.Println("Belum ada proyek.")
	}
}

func menuSorting() {
	fmt.Println("\n--- Sorting ---")
	fmt.Println("1. Berdasarkan Prioritas (selection sort)")
	fmt.Println("2. Berdasarkan Deadline (insertion sort)")
	fmt.Println("3. Berdasarkan Status (selection sort)")
	fmt.Print("Pilih jenis sorting: ")
	var opsi int
	fmt.Scanln(&opsi)

	switch opsi {
	case 1:
		selectionSortPrioritas()
	case 2:
		insertionSortDeadline()
	case 3:
		selectionSortStatus()
	default:
		fmt.Println("Opsi tidak valid.")
	}
	fmt.Println("Proyek berhasil diurutkan.")
}

func prioritasValue(p string) int {
	switch strings.ToLower(p) {
	case "tinggi":
		return 3
	case "sedang":
		return 2
	case "rendah":
		return 1
	default:
		return 0
	}
}

func statusValue(s string) int {
	switch strings.ToLower(s) {
	case "belum mulai":
		return 1
	case "dikerjakan":
		return 2
	case "selesai":
		return 3
	default:
		return 0
	}
}

func selectionSortPrioritas() {
	n := len(daftarProyek)
	for i := 0; i < n-1; i++ {
		max := i
		for j := i + 1; j < n; j++ {
			if prioritasValue(daftarProyek[j].Prioritas) > prioritasValue(daftarProyek[max].Prioritas) {
				max = j
			}
		}
		daftarProyek[i], daftarProyek[max] = daftarProyek[max], daftarProyek[i]
	}
}

func insertionSortDeadline() {
	for i := 1; i < len(daftarProyek); i++ {
		key := daftarProyek[i]
		j := i - 1
		for j >= 0 && daftarProyek[j].Deadline > key.Deadline {
			daftarProyek[j+1] = daftarProyek[j]
			j--
		}
		daftarProyek[j+1] = key
	}
}

func selectionSortStatus() {
	n := len(daftarProyek)
	for i := 0; i < n-1; i++ {
		min := i
		for j := i + 1; j < n; j++ {
			if statusValue(daftarProyek[j].Status) < statusValue(daftarProyek[min].Status) {
				min = j
			}
		}
		daftarProyek[i], daftarProyek[min] = daftarProyek[min], daftarProyek[i]
	}
}

func menuSearching() {
	fmt.Println("\n--- Searching ---")
	fmt.Println("1. Cari berdasarkan nama proyek (linear search)")
	fmt.Println("2. Cari berdasarkan nama klien (linear search)")
	fmt.Print("Pilih jenis pencarian: ")
	var opsi int
	fmt.Scanln(&opsi)

	fmt.Print("Masukkan kata kunci: ")
	scanner.Scan()
	kunci := scanner.Text()

	switch opsi {
	case 1:
		linearSearch("proyek", kunci)
	case 2:
		linearSearch("klien", kunci)
	default:
		fmt.Println("Opsi tidak valid.")
	}
}

func linearSearch(kategori, kunci string) {
	ditemukan := false
	for _, p := range daftarProyek {
		if (kategori == "proyek" && strings.Contains(strings.ToLower(p.Nama), strings.ToLower(kunci))) ||
			(kategori == "klien" && strings.Contains(strings.ToLower(p.Klien), strings.ToLower(kunci))) {
			fmt.Printf("Ditemukan: %s | Klien: %s | Prioritas: %s | Deadline: %s | Status: %s\n",
				p.Nama, p.Klien, p.Prioritas, p.Deadline, p.Status)
			ditemukan = true
		}
	}
	if !ditemukan {
		fmt.Println("Tidak ditemukan.")
	}
}

func updateStatus() {
	fmt.Print("Masukkan nama proyek yang ingin diupdate: ")
	scanner.Scan()
	nama := scanner.Text()

	for i := range daftarProyek {
		if strings.EqualFold(daftarProyek[i].Nama, nama) {
			fmt.Print("Masukkan status baru: ")
			scanner.Scan()
			daftarProyek[i].Status = scanner.Text()
			fmt.Println("Status berhasil diperbarui.")
			return
		}
	}
	fmt.Println("Proyek tidak ditemukan.")
}

func hapusProyek() {
	fmt.Print("Masukkan nama proyek yang ingin dihapus: ")
	scanner.Scan()
	nama := scanner.Text()

	for i := range daftarProyek {
		if strings.EqualFold(daftarProyek[i].Nama, nama) {
			daftarProyek = append(daftarProyek[:i], daftarProyek[i+1:]...)
			fmt.Println("Proyek berhasil dihapus.")
			return
		}
	}
	fmt.Println("Proyek tidak ditemukan.")
}