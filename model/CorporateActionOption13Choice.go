package model

// Choice between a standard code or a proprietary code to specify the type of the corporate action option.
type CorporateActionOption13Choice struct {

	// Specifies the corporate action options available to the account owner.
	Code *CorporateActionOption10Code `xml:"Cd"`

	// Proprietary identification of the type of corporate action option.
	Proprietary *GenericIdentification20 `xml:"Prtry"`
}

func (c *CorporateActionOption13Choice) SetCode(value string) {
	c.Code = (*CorporateActionOption10Code)(&value)
}

func (c *CorporateActionOption13Choice) AddProprietary() *GenericIdentification20 {
	c.Proprietary = new(GenericIdentification20)
	return c.Proprietary
}
