package model

// Choice between a standard code or a proprietary code to specify the type of the corporate action option.
type CorporateActionOption12Choice struct {

	// Specifies the corporate action options available to the account owner.
	Code *CorporateActionOption9Code `xml:"Cd"`

	// Proprietary identification of the type of corporate action option.
	Proprietary *GenericIdentification20 `xml:"Prtry"`
}

func (c *CorporateActionOption12Choice) SetCode(value string) {
	c.Code = (*CorporateActionOption9Code)(&value)
}

func (c *CorporateActionOption12Choice) AddProprietary() *GenericIdentification20 {
	c.Proprietary = new(GenericIdentification20)
	return c.Proprietary
}
