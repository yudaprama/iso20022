package model

// Choice between formats for a cancelled reason.
type CancelledReason12Choice struct {

	// Cancelled reason expressed as a code.
	Code *CancelledStatusReason2Code `xml:"Cd"`

	// Cancelled reason expressed as a proprietary code.
	Proprietary *GenericIdentification1 `xml:"Prtry"`

	// No reason available or to report for the cancelled status.
	NoSpecifiedReason *NoReasonCode `xml:"NoSpcfdRsn"`
}

func (c *CancelledReason12Choice) SetCode(value string) {
	c.Code = (*CancelledStatusReason2Code)(&value)
}

func (c *CancelledReason12Choice) AddProprietary() *GenericIdentification1 {
	c.Proprietary = new(GenericIdentification1)
	return c.Proprietary
}

func (c *CancelledReason12Choice) SetNoSpecifiedReason(value string) {
	c.NoSpecifiedReason = (*NoReasonCode)(&value)
}
