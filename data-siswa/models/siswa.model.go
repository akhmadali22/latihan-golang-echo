package models

import (
	"data-siswa/db"
	"log"
	"net/http"

	validator "github.com/go-playground/validator/v10"
)

type Siswa struct {
	ID      int    `json:"id"`
	Nama    string `json:"nama" validate:"required"`
	Kelas   string `json:"kelas" validate:"required"`
	Jenjang string `json:"jenjang" validate:"required"`
	Smt     string `json:"smt" validate:"required"`
}

func FetchAllSiswa() (Response, error) {
	var obj Siswa
	var arrobj []Siswa
	var res Response

	con := db.CreateCon()

	sqlStatement := "SELECT * FROM siswa"

	rows, err := con.Query(sqlStatement)
	defer rows.Close()

	if err != nil {
		return res, err
	}

	for rows.Next() {
		err := rows.Scan(&obj.ID, &obj.Nama, &obj.Kelas, &obj.Jenjang, &obj.Smt)

		if err != nil {
			return res, err
		}

		arrobj = append(arrobj, obj)
	}
	res.Status = 200
	res.Message = "Success"
	res.Data = arrobj

	return res, nil
}

func StoreSiswa(nama, kelas, jenjang, smt string) (Response, error) {
	var obj Siswa
	var id int
	var res Response

	v := validator.New()

	siswa := Siswa{
		Nama:    nama,
		Kelas:   kelas,
		Jenjang: jenjang,
		Smt:     smt,
	}
	err := v.Struct(siswa)
	if err != nil {
		return res, err
	}

	con := db.CreateCon()

	sqlStatement := "INSERT INTO siswa (nama, kelas, jenjang, smt) VALUES ($1, $2, $3, $4)"
	stmt, err := con.Prepare(sqlStatement)
	if err != nil {
		return res, err
	}

	if _, err := stmt.Exec(nama, kelas, jenjang, smt); err != nil {
		log.Fatal(err)
	}

	getMaxId := "SELECT max(id) FROM siswa;"
	rows, err := con.Query(getMaxId)
	if err != nil {
		return res, err
	}
	// defer rows.Close()
	for rows.Next() {
		err := rows.Scan(&obj.ID)
		if err != nil {
			return res, err
		}
		id = obj.ID
	}

	res.Status = 200
	res.Message = "Success"
	res.Data = map[string]int64{
		"last_inserted_id": int64(id),
	}

	return res, nil
}

func UpdateSiswa(id int, nama, kelas, jenjang, smt string) (Response, error) {
	var res Response

	con := db.CreateCon()

	sqlStatement := "UPDATE siswa SET nama=$1, kelas=$2, jenjang=$3, smt=$4 WHERE id=$5"
	stmt, err := con.Prepare(sqlStatement)
	if err != nil {
		return res, err
	}

	if _, err := stmt.Exec(nama, kelas, jenjang, smt, id); err != nil {
		log.Fatal(err)
	}

	res.Status = http.StatusOK
	res.Message = "Success"
	res.Data = map[string]int64{
		"last_updated_id": int64(id),
	}

	return res, nil
}

func DeleteSiswa(id int) (Response, error) {
	var res Response

	con := db.CreateCon()

	sqlStatement := "DELETE FROM siswa WHERE id=$1"
	stmt, err := con.Prepare(sqlStatement)
	if err != nil {
		return res, err
	}

	if _, err := stmt.Exec(id); err != nil {
		log.Fatal(err)
	}

	res.Status = http.StatusOK
	res.Message = "Success"
	res.Data = map[string]int64{
		"deleted_id": int64(id),
	}

	return res, nil
}
