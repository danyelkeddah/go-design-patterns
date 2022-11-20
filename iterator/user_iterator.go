package iterator

type UserIterator struct {
	Index int
	Users []*User
}

func (i *UserIterator) HasNext() bool {
	if i.Index < len(i.Users) {
		return true
	}
	return false
}

func (i *UserIterator) GetNext() *User {
	if i.HasNext() {
		user := i.Users[i.Index]
		i.Index++
		return user
	}

	return nil
}
