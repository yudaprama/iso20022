package model

// Choice of formats for the accounting status.
type AccountingStatus1Choice struct {

	// Accounting status expressed as a code.
	Code *AccountingStatus1Code `xml:"Cd"`

	// Accounting status expressed as a proprietary code.
	Proprietary *GenericIdentification47 `xml:"Prtry"`
}

func (a *AccountingStatus1Choice) SetCode(value string) {
	a.Code = (*AccountingStatus1Code)(&value)
}

func (a *AccountingStatus1Choice) AddProprietary() *GenericIdentification47 {
	a.Proprietary = new(GenericIdentification47)
	return a.Proprietary
}
