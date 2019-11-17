package model

// Choice of format for the investor capacity.
type InvestorCapacity4Choice struct {

	// Investor capacity expressed as an ISO 20022 code.
	Code *Eligibility1Code `xml:"Cd"`

	// Investor capacity expressed as a proprietary code.
	Proprietary *GenericIdentification30 `xml:"Prtry"`
}

func (i *InvestorCapacity4Choice) SetCode(value string) {
	i.Code = (*Eligibility1Code)(&value)
}

func (i *InvestorCapacity4Choice) AddProprietary() *GenericIdentification30 {
	i.Proprietary = new(GenericIdentification30)
	return i.Proprietary
}
