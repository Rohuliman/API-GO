package models

import (
	"database/sql"
	"fmt"

	"github.com/Rohuliman/API-GO/config"
	"github.com/Rohuliman/API-GO/entities"
)

type MahasiswaModel struct {
	conn *sql.DB
}

func NewMahasiswaModel() *MahasiswaModel {
	conn, err := config.DBConnection()
	if err != nil {
		panic(err)
	}

	return &MahasiswaModel{
		conn: conn,
	}

}

func (p *MahasiswaModel) Find() ([]entities.Mahasiswa, error) {

	rows, err := p.conn.Query("select * from mahasiswa")
	if err != nil {
		return []entities.Mahasiswa{}, err
	}
	defer rows.Close()

	var dataMahasiswa []entities.Mahasiswa
	for rows.Next() {
		var Mahasiswa entities.Mahasiswa
		rows.Scan(&Mahasiswa.Id,
			&Mahasiswa.Nama,
			&Mahasiswa.Usia,
			&Mahasiswa.Gender,
			&Mahasiswa.TanggalRegistrasi,
			)

		if Mahasiswa.Gender == "0" {
			Mahasiswa.Gender = "Pria"
		}	else {
			Mahasiswa.Gender = "perempuan"
		}



		dataMahasiswa = append(dataMahasiswa, Mahasiswa)
	}

	return dataMahasiswa, nil

}

func (p *MahasiswaModel) Create(mahasiswa entities.Mahasiswa) bool {

	result, err := p.conn.Exec("insert into mahasiswa (nama, usia, gender, tanggal_registrasi) values(?,?,?,?)", mahasiswa.Nama, mahasiswa.Usia, mahasiswa.Gender, mahasiswa.TanggalRegistrasi)
                

	if err != nil {
		fmt.Println(err)
		return false
	}

	lastInsertId, _ := result.LastInsertId()

	return lastInsertId > 0
}

func (p *DetailModel) Find4(id int64, detail *entities.Detail) error {

	return p.conn.QueryRow("select * From Mahasiswa where id = ?", id).Scan(&detail.Id,
		&detail.Nama,
		&detail.Usia,
		&detail.Gender,
		&detail.TanggalRegistrasi,)
}

func (p *MahasiswaModel) Apdate(mahasiswa entities.Mahasiswa) error{

	_, err := p.conn.Exec(
		"update Mahasiswa set nama = ?, usia = ?, gender = ?, tanggal_registrasi = ? where id = ?",
	    mahasiswa.Nama, mahasiswa.Usia, mahasiswa.Gender, mahasiswa.TanggalRegistrasi, mahasiswa.Id)

	if err != nil {
		return err
	}
	return nil
}

func (p *MahasiswaModel) Delate(id int64){
	p.conn.Exec("delete from Mahasiswa where id = ?", id)
}