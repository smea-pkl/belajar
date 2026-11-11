package main

type User struct {
	Name     string
	Username string
}

func main() {

	// users := []User{
	// 	User{Name: "anita", Username: "sdasd"},
	// 	User{Name: "dita", Username: "asd"},
	// 	User{Name: "anita", Username: "3"},
	// 	User{Name: "anita", Username: "1"},
	// 	User{Name: "dita", Username: "4"},
	// 	User{Name: "j", Username: "fgh"},
	// }

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
