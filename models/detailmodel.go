package models

import (
	"database/sql"
	"fmt"

	"github.com/Rohuliman/API-GO/config"
	"github.com/Rohuliman/API-GO/entities"
)

type DetailModel struct {
	conn *sql.DB
}

func NewDetailModel() *DetailModel {
	conn, err := config.DBConnection()
	if err != nil {
		panic(err)
	}

	return &DetailModel{
		conn: conn,
	}

}

func (p *DetailModel) FindAll() ([]entities.Detail, error) {

	rows, err := p.conn.Query("select * from DataMahasiswa")
	if err != nil {
		return []entities.Detail{}, err
	}
	defer rows.Close()

	var dataDetail []entities.Detail
	for rows.Next() {
		var Detail entities.Detail
		rows.Scan(&Detail.Id,
			&Detail.Nama,
			&Detail.Usia,
			&Detail.Gender,
			&Detail.TanggalRegistrasi,
			&Detail.Jurusan,
			&Detail.Hobi,
			)

		if Detail.Gender == "0" {
			Detail.Gender = "Pria"
		}	else {
			Detail.Gender = "perempuan"
		}



		dataDetail = append(dataDetail, Detail)
	}

	return dataDetail, nil

}

func (p *DetailModel) Create(detail entities.Detail) bool {

	result, err := p.conn.Exec("insert into DataMahasiswa (nama, usia, gender, tanggal_registrasi, jurusan, hobi) values(?,?,?,?,?,?)", detail.Nama, detail.Usia, detail.Gender, detail.TanggalRegistrasi, detail.Jurusan, detail.Hobi)
                

	if err != nil {
		fmt.Println(err)
		return false
	}

	lastInsertId, _ := result.LastInsertId()

	return lastInsertId > 0
}

func (p *DetailModel) Find(id int64, detail *entities.Detail) error {

	return p.conn.QueryRow("select * From DataMahasiswa where id = ?", id).Scan(&detail.Id,
		&detail.Nama,
		&detail.Usia,
		&detail.Gender,
		&detail.TanggalRegistrasi,
		&detail.Jurusan,
		&detail.Hobi)
}

func (p *DetailModel) Update(detail entities.Detail) error{

	_, err := p.conn.Exec(
		"update DataMahasiswa set nama = ?, usia = ?, gender = ?, tanggal_registrasi = ?, jurusan = ?, hobi = ? where id = ?",
	    detail.Nama, detail.Usia, detail.Gender, detail.TanggalRegistrasi, detail.Jurusan, detail.Hobi, detail.Id)

	if err != nil {
		return err
	}
	return nil
}

func (p *DetailModel) Delete(id int64) {
	p.conn.Exec("delete from DataMahasiswa where id = ?", id)
}