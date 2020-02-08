package term

import (
	"fmt"
	"github.com/sparkoo/payterm/model"
	"github.com/sparkoo/payterm/peripherals"
	server2 "github.com/sparkoo/payterm/server"
)

type Term struct {
	loggedUser *model.Account
	users      map[model.UserId]*model.Account

	running bool
	server  server2.Server

	io  termIO
	pay *payment
}

func NewTerm(server server2.Server, accounts map[model.UserId]*model.Account,
	k peripherals.InputReader, d peripherals.OutputWriter, b peripherals.OutputWriter, cr peripherals.InputReader) *Term {
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
	t.io.display.Write("Hello")
	t.mainLoop()
}

func (t *Term) mainLoop() {
	crChannel := make(chan string)
	keyboardChannel := make(chan string)
	for t.running {
		go t.read(crChannel, t.io.cardReader)
		go t.read(keyboardChannel, t.io.keyboard)

		select {
		case card := <-crChannel:
			t.cardRead(card)
		case key := <-keyboardChannel:
			t.keyPressed(key)
		}
	}
}

func (t *Term) cardRead(card string) {
	fmt.Println("card read", card)
	t.io.buzzer.Write("read-beep")
	if user, found := t.users[model.UserId(card)]; found {
		if t.pay == nil {
			t.io.display.Write(fmt.Sprintf("user [%s] has %d$", user.Name(), user.Balance()))
		} else {
			if err := t.processPayment(user, t.pay); err == nil {
				t.pay = nil
				t.io.buzzer.Write("success-beep")
				t.io.display.Write("Payment successfull")
				//time.Sleep(3 * time.Second)
			} else {
				t.io.buzzer.Write("error-beep")
				t.io.display.Write(err.Error())
				t.io.display.Write("cancelling payment")
				t.pay = nil
			}
		}
	} else {
		t.io.buzzer.Write("error-beep")
		fmt.Printf("user with key [%s] not found\n", card)
		t.io.display.Write("User not found!")
	}
}

func (t *Term) processPayment(user *model.Account, payment *payment) error {
	if amount, err := payment.amount(); err != nil {
		return err
	} else {
		if newBalance, err := user.Withdraw(amount); err == nil {
			t.io.display.Write("payment successful")
			t.io.display.Write(fmt.Sprintf("new balance of user [%s] is %d,- check %d", user.Name(), newBalance, user.Balance()))
			return nil
		} else {
			return err
		}
	}
}

func (t *Term) keyPressed(key string) {
	if t.pay == nil {
		t.pay = &payment{amountString: ""}
	}

	if err := t.pay.readKey(key); err == nil {
		t.io.buzzer.Write("key-beep")
		t.io.display.Write(fmt.Sprintf("%s,-", t.pay.amountString))
	} else {
		t.io.buzzer.Write("cancel-beep")
		fmt.Println("Cancel payment")
		t.io.display.Write("Payment cancelled ...")
		t.pay = nil
	}
}

func (t *Term) read(inputChannel chan string, inputReader peripherals.InputReader) {
	inputChannel <- inputReader.Read()
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
