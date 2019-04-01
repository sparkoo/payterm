package term

import "fmt"

type userid string

type Term struct {
	loggedUser Account
	users      map[userid]Account
}

func (t *Term) Authenticate(userid userid) (Account, error) {
	if account, ok := t.users[userid]; ok {
		t.loggedUser = account
		return account, nil
	} else {
		return nil, &UserNotFoundError{username: string(userid)}
	}
}

func (t *Term) GetLoggetAccount() Account {
	return t.loggedUser
}

func (t *Term) Logout(Account) {
	t.loggedUser = nil
}

type UserNotFoundError struct {
	username string
}

func (e *UserNotFoundError) Error() string {
	return fmt.Sprintf("user [%v] not found", e.username)
}
