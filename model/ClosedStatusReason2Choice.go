package model

// Choice of formats for a closed reason code.
type ClosedStatusReason2Choice struct {

	// Reason for the closed account status expressed as a code.
	Code *ClosedStatusReason1Code `xml:"Cd"`

	// Reason for the closed account status expressed as a proprietary code.
	Proprietary *GenericIdentification36 `xml:"Prtry"`
}

func (c *ClosedStatusReason2Choice) SetCode(value string) {
	c.Code = (*ClosedStatusReason1Code)(&value)
}

func (c *ClosedStatusReason2Choice) AddProprietary() *GenericIdentification36 {
	c.Proprietary = new(GenericIdentification36)
	return c.Proprietary
}
