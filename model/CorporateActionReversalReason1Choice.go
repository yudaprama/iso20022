package model

// Choice between a standard code or proprietary code to specify the reason for the reversal.
type CorporateActionReversalReason1Choice struct {

	// Standard code to specify the reason for the reversal.
	Code *CorporateActionReversalReason1Code `xml:"Cd"`

	// Proprietary identification for the reason of the reversal.
	Proprietary *GenericIdentification20 `xml:"Prtry"`
}

func (c *CorporateActionReversalReason1Choice) SetCode(value string) {
	c.Code = (*CorporateActionReversalReason1Code)(&value)
}

func (c *CorporateActionReversalReason1Choice) AddProprietary() *GenericIdentification20 {
	c.Proprietary = new(GenericIdentification20)
	return c.Proprietary
}
