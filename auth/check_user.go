package auth

func checkUser(name string, users []User) bool {
	for i := range users {
		if users[i].Name == name {
			return true
		}
	}
	return false
}

func checkPass(pass string, name string, users []User) bool {
	pass = encode(pass)
	for i := range users {
		if users[i].Name == name && users[i].Password == pass {
			return true
		}
	}
	return false
}
