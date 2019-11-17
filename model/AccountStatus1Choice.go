package model

// Choice of formats for the status of an account.
type AccountStatus1Choice struct {

	// Status of the account expressed as a code.
	Code *AccountStatus4Code `xml:"Cd"`

	// Status of the account expressed as a proprietary code.
	Proprietary *GenericIdentification47 `xml:"Prtry"`
}

func (a *AccountStatus1Choice) SetCode(value string) {
	a.Code = (*AccountStatus4Code)(&value)
}

func (a *AccountStatus1Choice) AddProprietary() *GenericIdentification47 {
	a.Proprietary = new(GenericIdentification47)
	return a.Proprietary
}
