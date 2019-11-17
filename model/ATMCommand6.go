package model

// Party which has requested the reconciliation.
type ATMCommand6 struct {

	// Type of command to be performed by the ATM.
	Type *ATMCommand3Code `xml:"Tp"`

	// Identification of the entity issuing the command.
	CommandIdentification *ATMCommandIdentification1 `xml:"CmdId,omitempty"`
}

func (a *ATMCommand6) SetType(value string) {
	a.Type = (*ATMCommand3Code)(&value)
}

func (a *ATMCommand6) AddCommandIdentification() *ATMCommandIdentification1 {
	a.CommandIdentification = new(ATMCommandIdentification1)
	return a.CommandIdentification
}
