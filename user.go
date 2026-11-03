package main

import "fmt"

type User struct {
	Name     string
	Username string
	Gender   string
	Email    string
}

func (usr User) ToString() string {

	var line string
	line = fmt.Sprintf("%s|%s|%s", usr.Username, usr.Name, usr.Gender)
	return line
}

func askUser() (User, error) {
	var usr User
	fmt.Print("Masukkan Nama Anda :")
	_, err := fmt.Scan(&usr.Name)

	if err != nil {
		return usr, err
	}

	fmt.Print("Masukkan Username Anda :")
	_, err = fmt.Scan(&usr.Username)

	if err != nil {
		return usr, err
	}

	fmt.Print("Masukkan Gender Anda :")
	_, err = fmt.Scan(&usr.Gender)

	if err != nil {
		return usr, err
	}

	fmt.Print("Masukkan Email Anda :")
	_, err = fmt.Scan(&usr.Email)

	if err != nil {
		return usr, err
	}

	return usr, err
}
