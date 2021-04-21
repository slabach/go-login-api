package models

// create user model
type User struct {
	Username string `json:"username"`
	Password []byte `json:"password"`
}

// create user list model
type UserList struct {
	Users []User
}

// function for initializing new user list
func NewUserList() *UserList {
	return &UserList{
		Users: []User{},
	}
}

// function for adding user to user list
func (u *UserList) Add(user User) {
	u.Users = append(u.Users, user)
}

// function to return all current users
func (r *UserList) GetUsers() []User {
	return r.Users
}
