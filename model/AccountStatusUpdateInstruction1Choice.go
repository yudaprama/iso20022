package model

// Choice of formats for an account status update instruction.
type AccountStatusUpdateInstruction1Choice struct {

	// Type of change to the account status expressed as a code.
	//
	Code *AccountStatusUpdateInstruction1Code `xml:"Cd"`

	// Type of change to the account status expressed as a proprietary code.
	Proprietary *GenericIdentification36 `xml:"Prtry"`
}

func (a *AccountStatusUpdateInstruction1Choice) SetCode(value string) {
	a.Code = (*AccountStatusUpdateInstruction1Code)(&value)
}

func (a *AccountStatusUpdateInstruction1Choice) AddProprietary() *GenericIdentification36 {
	a.Proprietary = new(GenericIdentification36)
	return a.Proprietary
}
