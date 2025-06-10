package auth

import (
	"fgo24-go-auth-flow/utils"
	"fmt"
)

func Register(users *[]utils.User) {
	for {
		fmt.Println("\n-----Menu Registrasi-----")
		name, errName := utils.GetInputString("Silakan masukkan nama anda: ")
		if errName != nil || name == "" {
			fmt.Println("Nama tidak valid! Silakan coba lagi")
		} else {
			password, errPass := utils.GetInputString("Silakan masukkan password anda: ")
			if errPass != nil || password == "" {
				fmt.Println("Password tidak valid! Silakan coba lagi")
			} else {
				user := utils.User{
					Name:     name,
					Password: password,
				}
				*users = append(*users, user)
				fmt.Println("Pendaftaran akun berhasil")
				fmt.Println("-------------------------")
				return
			}
		}
	}
}
