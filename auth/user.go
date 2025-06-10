package auth

type User struct {
	Name     string
	Password string
}

var Users []User

func (u User) addUser(Users *[]User) {
	*Users = append(*Users, u)
}
