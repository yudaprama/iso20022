package model

// Reason for an update to an account status.
type AccountStatusUpdateInstructionReason1 struct {

	// Reason for the instruction to change the account status.
	Code *AccountStatusUpdateInstructionReason2Choice `xml:"Cd,omitempty"`

	// Additional information about the reason for the instruction to change the account status.
	AdditionalInformation *Max350Text `xml:"AddtlInf,omitempty"`
}

func (a *AccountStatusUpdateInstructionReason1) AddCode() *AccountStatusUpdateInstructionReason2Choice {
	a.Code = new(AccountStatusUpdateInstructionReason2Choice)
	return a.Code
}

func (a *AccountStatusUpdateInstructionReason1) SetAdditionalInformation(value string) {
	a.AdditionalInformation = (*Max350Text)(&value)
}
