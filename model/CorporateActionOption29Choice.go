package model

// Choice between a standard code or a proprietary code to specify the type of the corporate action option.
type CorporateActionOption29Choice struct {

	// Specifies the corporate action options available to the account owner.
	Code *CorporateActionOption9Code `xml:"Cd"`

	// Proprietary identification of the type of corporate action option.
	Proprietary *GenericIdentification47 `xml:"Prtry"`
}

func (c *CorporateActionOption29Choice) SetCode(value string) {
	c.Code = (*CorporateActionOption9Code)(&value)
}

func (c *CorporateActionOption29Choice) AddProprietary() *GenericIdentification47 {
	c.Proprietary = new(GenericIdentification47)
	return c.Proprietary
}
