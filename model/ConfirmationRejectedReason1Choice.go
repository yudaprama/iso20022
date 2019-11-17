package model

// Choice of formats for a rejection reason for a confirmation.
type ConfirmationRejectedReason1Choice struct {

	// Reason for the rejected status expressed as a code.
	Code *RejectedConfirmationStatusReason1Code `xml:"Cd"`

	// Reason for the rejected status.
	Proprietary *GenericIdentification1 `xml:"Prtry"`
}

func (c *ConfirmationRejectedReason1Choice) SetCode(value string) {
	c.Code = (*RejectedConfirmationStatusReason1Code)(&value)
}

func (c *ConfirmationRejectedReason1Choice) AddProprietary() *GenericIdentification1 {
	c.Proprietary = new(GenericIdentification1)
	return c.Proprietary
}
