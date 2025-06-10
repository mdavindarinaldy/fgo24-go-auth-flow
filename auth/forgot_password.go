package auth

import (
	"fgo24-go-auth-flow/utils"
	"fmt"
)

func ForgotPassword(users *[]utils.User) {
	for {
		if len(*users) < 1 {
			fmt.Println("\nBelum ada akun yang terdaftar!")
			return
		}
		fmt.Println("----Menu Forgot Password----")
		name, err := utils.GetInputString("Silakan masukan nama Anda: ")
		if err != nil || name == "" {
			fmt.Println("Nama tidak valid! Silakan coba lagi")
		} else if checkUser(name, *users) {
			fmt.Println("\n[PROSES VERIFIKASI USER MISAL DENGAN EMAIL]")
			newPass, err := utils.GetInputString("Silakan masukkan password baru: ")
			if err != nil || newPass == "" {
				fmt.Println("Password baru tidak valid! Silakan coba lagi")
			} else {
				for i := range *users {
					// if users[i] == name {}
				}
			}
		}
	}
}
