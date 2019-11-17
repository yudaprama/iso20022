package model

// Choice between a standard code or a proprietary code to specify the type of the corporate action option.
type CorporateActionOption20Choice struct {

	// Specifies the corporate action options available to the account owner.
	Code *CorporateActionOption9Code `xml:"Cd"`

	// Proprietary identification of the type of corporate action option.
	Proprietary *GenericIdentification30 `xml:"Prtry"`
}

func (c *CorporateActionOption20Choice) SetCode(value string) {
	c.Code = (*CorporateActionOption9Code)(&value)
}

func (c *CorporateActionOption20Choice) AddProprietary() *GenericIdentification30 {
	c.Proprietary = new(GenericIdentification30)
	return c.Proprietary
}
