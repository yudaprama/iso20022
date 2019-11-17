package model

// Choice between a standard code or a proprietary code to specify the type of the corporate action option.
type CorporateActionOption21Choice struct {

	// Specifies the corporate action options available to the account owner.
	Code *CorporateActionOption10Code `xml:"Cd"`

	// Proprietary identification of the type of corporate action option.
	Proprietary *GenericIdentification30 `xml:"Prtry"`
}

func (c *CorporateActionOption21Choice) SetCode(value string) {
	c.Code = (*CorporateActionOption10Code)(&value)
}

func (c *CorporateActionOption21Choice) AddProprietary() *GenericIdentification30 {
	c.Proprietary = new(GenericIdentification30)
	return c.Proprietary
}
