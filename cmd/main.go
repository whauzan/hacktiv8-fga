package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
)

type data struct {
	nama      string
	alamat    string
	pekerjaan string
	alasan    string
}

type listData struct {
	listNama []string
	listSiswa []data
	absen int
}

type listName struct {
	listNama []string
}

type impl interface {
	checkAbsen() bool
	PrintSiswa() string
}

type generate interface {
	generateBiodata() []data
}

func (ls listData) checkAbsen() bool {
	return ls.absen <= len(ls.listSiswa) && ls.absen > 0
}

func (ls listData) PrintSiswa() string {
	d := ls.listSiswa
	absen := ls.absen - 1
	return fmt.Sprintf("Nama \t\t\t\t: %v\nAlamat \t\t\t\t: %v\nPekerjaan \t\t\t: %v\nAlasan memilih kelas Golang \t: %v", d[absen].nama, d[absen].alamat, d[absen].pekerjaan, d[absen].alasan)
}

func (ls listName) generateBiodata() []data {
	listNama := ls.listNama
	alamat := []string{"Bandung", "Jakarta", "Semarang", "Pemalang"}
	alasan := []string{"Menambah skill", "Ingin belajar aja", "Keperluan kantor"}
	pekerjaan := []string{"Software Engineer", "Mahasiswa", "IT Support"}
	generate := make([]data, 0)

	var s data
	for _, value := range listNama {
		s.nama = value
		s.alamat = alamat[rand.Intn(len(alamat))]
		s.pekerjaan = pekerjaan[rand.Intn(len(pekerjaan))]
		s.alasan = alasan[rand.Intn(len(alasan))]
		generate = append(generate, s)
	}

	return generate
}

func main() {
	fmt.Printf("\n\nHello! Wahyu here ^-^\n\n")

	args, _ := strconv.Atoi(os.Args[1])
	listNama := []string{"Wahyu", "Hauzan", "Rafi"}
	var generateBiodata generate = listName{listNama: listNama}
	var dataKelas impl = listData{listNama: listNama, listSiswa: generateBiodata.generateBiodata(), absen: args}

	if dataKelas.checkAbsen() {
		fmt.Println(dataKelas.PrintSiswa())
		os.Exit(0)
	} else {
		err := fmt.Errorf("siswa dengan absen %d tidak ada", args)
		fmt.Println(err)
	}

}
