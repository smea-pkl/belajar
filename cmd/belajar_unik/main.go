package main

import "fmt"

func main() {

	users := []string{
		"anita",
		"fitri",
		"asdasd",
		"none",
		"fitri",
		"asdasd",
		"james",
		"fitri",
		"asdasd",
		"james",
		"anita",
		"fitri",
		"asdasd",
		"none",
		"fitri",
		"asdasd",
		"james",
		"fitri",
		"asdasd",
		"james",
		"fitri",
		"asdasd",
		"james",
		"anita",
		"fitri",
		"asdasd",
		"none",
		"fitri",
		"asdasd",
	}

	user_unik := []string{}

	unik := map[string]bool{}

	for _, item := range users {
		unik[item] = true
	}

	for item := range unik {
		user_unik = append(user_unik, item)
	}

	fmt.Println(user_unik)

}

// contoh ruwet
// func main() {

// 	users := []string{
// 		"anita",
// 		"fitri",
// 		"asdasd",
// 		"none",
// 		"fitri",
// 		"asdasd",
// 		"james",
// 		"fitri",
// 		"asdasd",
// 		"james",
// 		"anita",
// 		"fitri",
// 		"asdasd",
// 		"none",
// 		"fitri",
// 		"asdasd",
// 		"james",
// 		"fitri",
// 		"asdasd",
// 		"james",
// 		"fitri",
// 		"asdasd",
// 		"james",
// 		"anita",
// 		"fitri",
// 		"asdasd",
// 		"none",
// 		"fitri",
// 		"asdasd",
// 	}

// 	user_unik := []string{}

// 	// 1. ambil satu persatu users.
// 	for _, item := range users {

// 		var duplikat bool = false
// 		for _, unik := range user_unik {
// 			// 2. check per item di user unik
// 			if item == unik {
// 				duplikat = true
// 			}
// 		}

// 		// 3. kalau ada dibuang
// 		// 4. kalau tidak ada, dimasukkan.
// 		if duplikat {
// 			continue
// 		} else {
// 			user_unik = append(user_unik, item)
// 		}

// 	}

// 	// tampilkan hasilnya
// 	fmt.Println(user_unik)

// }
