package models

import (
	"database/sql"
	"fmt"

	"github.com/Rohuliman/API-GO/config"
	"github.com/Rohuliman/API-GO/entities"
)

type IdModel struct {
	conn *sql.DB
}

func NewIdModel() *IdModel{
	conn, err := config.DBConnection()
	if err != nil {
		panic(err)
	}

	return &IdModel{
		conn: conn,
	}

}

func (p *IdModel) Create(Id entities.Id) bool {

	result, err := p.conn.Exec("insert into mahasiswa_hobi (id_mahasiswa, id_hobi) values(?,?)", Id.Id_Mahasiswa, Id.Id_Hobi)
                

	if err != nil {
		fmt.Println(err)
		return false
	}

	lastInsertId, _ := result.LastInsertId()

	return lastInsertId > 0
}