package models

import (
	"Kelompok2/config"
	"Kelompok2/entities"
	"database/sql"
)

type UserModel struct{
	db *sql.DB
}


func NewUserModel() *UserModel{
	conn, err:= config.DBConn()

	if err!= nil{
		panic(err)
	}

	return &UserModel{
		db: conn,
	}
}

func (u UserModel) Where(user *entities.User, fieldName, fieldValue string) error{
	//membuat query
	row, err := u.db.Query("select id_user, nama, password from user where " + fieldName + " = ? limit 1", fieldValue)

	if err != nil{
		return err
	}

	defer row.Close()//menutup database

	for row.Next(){
		row.Scan(&user.Username, &user.NamaLengkap, &user.Password)
	}

	return nil
}