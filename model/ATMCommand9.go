package model

// Party which has requested the reconciliation.
type ATMCommand9 struct {

	// Type of command to be performed by the ATM.
	Type *ATMCommand5Code `xml:"Tp"`

	// Identification of the entity issuing the command.
	CommandIdentification *ATMCommandIdentification1 `xml:"CmdId,omitempty"`
}

func (a *ATMCommand9) SetType(value string) {
	a.Type = (*ATMCommand5Code)(&value)
}

func (a *ATMCommand9) AddCommandIdentification() *ATMCommandIdentification1 {
	a.CommandIdentification = new(ATMCommandIdentification1)
	return a.CommandIdentification
}
