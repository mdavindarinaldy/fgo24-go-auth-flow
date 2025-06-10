package auth

import "fgo24-go-auth-flow/utils"

func checkUser(name string, users []utils.User) bool {
	for i := range users {
		if users[i].Name == name {
			return true
		}
	}
	return false
}
