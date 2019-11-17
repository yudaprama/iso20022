package model

// Posting of an item to a cash account, in the context of a cash transaction, that results in an increase or decrease to the balance of the account.
type AmountAndDirection30 struct {

	// Amount of money in the cash entry.
	Amount *ActiveOrHistoricCurrencyAndAmount `xml:"Amt"`

	// Indicates that the amount value is positive or negative.
	Sign *PlusOrMinusIndicator `xml:"Sgn,omitempty"`
}

func (a *AmountAndDirection30) SetAmount(value, currency string) {
	a.Amount = NewActiveOrHistoricCurrencyAndAmount(value, currency)
}

func (a *AmountAndDirection30) SetSign(value string) {
	a.Sign = (*PlusOrMinusIndicator)(&value)
}
