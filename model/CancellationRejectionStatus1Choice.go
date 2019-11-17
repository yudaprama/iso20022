package model

// Choice of cancellation rejection status.
type CancellationRejectionStatus1Choice struct {

	// Reason advising the rejection of the instruction cancellation request in the form of a code.
	Code *RejectionReason2Code `xml:"Cd"`

	// This code can be used in case another reason is required.
	Proprietary *GenericIdentification13 `xml:"Prtry"`
}

func (c *CancellationRejectionStatus1Choice) SetCode(value string) {
	c.Code = (*RejectionReason2Code)(&value)
}

func (c *CancellationRejectionStatus1Choice) AddProprietary() *GenericIdentification13 {
	c.Proprietary = new(GenericIdentification13)
	return c.Proprietary
}
