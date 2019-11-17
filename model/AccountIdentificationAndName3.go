package model

// Identification of the account expressed with a name and an account number.
type AccountIdentificationAndName3 struct {

	// Unique and unambiguous identification for the account between the account owner and the account servicer.
	Identification *CashAccountIdentification1Choice `xml:"Id"`

	// Name of the account. It provides an additional means of identification, and is designated by the account servicer in agreement with the account owner.
	Name *Max35Text `xml:"Nm,omitempty"`
}

func (a *AccountIdentificationAndName3) AddIdentification() *CashAccountIdentification1Choice {
	a.Identification = new(CashAccountIdentification1Choice)
	return a.Identification
}

func (a *AccountIdentificationAndName3) SetName(value string) {
	a.Name = (*Max35Text)(&value)
}
