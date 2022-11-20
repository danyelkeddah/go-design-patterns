package iterator

type UserCollection struct {
	Users []*User
}

func (c *UserCollection) CreateIterator() Iterator {
	return &UserIterator{Users: c.Users}
}
