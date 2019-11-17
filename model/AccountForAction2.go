package model

// Account to or from which a cash entry is made.
type AccountForAction2 struct {

	// Unique and unambiguous identification for the account between the account owner and the account servicer.
	Identification *AccountIdentification4Choice `xml:"Id"`

	// Name of the account. In case of conflict between the Account Identification and the Account Name, it is recommended that the Account Servicer ask for clarification by means of the Request for Additional Information message.
	//
	//
	Name *Max70Text `xml:"Nm,omitempty"`

	// Medium of exchange of value.
	Currency *ActiveCurrencyCode `xml:"Ccy"`
}

func (a *AccountForAction2) AddIdentification() *AccountIdentification4Choice {
	a.Identification = new(AccountIdentification4Choice)
	return a.Identification
}

func (a *AccountForAction2) SetName(value string) {
	a.Name = (*Max70Text)(&value)
}

func (a *AccountForAction2) SetCurrency(value string) {
	a.Currency = (*ActiveCurrencyCode)(&value)
}
