package model

// Choice of formats for the status of the investor profile.
type InvestorProfileStatus1Choice struct {

	// Investor profile status expressed as a code.
	Code *InvestorProfileStatus1Code `xml:"Cd"`

	// Investor profile status expressed as a proprietary code.
	Proprietary *GenericIdentification47 `xml:"Prtry"`
}

func (i *InvestorProfileStatus1Choice) SetCode(value string) {
	i.Code = (*InvestorProfileStatus1Code)(&value)
}

func (i *InvestorProfileStatus1Choice) AddProprietary() *GenericIdentification47 {
	i.Proprietary = new(GenericIdentification47)
	return i.Proprietary
}
