package model

// Choice of formats for the account opening type.
type AccountOpeningType1Choice struct {

	// Type of account opening instruction expressed as a code.
	Code *AccountOpeningType1Code `xml:"Cd"`

	// Type of account opening instruction expressed as a proprietary code.
	Proprietary *GenericIdentification47 `xml:"Prtry"`
}

func (a *AccountOpeningType1Choice) SetCode(value string) {
	a.Code = (*AccountOpeningType1Code)(&value)
}

func (a *AccountOpeningType1Choice) AddProprietary() *GenericIdentification47 {
	a.Proprietary = new(GenericIdentification47)
	return a.Proprietary
}
