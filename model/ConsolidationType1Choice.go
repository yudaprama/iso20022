package model

// Choice of formats for the type of consolidation.
type ConsolidationType1Choice struct {

	// Consolidation type expressed as a code.
	Code *ConsolidationType1Code `xml:"Cd"`

	// Consolidation type expressed as a proprietary code.
	Proprietary *GenericIdentification47 `xml:"Prtry"`
}

func (c *ConsolidationType1Choice) SetCode(value string) {
	c.Code = (*ConsolidationType1Code)(&value)
}

func (c *ConsolidationType1Choice) AddProprietary() *GenericIdentification47 {
	c.Proprietary = new(GenericIdentification47)
	return c.Proprietary
}
