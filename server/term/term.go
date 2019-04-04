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

	running bool
	server  websocket.Server

	io termIO
}

func NewTerm(server websocket.Server, accounts map[model.UserId]*model.Account,
	k *peripherals.Keyboard, d *peripherals.Display, b *peripherals.Buzzer, cr *peripherals.CardReader) *Term {
	return &Term{
		running:    false,
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
	t.running = true
	for t.running {
		key := (*t.io.cardReader).Read()
		if user, found := t.users[model.UserId(key)]; found {
			(*t.io.display).Write(fmt.Sprintf("user [%s] has %d$", user.Name(), user.Balance()))
		} else {
			fmt.Printf("user with key [%s] not found\n", key)
		}
	}
}

func (t *Term) Close() {
	t.running = false
	t.Close()
}

type UserNotFoundError struct {
	username string
}

func (e *UserNotFoundError) Error() string {
	return fmt.Sprintf("user [%v] not found", e.username)
}
