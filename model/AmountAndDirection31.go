package model

// Posting of an item to a cash account, in the context of a cash transaction, that results in an increase or decrease to the balance of the account.
type AmountAndDirection31 struct {

	// Currency and value.
	Amount *ActiveOrHistoricCurrencyAndAmount `xml:"Amt"`

	// Indication that the position is short or long.
	ShortLongIndicator *ShortLong1Code `xml:"ShrtLngInd,omitempty"`
}

func (a *AmountAndDirection31) SetAmount(value, currency string) {
	a.Amount = NewActiveOrHistoricCurrencyAndAmount(value, currency)
}

func (a *AmountAndDirection31) SetShortLongIndicator(value string) {
	a.ShortLongIndicator = (*ShortLong1Code)(&value)
}
