package model

// Choice of formats for the designation of an account.
type AccountDesignation1Choice struct {

	// Account designation expressed as a code.
	Code *Rank1Code `xml:"Cd"`

	// Account designation expressed as a proprietary code.
	Proprietary *GenericIdentification47 `xml:"Prtry"`
}

func (a *AccountDesignation1Choice) SetCode(value string) {
	a.Code = (*Rank1Code)(&value)
}

func (a *AccountDesignation1Choice) AddProprietary() *GenericIdentification47 {
	a.Proprietary = new(GenericIdentification47)
	return a.Proprietary
}
