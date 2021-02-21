package iterator01

type User struct {
	name string
	age  int
}

func NewUser(name string, age int) *User {
	return &User{
		name: name,
		age: age,
	}
}

type ISet interface {
	CreateIterator() Iterator
	Add(*User)
}

type UserSet struct {
	users []*User
}

func NewUserSet() *UserSet {
	return &UserSet{}
}

func (u *UserSet) Add(user *User) {
	u.users = append(u.users, user)
}

func (u *UserSet) CreateIterator() Iterator {
	return &UserIterator{
		users: u.users,
	}
}