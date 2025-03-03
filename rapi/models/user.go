package models

import (
		"example.com/rapi/db"
		"example.com/rapi/utils"
	)


type User struct{

	ID int64 `json:"id"`	
    Email string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
	
}
func (u *User) Save() error {
	query := `INSERT INTO users (email, password) VALUES (?, ?)`
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()
    s,err:=utils.HashPassword(u.Password)
	if err != nil {
			return err}

	res, err := stmt.Exec(u.Email,s)
	if err != nil {
		return err
	}

	id, err := res.LastInsertId()
	if err != nil {
		return err
	}
	u.ID = id
	return nil
}
func (u *User) ValidateCredentials()  error {
	query := `SELECT id,password FROM users WHERE email = ?`
	var password string
	err := db.DB.QueryRow(query, u.Email).Scan(&u.ID,&password)
	if err != nil {
		return err
	}
	paasswordIsValid:= utils.CheckPasswordHash(u.Password, password)
	if paasswordIsValid != nil {
		return nil}
	return nil

}