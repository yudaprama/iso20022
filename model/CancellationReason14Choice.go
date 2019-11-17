package model

// Specifies the reason for the cancellation request.
type CancellationReason14Choice struct {

	// Reason for the cancellation request, in a coded form.
	Code *CancellationReason5Code `xml:"Cd"`

	// Reason for the cancellation request, in a proprietary form
	Proprietary *Max35Text `xml:"Prtry"`
}

func (c *CancellationReason14Choice) SetCode(value string) {
	c.Code = (*CancellationReason5Code)(&value)
}

func (c *CancellationReason14Choice) SetProprietary(value string) {
	c.Proprietary = (*Max35Text)(&value)
}
