package model

// Choice of format for the investor capacity.
type InvestorCapacity3Choice struct {

	// Investor capacity expressed as an ISO 20022 code.
	Code *Eligibility1Code `xml:"Cd"`

	// Investor capacity expressed as a proprietary code.
	Proprietary *GenericIdentification38 `xml:"Prtry"`
}

func (i *InvestorCapacity3Choice) SetCode(value string) {
	i.Code = (*Eligibility1Code)(&value)
}

func (i *InvestorCapacity3Choice) AddProprietary() *GenericIdentification38 {
	i.Proprietary = new(GenericIdentification38)
	return i.Proprietary
}
