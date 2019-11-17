package model

// Choice of formats to  express the type of changes.
type CorporateActionChangeType1FormatChoice struct {

	// Standard code to  specify the type of changes.
	Code *CorporateActionChangeType1Code `xml:"Cd"`

	// Proprietary code to  express the type of changes.
	Proprietary *GenericIdentification13 `xml:"Prtry"`
}

func (c *CorporateActionChangeType1FormatChoice) SetCode(value string) {
	c.Code = (*CorporateActionChangeType1Code)(&value)
}

func (c *CorporateActionChangeType1FormatChoice) AddProprietary() *GenericIdentification13 {
	c.Proprietary = new(GenericIdentification13)
	return c.Proprietary
}
