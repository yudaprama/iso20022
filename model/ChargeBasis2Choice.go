package model

// Choice of formats for the charge or commission basis.
type ChargeBasis2Choice struct {

	// Fee (charge/commission) basis expressed as a code.
	Code *TaxationBasis5Code `xml:"Cd"`

	// Fee (charge/commission) basis expressed as a proprietary code.
	Proprietary *GenericIdentification47 `xml:"Prtry"`
}

func (c *ChargeBasis2Choice) SetCode(value string) {
	c.Code = (*TaxationBasis5Code)(&value)
}

func (c *ChargeBasis2Choice) AddProprietary() *GenericIdentification47 {
	c.Proprietary = new(GenericIdentification47)
	return c.Proprietary
}
