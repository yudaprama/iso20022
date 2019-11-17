package model

// Choice of format for the charge tax basis.
type ChargeTaxBasisType1Choice struct {

	// Securities transaction type expressed as an ISO 20022 code.
	Code *ChargeTaxBasis1Code `xml:"Cd"`

	// Securities transaction type expressed as a proprietary code.
	Proprietary *GenericIdentification38 `xml:"Prtry"`
}

func (c *ChargeTaxBasisType1Choice) SetCode(value string) {
	c.Code = (*ChargeTaxBasis1Code)(&value)
}

func (c *ChargeTaxBasisType1Choice) AddProprietary() *GenericIdentification38 {
	c.Proprietary = new(GenericIdentification38)
	return c.Proprietary
}
