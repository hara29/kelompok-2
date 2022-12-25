package models

import (
	"Kelompok2/config"
	"Kelompok2/entities"
	"database/sql"
	"fmt"
)

type MhsModel struct{
	db *sql.DB
}

func NewMhsModel() *MhsModel{
	conn, err:= config.DBConn()

	if err!= nil{
		panic(err)
	}

	return &MhsModel{
		db: conn,
	}
}

func (m *MhsModel) FindAll(mhs *entities.Mahasiswa, fieldName, fieldValue string) ([]entities.Mahasiswa, error){
	// rows, err:= m.db.Query("select * from mahasiswa")
	rows, err:= m.db.Query("select * from mahasiswa where " + fieldName + " = ? ", fieldValue)
	
	if err != nil {
		return []entities.Mahasiswa{}, err
	}
	defer rows.Close()

	var dataMhs []entities.Mahasiswa
	for rows.Next(){
		var mhs entities.Mahasiswa
		rows.Scan(&mhs.Pert, &mhs.No, &mhs.Nama, &mhs.NPM, &mhs.Kelas, &mhs.UTS, &mhs.UAS, &mhs.Total)
		dataMhs = append(dataMhs, mhs)
	}
	return dataMhs, nil
}

func (m *MhsModel) Create(mhs entities.Mahasiswa, uts, uas, total float64, pert string, no int)bool{
	result, err := m.db.Exec("UPDATE mahasiswa SET uts=?, uas=?, total=? WHERE pert=? AND no=?", uts, uas, total, pert, no)
	if err != nil {
		fmt.Println(err)
		return false
	}
	lastInsertId, _ :=result.LastInsertId()
	return lastInsertId > 0
}

// func (m *MhsModel) Find(no int64, pert string, mhs *entities.Mahasiswa) error {

// 	return m.db.QueryRow("select * from mahasiswa where no = ? AND pert='?'", no, pert).Scan(
// 		&mhs.No, &mhs.Pert, &mhs.NPM, &mhs.Nama, &mhs.Kelas, &mhs.UTS, &mhs.UAS, &mhs.Total)
// }