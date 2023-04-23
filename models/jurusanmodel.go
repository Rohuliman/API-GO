package models

import (
	"database/sql"
	"fmt"

	"github.com/Rohuliman/API-GO/config"
	"github.com/Rohuliman/API-GO/entities"
)

type JurusanModel struct {
	conn *sql.DB
}

func NewJurusanModel() *JurusanModel {
	conn, err := config.DBConnection()
	if err != nil {
		panic(err)
	}

	return &JurusanModel{
		conn: conn,
	}

}
func (p *JurusanModel) Find() ([]entities.Jurusan, error) {

	rows, err := p.conn.Query("select * from jurusan")
	if err != nil {
		return []entities.Jurusan{}, err
	}
	defer rows.Close()

	var dataJurusan []entities.Jurusan
	for rows.Next() {
		var Jurusan entities.Jurusan
		rows.Scan(&Jurusan.Id,
			&Jurusan.Jurusan,
			)

		dataJurusan = append(dataJurusan, Jurusan)
	}

	return dataJurusan, nil

}

func (p *JurusanModel) Create(jurusan entities.Jurusan) bool {

	result, err := p.conn.Exec("insert into jurusan (nama_jurusan) values(?)", jurusan.Jurusan)
                

	if err != nil {
		fmt.Println(err)
		return false
	}

	lastInsertId, _ := result.LastInsertId()

	return lastInsertId > 0
}

func (p *DetailModel) Find3(id int64, detail *entities.Detail) error {

	return p.conn.QueryRow("select * From jurusan where id = ?", id).Scan(&detail.Id,
		&detail.Jurusan,)
}

func (p *JurusanModel) Upduate(jurusan entities.Jurusan) error{

	_, err := p.conn.Exec(
		"update jurusan set nama_jurusan = ? where id = ?",
	    jurusan.Jurusan, jurusan.Id)

	if err != nil {
		return err
	}
	return nil
}
func (p *JurusanModel) Dalate(id int64){
	p.conn.Exec("delete from jurusan where id = ?", id)
}