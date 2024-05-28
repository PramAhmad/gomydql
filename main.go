package main

import (
	"bufio"
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

type Mahasiswa struct {
	NPM          string
	Nama         string
	Alamat       string
	JenisKelamin string
	GolDarah     string
}

var db *sql.DB

func main() {

	var err error
	db, err = sql.Open("mysql", "root:Tasikmalaya123..@tcp(127.0.0.1:3306)/akademik")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	for {
		fmt.Println("\nMenu:")
		fmt.Println("1. Tambah Data Mahasiswa")
		fmt.Println("2. Tampilkan Data Mahasiswa")
		fmt.Println("3. Update Data Mahasiswa")
		fmt.Println("4. Hapus Data Mahasiswa")
		fmt.Println("5. Keluar")
		fmt.Print("Pilih menu: ")
		var choice string
		fmt.Scanln(&choice)

		switch choice {
		case "1":
			addMahasiswa()
		case "2":
			showMahasiswa()
		case "3":
			updateMahasiswa()
		case "4":
			deleteMahasiswa()
		case "5":
			fmt.Println("Terima kasih!")
			return
		default:
			fmt.Println("Pilihan tidak valid.")
		}
	}
}

func addMahasiswa() {
	var m Mahasiswa
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("NPM: ")
	scanner.Scan()
	m.NPM = scanner.Text()
	fmt.Print("Nama: ")
	scanner.Scan()
	m.Nama = scanner.Text()
	fmt.Print("Alamat: ")
	scanner.Scan()
	m.Alamat = scanner.Text()
	fmt.Print("Jenis Kelamin: ")
	scanner.Scan()
	m.JenisKelamin = scanner.Text()
	fmt.Print("Golongan Darah: ")
	scanner.Scan()
	m.GolDarah = scanner.Text()
	err := createMahasiswa(m)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Data mahasiswa berhasil ditambahkan")
}

func showMahasiswa() {
	mahasiswas, err := readMahasiswa()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Data Mahasiswa:")
	for _, m := range mahasiswas {
		fmt.Printf("NPM: %s, Nama: %s, Alamat: %s, Jenis Kelamin: %s, Gol Darah: %s\n", m.NPM, m.Nama, m.Alamat, m.JenisKelamin, m.GolDarah)
	}
}

func updateMahasiswa() {
	var m Mahasiswa
	fmt.Print("Masukkan NPM mahasiswa yang akan diupdate: ")
	fmt.Scanln(&m.NPM)
	var newData Mahasiswa
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Nama Baru: ")
	scanner.Scan()
	newData.Nama = scanner.Text()
	fmt.Print("Alamat Baru: ")
	scanner.Scan()
	newData.Alamat = scanner.Text()
	fmt.Print("Jenis Kelamin Baru: ")
	scanner.Scan()
	newData.JenisKelamin = scanner.Text()
	fmt.Print("Golongan Darah Baru: ")
	scanner.Scan()
	newData.GolDarah = scanner.Text()
	err := updateMahasiswaData(m.NPM, newData)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Data mahasiswa berhasil diperbarui")
}

func deleteMahasiswa() {
	var npm string
	fmt.Print("Masukkan NPM mahasiswa yang akan dihapus: ")
	fmt.Scanln(&npm)
	err := deleteMahasiswaData(npm)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Data mahasiswa berhasil dihapus")
}

func createMahasiswa(m Mahasiswa) error {
	_, err := db.Exec("INSERT INTO mahasiswa(npm, nama, alamat, jenis_kelamin, gol_darah) VALUES(?, ?, ?, ?, ?)", m.NPM, m.Nama, m.Alamat, m.JenisKelamin, m.GolDarah)
	return err
}

func readMahasiswa() ([]Mahasiswa, error) {
	rows, err := db.Query("SELECT * FROM mahasiswa")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var mahasiswas []Mahasiswa
	for rows.Next() {
		var m Mahasiswa
		if err := rows.Scan(&m.NPM, &m.Nama, &m.Alamat, &m.JenisKelamin, &m.GolDarah); err != nil {
			return nil, err
		}
		mahasiswas = append(mahasiswas, m)
	}
	return mahasiswas, nil
}

func updateMahasiswaData(npm string, m Mahasiswa) error {
	_, err := db.Exec("UPDATE mahasiswa SET nama=?, alamat=?, jenis_kelamin=?, gol_darah=? WHERE npm=?", m.Nama, m.Alamat, m.JenisKelamin, m.GolDarah, npm)
	return err
}

func deleteMahasiswaData(npm string) error {
	_, err := db.Exec("DELETE FROM mahasiswa WHERE npm=?", npm)
	return err
}
