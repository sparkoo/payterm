package term

import (
	"fmt"
	"strconv"
)

type payment struct {
	amountString string
	amount       int
}

type CancelPaymentError struct {
}

func (*CancelPaymentError) Error() string {
	return "Payment cancelled"
}

func (p *payment) readKey(key string) error {
	// is number
	if _, err := strconv.ParseInt(key, 10, 32); err == nil {
		p.amountString = p.amountString + key
	} else {
		fmt.Println("pressed", key, "now what to do?")
		return &CancelPaymentError{}
	}
	return nil
}
