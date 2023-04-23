package mahasiswacontroller

import (
	"html/template"
	"net/http"
	"strconv"

	"github.com/Rohuliman/API-GO/libraries"
	"github.com/Rohuliman/API-GO/models"

	"github.com/Rohuliman/API-GO/entities"
)

var validation = libraries.NewValidation()
var detailModel = models.NewDetailModel()
var mahasiswaModel = models.NewMahasiswaModel() 
var jurusanModel = models.NewJurusanModel()
var hobiModel = models.NewHobiModel()


func Index(response http.ResponseWriter, request *http.Request) {
	detail, _ := detailModel.FindAll()
	mahasiswa, _ := mahasiswaModel.Find()
	jurusan, _ := jurusanModel.Find()
	hobi, _ := hobiModel.Find()
	data := map[string]interface{}{
		"detail": detail,
		"mahasiswa": mahasiswa,
		"jurusan": jurusan,
		"hobi": hobi,
	}
	temp, err := template.ParseFiles("views/mahasiswa/index.html")
	if err != nil {
		panic(err)
	}
	temp.Execute(response, data)



}

func Tambah(response http.ResponseWriter, request *http.Request) {

	if request.Method == http.MethodGet {
		temp, err := template.ParseFiles("views/mahasiswa/tambah.html")
	if err != nil {
		panic(err)
	}
	temp.Execute(response, nil)
	} else if request.Method == http.MethodPost {

		request.ParseForm()

		var detail entities.Detail
		detail.Nama = request.Form.Get("nama")
		detail.Usia = request.Form.Get("usia")
		detail.Gender = request.Form.Get("gender")
		detail.TanggalRegistrasi = request.Form.Get("tanggal_registrasi")
		detail.Jurusan = request.Form.Get("jurusan")
		detail.Hobi = request.Form.Get("hobi")

		var mahasiswa entities.Mahasiswa
		mahasiswa.Nama = request.Form.Get("nama")
		mahasiswa.Usia = request.Form.Get("usia")
		mahasiswa.Gender = request.Form.Get("gender")
		mahasiswa.TanggalRegistrasi = request.Form.Get("tanggal_registrasi")
		
		var jurusan entities.Jurusan
		jurusan.Jurusan = request.Form.Get("jurusan")

		var hobi entities.Hobi
		hobi.Hobi = request.Form.Get("hobi")

		var data = make(map[string]interface{})
		

		vErrors := validation.Struct(detail)
		if vErrors != nil {
			    data["detail"] = detail
				data["validation"] = vErrors
		} else {
			data ["pesan"] = "Data Berhasil Disimpan"
		detailModel.Create(detail)
		mahasiswaModel.Create(mahasiswa)
		jurusanModel.Create(jurusan)
		hobiModel.Create(hobi)
		}


	temp, _ := template.ParseFiles("views/mahasiswa/tambah.html")
	temp.Execute(response, data)
	}
	
	
}

func Edit(response http.ResponseWriter, request *http.Request) {
	if request.Method == http.MethodGet {
        
		queryString := request.URL.Query()
		id, _ := strconv.ParseInt(queryString.Get("id"), 10, 64)


		var detail entities.Detail
		var mahasiswa entities.Mahasiswa
		var jurusan entities.Jurusan
		var hobi entities.Hobi
		detailModel.Find(id, &detail)
		

		data := map[string]interface{}{
			"detail": detail,
			"mahasiswa": mahasiswa,
			"jurusan": jurusan,
			"hobi": hobi,
			
		}

		temp, err := template.ParseFiles("views/mahasiswa/edit.html")
	if err != nil {
		panic(err)
	}
	temp.Execute(response, data)

	} else if request.Method == http.MethodPost {

		request.ParseForm()

		var detail entities.Detail
		detail.Id, _= strconv.ParseInt(request.Form.Get("id"), 10, 64)
		detail.Nama = request.Form.Get("nama")
		detail.Usia = request.Form.Get("usia")
		detail.Gender = request.Form.Get("gender")
		detail.TanggalRegistrasi = request.Form.Get("tanggal_registrasi")
		detail.Jurusan = request.Form.Get("jurusan")
		detail.Hobi = request.Form.Get("hobi")

		var mahasiswa entities.Mahasiswa
		mahasiswa.Nama = request.Form.Get("nama")
		mahasiswa.Usia = request.Form.Get("usia")
		mahasiswa.Gender = request.Form.Get("gender")
		mahasiswa.TanggalRegistrasi = request.Form.Get("tanggal_registrasi")
		
		var jurusan entities.Jurusan
		jurusan.Jurusan = request.Form.Get("jurusan")

		var hobi entities.Hobi
		hobi.Hobi = request.Form.Get("hobi")

		var data = make(map[string]interface{})
		

		vErrors := validation.Struct(detail)
		if vErrors != nil {
			    data["detail"] = detail
				data["validation"] = vErrors
		} else {
			data ["pesan"] = "Data Berhasil Diperbaharui"
		detailModel.Update(detail)
		
		}

	temp, _ := template.ParseFiles("views/mahasiswa/edit.html")
	temp.Execute(response, data)
	}
	
}

func Hapus(response http.ResponseWriter, request *http.Request) {

	queryString := request.URL.Query()
	id, _ := strconv.ParseInt(queryString.Get("id"), 10, 64)

	detailModel.Delete(id)
	

	http.Redirect(response, request, "/mahasiswa", http.StatusSeeOther)
	
}