package moviestore

type UserID uint16



type User struct {
	Name string
	Age Age
	UserID UserID
}

func (u *User) String() string {
return "user"
}
