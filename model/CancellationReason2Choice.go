package model

// Specifies the reason for the cancellation request.
type CancellationReason2Choice struct {

	// Reason for the cancellation request, in a coded form.
	Code *CancellationReason4Code `xml:"Cd"`

	// Reason for the cancellation request, in a proprietary form
	Proprietary *Max35Text `xml:"Prtry"`
}

func (c *CancellationReason2Choice) SetCode(value string) {
	c.Code = (*CancellationReason4Code)(&value)
}

func (c *CancellationReason2Choice) SetProprietary(value string) {
	c.Proprietary = (*Max35Text)(&value)
}
