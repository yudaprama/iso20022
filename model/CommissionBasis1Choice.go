package model

// Choice of formats for the commission basis.
type CommissionBasis1Choice struct {

	// Commission basis expressed as a code.
	Code *TaxationBasis4Code `xml:"Cd"`

	// Commission basis expressed as a proprietary code.
	Proprietary *GenericIdentification47 `xml:"Prtry"`
}

func (c *CommissionBasis1Choice) SetCode(value string) {
	c.Code = (*TaxationBasis4Code)(&value)
}

func (c *CommissionBasis1Choice) AddProprietary() *GenericIdentification47 {
	c.Proprietary = new(GenericIdentification47)
	return c.Proprietary
}
