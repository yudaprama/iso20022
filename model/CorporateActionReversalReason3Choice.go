package model

// Choice between a standard code or proprietary code to specify the reason for the reversal.
type CorporateActionReversalReason3Choice struct {

	// Standard code to specify the reason for the reversal.
	Code *CorporateActionReversalReason1Code `xml:"Cd"`

	// Proprietary identification for the reason of the reversal.
	Proprietary *GenericIdentification30 `xml:"Prtry"`
}

func (c *CorporateActionReversalReason3Choice) SetCode(value string) {
	c.Code = (*CorporateActionReversalReason1Code)(&value)
}

func (c *CorporateActionReversalReason3Choice) AddProprietary() *GenericIdentification30 {
	c.Proprietary = new(GenericIdentification30)
	return c.Proprietary
}
