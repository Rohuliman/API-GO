package models

import (
	"database/sql"
	"fmt"

	"github.com/Rohuliman/API-GO/config"
	"github.com/Rohuliman/API-GO/entities"
)

type HobiModel struct {
	conn *sql.DB
}

func NewHobiModel() *HobiModel {
	conn, err := config.DBConnection()
	if err != nil {
		panic(err)
	}

	return &HobiModel{
		conn: conn,
	}

}
func (p *HobiModel) Find() ([]entities.Hobi, error) {

	rows, err := p.conn.Query("select * from hobi")
	if err != nil {
		return []entities.Hobi{}, err
	}
	defer rows.Close()

	var dataHobi []entities.Hobi
	for rows.Next() {
		var Hobi entities.Hobi
		rows.Scan(&Hobi.Id,
			&Hobi.Hobi,
			)


		dataHobi = append(dataHobi, Hobi)
	}

	return dataHobi, nil

}

func (p *HobiModel) Create(hobi entities.Hobi) bool {

	result, err := p.conn.Exec("insert into hobi (nama_hobi) values(?)", hobi.Hobi)
                

	if err != nil {
		fmt.Println(err)
		return false
	}

	lastInsertId, _ := result.LastInsertId()

	return lastInsertId > 0
}

func (p *DetailModel) Findd(id int64, detail *entities.Detail) error {

	return p.conn.QueryRow("select * From hobi where id = ?", id).Scan(&detail.Id,
		&detail.Hobi)
}

func (p *HobiModel) Updite(hobi entities.Hobi) error{

	_, err := p.conn.Exec(
		"update Hobi set nama_hobi = ?, whre id = ?",
	     hobi.Hobi, hobi.Id)

	if err != nil {
		return err
	}
	return nil
}

func (p *HobiModel) Dalete(id int64){
	p.conn.Exec("delete from DataMahasiswa where id = ?", id)
}