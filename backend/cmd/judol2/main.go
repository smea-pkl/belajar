package main

import (
	"fmt"
	"math/rand/v2"
)

type User struct {
	ID       int
	Saldo    uint
	Username string
	Password string
}

func (j *User) IncSaldo(amount uint) error {
	j.Saldo = j.Saldo + amount
	return nil
}

func (j *User) DecSaldo(amount uint) error {
	j.Saldo = j.Saldo - amount
	return nil
}

func (j *User) Main() error {
	return nil
}

func (j *User) Login() error {
	return nil
}

func (j *User) Topup() error {
	return nil
}

func (j *User) Tarik() error {
	return nil
}

func (j *User) Tebak() error {
	return nil
}

type Game struct {
	Posisi string // 'angka' or 'gambar'
}

func (j *Game) Acak() error {

	i := rand.IntN(2)
	koin := []string{"angka", "gambar"}
	j.Posisi = koin[i]

	return nil
}

type JudolApp struct {
	// Property
	Users       []*User
	CurrentUser *User
	Game        Game
}

func (j *JudolApp) Daftar() error {

	newUser := User{
		ID:       1,
		Saldo:    10_000_000,
		Username: "test",
		Password: "test",
	}

	j.Users = append(j.Users, &newUser)
	j.CurrentUser = &newUser

	fmt.Printf("Anda terdaftar dengan username %s dengan password %s \n", "test", "test")

	return nil
}

func (j *JudolApp) Login() error {
	panic("belum ditulis")
}

func (j *JudolApp) Run() error {

	fmt.Println(`
		Mau Apa ?
		1. Daftar
		2. Login
	`)

	var action string
	_, err := fmt.Scan(&action)
	if err != nil {
		return err
	}

	actionMap := map[string]func() error{
		"1": j.Daftar,
		"2": j.Login,
	}

	err = actionMap[action]()
	if err != nil {
		return nil
	}

	// slog.Info("posisi koin", "posisi", j.Game.Posisi)

	for {

		err = j.Game.Acak() // diacak
		if err != nil {
			return err
		}

		// tanya tebakan
		var tebakan string

		fmt.Println("tebak ? \n1. 'angka'\n2. 'gambar'")
		_, err = fmt.Scan(&tebakan)

		if err != nil {
			return err
		}

		// cocokkan tebakan
		if j.Game.Posisi == tebakan {
			j.CurrentUser.IncSaldo(500_000)

			fmt.Printf("selamat kamu benar, saldo anda %d\n", j.CurrentUser.Saldo)

		} else {
			j.CurrentUser.DecSaldo(700_000)

			fmt.Printf("kamu salah, saldo anda %d\n", j.CurrentUser.Saldo)
		}
	}

	// return nil
}

func main() {

	var app JudolApp = JudolApp{
		Game: Game{},
	} // buat onject
	err := app.Run() // jalan aplikasi

	if err != nil {
		panic(err)
	}

}
