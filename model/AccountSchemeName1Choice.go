package model

// Sets of elements to identify a name of the identification scheme
type AccountSchemeName1Choice struct {

	// Name of the identification scheme, in a coded form as published in an external list.
	Code *ExternalAccountIdentification1Code `xml:"Cd"`

	// Name of the identification scheme, in a free text form.
	Proprietary *Max35Text `xml:"Prtry"`
}

func (a *AccountSchemeName1Choice) SetCode(value string) {
	a.Code = (*ExternalAccountIdentification1Code)(&value)
}

func (a *AccountSchemeName1Choice) SetProprietary(value string) {
	a.Proprietary = (*Max35Text)(&value)
}
