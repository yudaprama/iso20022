package model

// Choice of formats for the specification of the type of account usage.
type AccountUsageType2Choice struct {

	// Type of account usage expressed as a code.
	Code *AccountUsageType2Code `xml:"Cd"`

	// Type of account usage expressed as a proprietary code.
	Proprietary *GenericIdentification47 `xml:"Prtry"`
}

func (a *AccountUsageType2Choice) SetCode(value string) {
	a.Code = (*AccountUsageType2Code)(&value)
}

func (a *AccountUsageType2Choice) AddProprietary() *GenericIdentification47 {
	a.Proprietary = new(GenericIdentification47)
	return a.Proprietary
}
