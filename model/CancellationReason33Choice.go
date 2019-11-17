package model

// Specifies the reason for the cancellation request.
type CancellationReason33Choice struct {

	// Reason for the cancellation request, in a coded form.
	Code *ExternalCancellationReason1Code `xml:"Cd"`

	// Reason for the cancellation request, in a proprietary form
	Proprietary *Max35Text `xml:"Prtry"`
}

func (c *CancellationReason33Choice) SetCode(value string) {
	c.Code = (*ExternalCancellationReason1Code)(&value)
}

func (c *CancellationReason33Choice) SetProprietary(value string) {
	c.Proprietary = (*Max35Text)(&value)
}
