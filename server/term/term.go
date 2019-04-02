package term

import (
	"fmt"
	"github.com/sparkoo/payterm/model"
	"github.com/sparkoo/payterm/peripherals"
	"github.com/sparkoo/payterm/websocket"
)

type Term struct {
	loggedUser *model.Account
	users      map[model.UserId]*model.Account

	server websocket.Server

	io termIO
}

func NewTerm(server websocket.Server, accounts map[model.UserId]*model.Account,
	k *peripherals.Keyboard, d *peripherals.Display, b *peripherals.Buzzer, cr *peripherals.CardReader) *Term {
	return &Term{
		server:     server,
		loggedUser: nil,
		users:      accounts,
		io: termIO{
			keyboard:   k,
			display:    d,
			buzzer:     b,
			cardReader: cr,
		},
	}
}

func (t *Term) Authenticate(userid model.UserId) (*model.Account, error) {
	if account, ok := t.users[userid]; ok {
		t.loggedUser = account
		return account, nil
	} else {
		return nil, &UserNotFoundError{username: string(userid)}
	}
}

func (t *Term) GetLoggedAccount() *model.Account {
	return t.loggedUser
}

func (t *Term) Logout() {
	t.loggedUser = nil
}

func (t *Term) Start() {
	go t.server.Start()
	for {
		key := (*t.io.keyboard).Read()
		(*t.io.display).Write(key)
	}
}

type UserNotFoundError struct {
	username string
}

func (e *UserNotFoundError) Error() string {
	return fmt.Sprintf("user [%v] not found", e.username)
}
