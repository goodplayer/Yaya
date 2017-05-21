package repo

type User struct {
	Id string
}

func NewNonUser() *User {
	u := new(User)
	u.Id = "00000000000000000000000000000000"
	return u
}
