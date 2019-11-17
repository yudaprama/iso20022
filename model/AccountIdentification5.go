package model

// Business relationship between two entities; one entity is the account owner, the other entity is the account servicer.
type AccountIdentification5 struct {

	// Unambiguous identification for the account between the account owner and the account servicer.
	Identification *Max35Text `xml:"Id"`

	// Description of the account.
	Name *Max35Text `xml:"Nm,omitempty"`

	// Specifies the type of account.
	Type *GenericIdentification30 `xml:"Tp,omitempty"`
}

func (a *AccountIdentification5) SetIdentification(value string) {
	a.Identification = (*Max35Text)(&value)
}

func (a *AccountIdentification5) SetName(value string) {
	a.Name = (*Max35Text)(&value)
}

func (a *AccountIdentification5) AddType() *GenericIdentification30 {
	a.Type = new(GenericIdentification30)
	return a.Type
}
