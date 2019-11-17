package model

// Choice of formats for a closure pending reason code.
type ClosurePendingStatusReason2Choice struct {

	// Reason for the closure pending status expressed as a code.
	Code *ClosurePendingStatusReason1Code `xml:"Cd"`

	// Reason for the closure pending status expressed as a proprietary code.
	Proprietary *GenericIdentification36 `xml:"Prtry"`
}

func (c *ClosurePendingStatusReason2Choice) SetCode(value string) {
	c.Code = (*ClosurePendingStatusReason1Code)(&value)
}

func (c *ClosurePendingStatusReason2Choice) AddProprietary() *GenericIdentification36 {
	c.Proprietary = new(GenericIdentification36)
	return c.Proprietary
}
