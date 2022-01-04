package main

type User struct {
	ID        int
	FirstName string
	LastName  string
}

var (
	users  []*User
	nextId = 1
)

func GetUser() []*User {
	return users
}

fun AddUser(u User) (User, error){
	u.ID = nextId
	nextId++
	users =append(users, &u)
	return u, nil
}