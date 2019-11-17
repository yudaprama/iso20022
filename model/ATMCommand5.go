package model

// Command result for reinitialization of the transaction counters.
type ATMCommand5 struct {

	// Type of command to be performed by the ATM.
	Type *ATMCommand3Code `xml:"Tp"`

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

func (a *ATMCommand5) SetType(value string) {
	a.Type = (*ATMCommand3Code)(&value)
}

func (a *ATMCommand5) SetRequiredDateTime(value string) {
	a.RequiredDateTime = (*ISODateTime)(&value)
}

func (a *ATMCommand5) SetProcessedDateTime(value string) {
	a.ProcessedDateTime = (*ISODateTime)(&value)
}

func (a *ATMCommand5) AddCommandIdentification() *ATMCommandIdentification1 {
	a.CommandIdentification = new(ATMCommandIdentification1)
	return a.CommandIdentification
}

func (a *ATMCommand5) SetResult(value string) {
	a.Result = (*TerminalManagementActionResult2Code)(&value)
}

func (a *ATMCommand5) SetAdditionalErrorInformation(value string) {
	a.AdditionalErrorInformation = (*Max140Text)(&value)
}
