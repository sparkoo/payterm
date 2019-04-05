package term

import (
	"fmt"
	"strconv"
)

type payment struct {
	amountString string
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

func (p *payment) amount() (int, error) {
	if amountNumber, err := strconv.ParseInt(p.amountString, 10, 32); err == nil {
		return int(amountNumber), nil
	} else {
		return 0, &IllegalAmountError{amountString: p.amountString, err: err}
	}
}

type IllegalAmountError struct {
	amountString string
	err          error
}

func (err *IllegalAmountError) Error() string {
	return fmt.Sprintf("Amount string can't be parsed. [%v], %v", err.amountString, err.err)
}
