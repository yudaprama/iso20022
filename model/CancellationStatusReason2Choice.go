package model

// Specifies the reason for the transaction cancellation status.
type CancellationStatusReason2Choice struct {

	// Reason for the cancellation status, in a coded form.
	Code *PaymentCancellationRejection2Code `xml:"Cd"`

	// Reason for the status, in a proprietary form.
	Proprietary *Max35Text `xml:"Prtry"`
}

func (c *CancellationStatusReason2Choice) SetCode(value string) {
	c.Code = (*PaymentCancellationRejection2Code)(&value)
}

func (c *CancellationStatusReason2Choice) SetProprietary(value string) {
	c.Proprietary = (*Max35Text)(&value)
}
