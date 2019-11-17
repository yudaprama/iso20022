package model

// Maintenance command the ATM must perform.
type ATMCommand4 struct {

	// Type of command to be performed by the ATM.
	Type *ATMCommand2Code `xml:"Tp"`

	// Urgency of the command.
	Urgency *TMSContactLevel2Code `xml:"Urgcy"`

	// Date time on which the command must be performed.
	DateTime *ISODateTime `xml:"DtTm,omitempty"`

	// Identification of the entity issuing the command.
	CommandIdentification *ATMCommandIdentification1 `xml:"CmdId,omitempty"`

	// Reason for sending the command.
	Reason *ATMCommandReason1Code `xml:"Rsn,omitempty"`

	// Trace of reasons by the entities in the path from the origin of the command to the ATM.
	TraceReason []*ATMCommandReason1Code `xml:"TracRsn,omitempty"`

	// Additional information about the reason to request this command.
	AdditionalReasonInformation *Max70Text `xml:"AddtlRsnInf,omitempty"`

	// Specific parameters attached to the command.
	CommandParameters *ATMCommandParameters2Choice `xml:"CmdParams,omitempty"`
}

func (a *ATMCommand4) SetType(value string) {
	a.Type = (*ATMCommand2Code)(&value)
}

func (a *ATMCommand4) SetUrgency(value string) {
	a.Urgency = (*TMSContactLevel2Code)(&value)
}

func (a *ATMCommand4) SetDateTime(value string) {
	a.DateTime = (*ISODateTime)(&value)
}

func (a *ATMCommand4) AddCommandIdentification() *ATMCommandIdentification1 {
	a.CommandIdentification = new(ATMCommandIdentification1)
	return a.CommandIdentification
}

func (a *ATMCommand4) SetReason(value string) {
	a.Reason = (*ATMCommandReason1Code)(&value)
}

func (a *ATMCommand4) AddTraceReason(value string) {
	a.TraceReason = append(a.TraceReason, (*ATMCommandReason1Code)(&value))
}

func (a *ATMCommand4) SetAdditionalReasonInformation(value string) {
	a.AdditionalReasonInformation = (*Max70Text)(&value)
}

func (a *ATMCommand4) AddCommandParameters() *ATMCommandParameters2Choice {
	a.CommandParameters = new(ATMCommandParameters2Choice)
	return a.CommandParameters
}
