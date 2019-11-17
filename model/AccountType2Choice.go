package model

// Choice of formats for the specification of the type of account.
type AccountType2Choice struct {

	// Type of account expressed as a code.
	Code *FundCashAccount4Code `xml:"Cd"`

	// Type of account expressed as a proprietary code.
	Proprietary *GenericIdentification47 `xml:"Prtry"`
}

func (a *AccountType2Choice) SetCode(value string) {
	a.Code = (*FundCashAccount4Code)(&value)
}

func (a *AccountType2Choice) AddProprietary() *GenericIdentification47 {
	a.Proprietary = new(GenericIdentification47)
	return a.Proprietary
}
