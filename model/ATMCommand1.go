package model

// Maintenance command to perform on an ATM.
type ATMCommand1 struct {

	// Type of command to be performed by the ATM.
	Type *ATMCommand1Code `xml:"Tp"`

	// Urgency of the command.
	Urgency *TMSContactLevel2Code `xml:"Urgcy"`

	// Date time on which the command must be performed.
	DateTime *ISODateTime `xml:"DtTm,omitempty"`

	// Identification of the entity issuing the command.
	CommandIdentification *ATMCommandIdentification1 `xml:"CmdId,omitempty"`

	// Specific parameters attached to the command.
	CommandParameters *ATMCommandParameters1Choice `xml:"CmdParams,omitempty"`
}

func (a *ATMCommand1) SetType(value string) {
	a.Type = (*ATMCommand1Code)(&value)
}

func (a *ATMCommand1) SetUrgency(value string) {
	a.Urgency = (*TMSContactLevel2Code)(&value)
}

func (a *ATMCommand1) SetDateTime(value string) {
	a.DateTime = (*ISODateTime)(&value)
}

func (a *ATMCommand1) AddCommandIdentification() *ATMCommandIdentification1 {
	a.CommandIdentification = new(ATMCommandIdentification1)
	return a.CommandIdentification
}

func (a *ATMCommand1) AddCommandParameters() *ATMCommandParameters1Choice {
	a.CommandParameters = new(ATMCommandParameters1Choice)
	return a.CommandParameters
}
