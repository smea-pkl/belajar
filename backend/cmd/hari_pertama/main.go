package main

import (
	"errors"
	"fmt"
	"os"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// harus disimpan ndek 2 cara.
// 1. database
// 2. file

func main() {

	dbfilename := "pkl.db"
	db, err := gorm.Open(sqlite.Open(dbfilename), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	err = db.AutoMigrate(&User{})
	if err != nil {
		panic(err)
	}

	// test

	// var blacklist []string = []string{
	// 	"sutarman",
	// 	"andika",
	// 	"fitri", // P --> anda ditolak karena janda.
	// }

	// var nama string
	ask := "y"
	filename := "user.txt"

	file, err := os.OpenFile(filename, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0644)
	if err != nil {
		panic(err)
	}

	defer file.Close()

	// var users []*User = []*User{}

	for ask == "y" {

		// input
		usr, err := askUser()
		if err != nil {
			panic(err)
		}

		// filter
		check, err := filterUser(usr)
		if err != nil {
			panic(err)
		}

		if !check {
			err := errors.New("user kena filter")
			panic(err)
		}

		// hasil / output
		fmt.Println(usr) // print cli

		// save file
		line := usr.ToString()
		file.WriteString(line + "\n")

		err = db.Save(&usr).Error
		if err != nil {
			panic(err)
		}

		fmt.Println("User Berhasil Disimpan")

		fmt.Println("Lagi ?")
		_, err = fmt.Scan(&ask)
		if err != nil {
			panic(err)
		}
	}

}
