package model

// Choice of format for the investor capacity.
type InvestorCapacity5Choice struct {

	// Investor capacity expressed as an ISO 20022 code.
	Code *Eligibility1Code `xml:"Cd"`

	// Investor capacity expressed as a proprietary code.
	Proprietary *GenericIdentification47 `xml:"Prtry"`
}

func (i *InvestorCapacity5Choice) SetCode(value string) {
	i.Code = (*Eligibility1Code)(&value)
}

func (i *InvestorCapacity5Choice) AddProprietary() *GenericIdentification47 {
	i.Proprietary = new(GenericIdentification47)
	return i.Proprietary
}
