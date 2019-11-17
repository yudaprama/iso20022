package model

// Choice between a reason or no reason for the corporate action instruction or instruction cancellation cancelled/cancellation completed status.
type CancelledStatus3Choice struct {

	// Reason not specified.
	NoSpecifiedReason *NoReasonCode `xml:"NoSpcfdRsn"`

	// Reason for the cancelled status.
	Reason []*CancelledStatusReason6 `xml:"Rsn"`
}

func (c *CancelledStatus3Choice) SetNoSpecifiedReason(value string) {
	c.NoSpecifiedReason = (*NoReasonCode)(&value)
}

func (c *CancelledStatus3Choice) AddReason() *CancelledStatusReason6 {
	newValue := new(CancelledStatusReason6)
	c.Reason = append(c.Reason, newValue)
	return newValue
}
