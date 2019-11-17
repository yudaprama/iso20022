package model

// Choice of formats for the specification of the type of account usage.
type AccountUsageType1Choice struct {

	// Type of account usage expressed as a code.
	Code *AccountUsageType1Code `xml:"Cd"`

	// Type of account usage expressed as a proprietary code.
	Proprietary *GenericIdentification47 `xml:"Prtry"`
}

func (a *AccountUsageType1Choice) SetCode(value string) {
	a.Code = (*AccountUsageType1Code)(&value)
}

func (a *AccountUsageType1Choice) AddProprietary() *GenericIdentification47 {
	a.Proprietary = new(GenericIdentification47)
	return a.Proprietary
}
