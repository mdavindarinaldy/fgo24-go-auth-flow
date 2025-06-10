package auth

import (
	"fgo24-go-auth-flow/utils"
	"fmt"
)

func checkPass(pass string, name string, users []utils.User) bool {
	pass = encode(pass)
	for i := range users {
		if users[i].Name == name && users[i].Password == pass {
			return true
		}
	}
	return false
}

func loginSuccess() {
	for {
		fmt.Println("Berhasil login!")
		option, err := utils.GetInputInt("Silakan masukkan angka 1 untuk logout dan kembali ke menu utama: ")
		if err != nil || option != 1 {
			fmt.Println("Input tidak valid silakan coba lagi")
		} else if option == 1 {
			fmt.Println("Logout berhasil dilakukan")
			fmt.Println("-------------------------")
			return
		}
	}
}

func Login(users *[]utils.User) {
	for {
		if len(*users) < 1 {
			fmt.Println("\nBelum ada akun yang terdaftar!")
			return
		}
		fmt.Println("\n-----Menu Login-----")
		name, errName := utils.GetInputString("Silakan masukkan nama anda: ")
		password, errPass := utils.GetInputString("Silakan masukkan password anda: ")
		if errName != nil || errPass != nil || name == "" || password == "" {
			fmt.Println("Data input tidak valid! Silakan coba lagi")
		} else if checkPass(password, name, *users) {
			loginSuccess()
			return
		} else {
			fmt.Println("Nama dan/atau Password salah! Silakan coba lagi")
		}
	}
}
