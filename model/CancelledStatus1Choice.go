package model

// Choice between a reason or no reason for the corporate action instruction or instruction cancellation cancelled/cancellation completed status.
type CancelledStatus1Choice struct {

	// Reason not specified.
	NoSpecifiedReason *NoReasonCode `xml:"NoSpcfdRsn"`

	// Reason for the cancelled status.
	Reason []*CancelledStatusReason4 `xml:"Rsn"`
}

func (c *CancelledStatus1Choice) SetNoSpecifiedReason(value string) {
	c.NoSpecifiedReason = (*NoReasonCode)(&value)
}

func (c *CancelledStatus1Choice) AddReason() *CancelledStatusReason4 {
	newValue := new(CancelledStatusReason4)
	c.Reason = append(c.Reason, newValue)
	return newValue
}
