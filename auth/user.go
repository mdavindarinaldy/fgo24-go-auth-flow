package auth

type dataUser interface {
	addUser()
	removeUser()
}

type User struct {
	Name     string
	Password string
}

var Users []User

func (u User) addUser(Users *[]User) {
	*Users = append(*Users, u)
}

func (u User) removeUser(Users *[]User) {
	// kode hapus user
}
