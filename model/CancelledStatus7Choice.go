package model

// Choice between a reason or no reason for the corporate action instruction or instruction cancellation cancelled/cancellation completed status.
type CancelledStatus7Choice struct {

	// Reason not specified.
	NoSpecifiedReason *NoReasonCode `xml:"NoSpcfdRsn"`

	// Reason for the cancelled status.
	Reason []*CancelledStatusReason8 `xml:"Rsn"`
}

func (c *CancelledStatus7Choice) SetNoSpecifiedReason(value string) {
	c.NoSpecifiedReason = (*NoReasonCode)(&value)
}

func (c *CancelledStatus7Choice) AddReason() *CancelledStatusReason8 {
	newValue := new(CancelledStatusReason8)
	c.Reason = append(c.Reason, newValue)
	return newValue
}
