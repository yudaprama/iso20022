package model

// Posting of an item to a cash account, in the context of a cash transaction, that results in an increase or decrease to the balance of the account.
type PayInScheduleItems1 struct {

	// Currency and amount to be paid in.
	Amount *ActiveCurrencyAndAmount `xml:"Amt"`

	// Time by which the amount must be paid in.
	Deadline *ISODateTime `xml:"Ddln"`
}

func (p *PayInScheduleItems1) SetAmount(value, currency string) {
	p.Amount = NewActiveCurrencyAndAmount(value, currency)
}

func (p *PayInScheduleItems1) SetDeadline(value string) {
	p.Deadline = (*ISODateTime)(&value)
}
