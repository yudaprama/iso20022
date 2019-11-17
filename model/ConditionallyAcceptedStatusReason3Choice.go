package model

// Choice of formats for a conditionally accepted status.
type ConditionallyAcceptedStatusReason3Choice struct {

	// Conditionally accepted reason expressed as a code.
	Code *ConditionallyAcceptedStatusReason2Code `xml:"Cd"`

	// Conditionally accepted reason expressed as a proprietary code.
	Proprietary *GenericIdentification1 `xml:"Prtry"`
}

func (c *ConditionallyAcceptedStatusReason3Choice) SetCode(value string) {
	c.Code = (*ConditionallyAcceptedStatusReason2Code)(&value)
}

func (c *ConditionallyAcceptedStatusReason3Choice) AddProprietary() *GenericIdentification1 {
	c.Proprietary = new(GenericIdentification1)
	return c.Proprietary
}
