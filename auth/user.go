package auth

type dataUser interface {
	addUser(users *[]User)
}

type User struct {
	Name     string
	Password string
}

var Users []User

func showData(du dataUser, users *[]User) {
	du.addUser(users)
}

func (u User) addUser(Users *[]User) {
	*Users = append(*Users, u)
}
