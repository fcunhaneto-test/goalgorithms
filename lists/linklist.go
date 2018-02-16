package lists

// User struct of users only as an example
type User struct {
	name string
	id   string
}

// LinkUser link list of users
type LinkUser struct {
	size int
	next *User
}

// func InsertLL(u User, )
