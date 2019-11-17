package model

// Choice between a reason or no reason for the corporate action instruction or instruction cancellation cancelled/cancellation completed status.
type CancelledStatus11Choice struct {

	// Reason not specified.
	NoSpecifiedReason *NoReasonCode `xml:"NoSpcfdRsn"`

	// Reason for the cancelled status.
	Reason []*CancelledStatusReason12 `xml:"Rsn"`
}

func (c *CancelledStatus11Choice) SetNoSpecifiedReason(value string) {
	c.NoSpecifiedReason = (*NoReasonCode)(&value)
}

func (c *CancelledStatus11Choice) AddReason() *CancelledStatusReason12 {
	newValue := new(CancelledStatusReason12)
	c.Reason = append(c.Reason, newValue)
	return newValue
}
