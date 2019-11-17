package model

// Choice between formats for a partially settled reason.
type PartiallySettled21Choice struct {

	// Partially settled reason expressed as a code.
	Code *SettledStatusReason2Code `xml:"Cd"`

	// Partially settled reason expressed as a proprietary code.
	Proprietary *GenericIdentification1 `xml:"Prtry"`
}

func (p *PartiallySettled21Choice) SetCode(value string) {
	p.Code = (*SettledStatusReason2Code)(&value)
}

func (p *PartiallySettled21Choice) AddProprietary() *GenericIdentification1 {
	p.Proprietary = new(GenericIdentification1)
	return p.Proprietary
}
