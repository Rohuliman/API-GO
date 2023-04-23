package main

import (
	"net/http"

	"github.com/Rohuliman/API-GO/controllers/mahasiswacontroller"
)


func main()  {
	
	http.HandleFunc("/", mahasiswacontroller.Index)
	http.HandleFunc("/mahasiswa", mahasiswacontroller.Index)
	http.HandleFunc("/mahasiswa/index", mahasiswacontroller.Index)
	http.HandleFunc("/mahasiswa/tambah", mahasiswacontroller.Tambah)
	http.HandleFunc("/mahasiswa/edit", mahasiswacontroller.Edit)
	http.HandleFunc("/mahasiswa/hapus", mahasiswacontroller.Hapus)

	http.ListenAndServe(":3000", nil)
}
