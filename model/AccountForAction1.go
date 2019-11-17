package model

// Account to or from which a cash entry is made.
type AccountForAction1 struct {

	// Unique and unambiguous identification for the account between the account owner and the account servicer.
	Identification *AccountIdentification4Choice `xml:"Id"`

	// Medium of exchange of value.
	Currency *ActiveCurrencyCode `xml:"Ccy"`
}

func (a *AccountForAction1) AddIdentification() *AccountIdentification4Choice {
	a.Identification = new(AccountIdentification4Choice)
	return a.Identification
}

func (a *AccountForAction1) SetCurrency(value string) {
	a.Currency = (*ActiveCurrencyCode)(&value)
}
