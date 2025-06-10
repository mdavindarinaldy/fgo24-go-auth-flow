package auth

import (
	"crypto/md5"
	"fgo24-go-auth-flow/utils"
	"fmt"
)

func encode(pass string) string {
	hash := md5.New()
	_, _ = hash.Write([]byte(pass))
	md5 := hash.Sum(nil)
	return fmt.Sprintf("%x", md5)
}

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
				password = encode(password)
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
