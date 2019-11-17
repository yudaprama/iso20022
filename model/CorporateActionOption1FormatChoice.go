package model

// Choice of format for the option type.
type CorporateActionOption1FormatChoice struct {

	// Option type expressed as a code.
	Code *CorporateActionOptionType1Code `xml:"Cd"`

	// Option type expressed as a prorprietary code.
	Proprietary *GenericIdentification13 `xml:"Prtry"`
}

func (c *CorporateActionOption1FormatChoice) SetCode(value string) {
	c.Code = (*CorporateActionOptionType1Code)(&value)
}

func (c *CorporateActionOption1FormatChoice) AddProprietary() *GenericIdentification13 {
	c.Proprietary = new(GenericIdentification13)
	return c.Proprietary
}
