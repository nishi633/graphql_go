package model

type User struct {
	ID    int
	Name  string
	Email string
}

func UserAll() []*User {
	return []*User{
		NewUser(1, "山田一郎", "yamada1@test.com"),
		NewUser(2, "山田二郎", "yamada2@test.com"),
		NewUser(3, "山田三郎", "yamada3@test.com"),
	}
}

func UserDetail() *User {
	return NewUser(3, "山田三郎", "yamada3@test.com")
}

func NewUser(id int, name, email string) *User {
	return &User{
		ID:    id,
		Name:  name,
		Email: email,
	}
}
