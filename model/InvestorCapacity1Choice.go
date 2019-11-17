package model

// Choice of format for the investor capacity.
type InvestorCapacity1Choice struct {

	// Investor capacity expressed as an ISO 20022 code.
	Code *Eligibility1Code `xml:"Cd"`

	// Investor capacity expressed as a proprietary code.
	Proprietary *GenericIdentification20 `xml:"Prtry"`
}

func (i *InvestorCapacity1Choice) SetCode(value string) {
	i.Code = (*Eligibility1Code)(&value)
}

func (i *InvestorCapacity1Choice) AddProprietary() *GenericIdentification20 {
	i.Proprietary = new(GenericIdentification20)
	return i.Proprietary
}
