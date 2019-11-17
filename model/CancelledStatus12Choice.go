package model

// Choice between a reason or no reason for the corporate action instruction or instruction cancellation cancelled/cancellation completed status.
type CancelledStatus12Choice struct {

	// Reason not specified.
	NoSpecifiedReason *NoReasonCode `xml:"NoSpcfdRsn"`

	// Reason for the cancelled status.
	Reason []*CancelledStatusReason11 `xml:"Rsn"`
}

func (c *CancelledStatus12Choice) SetNoSpecifiedReason(value string) {
	c.NoSpecifiedReason = (*NoReasonCode)(&value)
}

func (c *CancelledStatus12Choice) AddReason() *CancelledStatusReason11 {
	newValue := new(CancelledStatusReason11)
	c.Reason = append(c.Reason, newValue)
	return newValue
}
