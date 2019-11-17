package model

// Reason for the cancellation request of the transaction.
type CancellationReason1Choice struct {

	// Reason for the cancellation request in a coded form.
	Code *CancellationReason2Code `xml:"Cd"`

	// Reason for the cancellation request not catered for by the available codes.
	Proprietary *Max35Text `xml:"Prtry"`
}

func (c *CancellationReason1Choice) SetCode(value string) {
	c.Code = (*CancellationReason2Code)(&value)
}

func (c *CancellationReason1Choice) SetProprietary(value string) {
	c.Proprietary = (*Max35Text)(&value)
}
