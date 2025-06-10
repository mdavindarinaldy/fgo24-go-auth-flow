package main

import (
	"fgo24-go-auth-flow/auth"
	"fgo24-go-auth-flow/utils"
	"fmt"
	"os"
)

func main() {
	// users := []utils.User{}
	for {
		utils.Menu()
		option, err := utils.GetInputInt("Pilih menu [1-4]: ")
		if err != nil || option > 4 {
			fmt.Println("Input tidak valid, mohon masukkan angka 1-4!")
		} else if option == 4 {
			os.Exit(0)
		} else if option == 1 {
			auth.Login(&utils.Users)
		} else if option == 2 {
			auth.Register(&utils.Users)
			fmt.Println(utils.Users)
		} else if option == 3 {
			auth.ForgotPassword(&utils.Users)
		}
	}
}
