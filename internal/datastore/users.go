package datastore

type Users struct {
	userSet map[string]bool // map of strings that point to a bool, bool is only used to see if a user is already in our list
	Users   []string        `json:"users"`
}

func (u *Users) CheckAndAddNewUser(user string) error {
	_, pres := u.userSet[user]
	if pres == false {
		u.userSet[user] = true
		u.Users = append(u.Users, user)
	}
	return nil
}

func ConstructUsers() *Users {
	var u Users
	u.userSet = make(map[string]bool)
	return &u
}
