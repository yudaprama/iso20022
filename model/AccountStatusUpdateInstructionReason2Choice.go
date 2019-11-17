package model

// Choice of formats for a reason for an instruction to change the status of an account.
type AccountStatusUpdateInstructionReason2Choice struct {

	// Reason for the instruction to change the account status expressed as a code.
	Code *AccountStatusUpdateRequestReason1Code `xml:"Cd"`

	// Reason for the  instruction to change the account status expressed as a proprietary code.
	Proprietary *GenericIdentification36 `xml:"Prtry"`
}

func (a *AccountStatusUpdateInstructionReason2Choice) SetCode(value string) {
	a.Code = (*AccountStatusUpdateRequestReason1Code)(&value)
}

func (a *AccountStatusUpdateInstructionReason2Choice) AddProprietary() *GenericIdentification36 {
	a.Proprietary = new(GenericIdentification36)
	return a.Proprietary
}
