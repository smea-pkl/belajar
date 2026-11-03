package main

import (
	"errors"
	"fmt"
	"os"
)

func main() {

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

		fmt.Println("Lagi ?")
		_, err = fmt.Scan(&ask)
		if err != nil {
			panic(err)
		}
	}

}
