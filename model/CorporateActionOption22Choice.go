package model

// Choice between a standard code or a proprietary code to specify the type of the corporate action option.
type CorporateActionOption22Choice struct {

	// Specifies the corporate action options available to the account owner.
	Code *CorporateActionOption10Code `xml:"Cd"`

	// Proprietary identification of the type of corporate action option.
	Proprietary *GenericIdentification47 `xml:"Prtry"`
}

func (c *CorporateActionOption22Choice) SetCode(value string) {
	c.Code = (*CorporateActionOption10Code)(&value)
}

func (c *CorporateActionOption22Choice) AddProprietary() *GenericIdentification47 {
	c.Proprietary = new(GenericIdentification47)
	return c.Proprietary
}
