package model

// Choice between a standard code or proprietary code to specify the reason for the reversal.
type CorporateActionReversalReason4Choice struct {

	// Standard code to specify the reason for the reversal.
	Code *CorporateActionReversalReason1Code `xml:"Cd"`

	// Proprietary identification for the reason of the reversal.
	Proprietary *GenericIdentification47 `xml:"Prtry"`
}

func (c *CorporateActionReversalReason4Choice) SetCode(value string) {
	c.Code = (*CorporateActionReversalReason1Code)(&value)
}

func (c *CorporateActionReversalReason4Choice) AddProprietary() *GenericIdentification47 {
	c.Proprietary = new(GenericIdentification47)
	return c.Proprietary
}
