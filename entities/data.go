package entities

type Detail struct {
	Id int64
	Nama string  `validate:"required" `
	Usia string   `validate:"required"`
	Gender string  `validate:"required"`
	TanggalRegistrasi string `validate:"required" label:"Tanggal Registrasi"`
	Jurusan string `validate:"required"`
	Hobi string `validate:"required"`
}

type Mahasiswa struct {
	Id int64
	Nama string  `validate:"required"`
	Usia string   `validate:"required"`
	Gender string  `validate:"required"`
	TanggalRegistrasi string `validate:"required"`
}

type Jurusan struct {
	Id int64
	Jurusan string `validate:"required"`
}

type Hobi struct {
	Id int64
	Hobi string `validate:"required"`
}

type Id struct {
	Id_Mahasiswa int
	Id_Hobi int
}