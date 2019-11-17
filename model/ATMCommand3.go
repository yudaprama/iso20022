package model

// Maintenance command which has requested the device report.
type ATMCommand3 struct {

	// Type of command to be performed by the ATM.
	Type *ATMCommand2Code `xml:"Tp"`

	// Identification of the entity issuing the command.
	CommandIdentification *ATMCommandIdentification1 `xml:"CmdId,omitempty"`
}

func (a *ATMCommand3) SetType(value string) {
	a.Type = (*ATMCommand2Code)(&value)
}

func (a *ATMCommand3) AddCommandIdentification() *ATMCommandIdentification1 {
	a.CommandIdentification = new(ATMCommandIdentification1)
	return a.CommandIdentification
}
