package model

// Result of a maintenance command performed by the ATM.
type ATMCommand2 struct {

	// Type of command to be performed by the ATM.
	Type *ATMCommand2Code `xml:"Tp"`

	// Date time on which the command has been requested to be performed.
	RequiredDateTime *ISODateTime `xml:"ReqrdDtTm,omitempty"`

	// Date time on which the command has been performed.
	ProcessedDateTime *ISODateTime `xml:"PrcdDtTm"`

	// Identification of the entity issuing the command.
	CommandIdentification *ATMCommandIdentification1 `xml:"CmdId,omitempty"`

	// Final result of the processed command at the ATM.
	Result *TerminalManagementActionResult2Code `xml:"Rslt"`

	// Additional information on the failure to be logged for further examination.
	AdditionalErrorInformation *Max140Text `xml:"AddtlErrInf,omitempty"`
}

func (a *ATMCommand2) SetType(value string) {
	a.Type = (*ATMCommand2Code)(&value)
}

func (a *ATMCommand2) SetRequiredDateTime(value string) {
	a.RequiredDateTime = (*ISODateTime)(&value)
}

func (a *ATMCommand2) SetProcessedDateTime(value string) {
	a.ProcessedDateTime = (*ISODateTime)(&value)
}

func (a *ATMCommand2) AddCommandIdentification() *ATMCommandIdentification1 {
	a.CommandIdentification = new(ATMCommandIdentification1)
	return a.CommandIdentification
}

func (a *ATMCommand2) SetResult(value string) {
	a.Result = (*TerminalManagementActionResult2Code)(&value)
}

func (a *ATMCommand2) SetAdditionalErrorInformation(value string) {
	a.AdditionalErrorInformation = (*Max140Text)(&value)
}
