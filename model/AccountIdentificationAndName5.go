package model

// Identification of the account expressed with a name and an account number.
type AccountIdentificationAndName5 struct {

	// Unique and unambiguous identification for the account between the account owner and the account servicer.
	Identification *AccountIdentification4Choice `xml:"Id"`

	// Name of the account. It provides an additional means of identification, and is designated by the account servicer in agreement with the account owner.
	Name *Max35Text `xml:"Nm,omitempty"`
}

func (a *AccountIdentificationAndName5) AddIdentification() *AccountIdentification4Choice {
	a.Identification = new(AccountIdentification4Choice)
	return a.Identification
}

func (a *AccountIdentificationAndName5) SetName(value string) {
	a.Name = (*Max35Text)(&value)
}
