package model

// Speficies the reason for the transaction cancellation status.
type CancellationStatusReason1Choice struct {

	// Reason for the cancellation status, in a coded form.
	Code *PaymentCancellationRejection1Code `xml:"Cd"`

	// Reason for the status, in a proprietary form.
	Proprietary *Max35Text `xml:"Prtry"`
}

func (c *CancellationStatusReason1Choice) SetCode(value string) {
	c.Code = (*PaymentCancellationRejection1Code)(&value)
}

func (c *CancellationStatusReason1Choice) SetProprietary(value string) {
	c.Proprietary = (*Max35Text)(&value)
}
