package model

// Choice of formats for the specification of the type of account.
type AccountType1Choice struct {

	// Type of account expressed as a code.
	Code *FundCashAccount3Code `xml:"Cd"`

	// Type of account expressed as a proprietary code.
	Proprietary *GenericIdentification47 `xml:"Prtry"`
}

func (a *AccountType1Choice) SetCode(value string) {
	a.Code = (*FundCashAccount3Code)(&value)
}

func (a *AccountType1Choice) AddProprietary() *GenericIdentification47 {
	a.Proprietary = new(GenericIdentification47)
	return a.Proprietary
}
