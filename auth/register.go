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

func Register(users *[]User) {
	for {
		fmt.Println("\n-----Menu Registrasi-----")
		name, errName := utils.GetInputString("Silakan masukkan nama anda: ")
		password, errPass := utils.GetInputString("Silakan masukkan password anda: ")
		if errName != nil || errPass != nil || name == "" || password == "" {
			fmt.Println("Data tidak valid! Silakan coba lagi")
		} else {
			password = encode(password)
			u := User{
				Name:     name,
				Password: password,
			}
			// u.addUser(users) // method
			showData(u, users) // interface
			fmt.Println("Pendaftaran akun berhasil")
			fmt.Println("-------------------------")
			return
		}
	}
}
