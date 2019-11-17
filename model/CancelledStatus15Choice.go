package model

// Choice between a reason or no reason for the corporate action instruction or instruction cancellation cancelled/cancellation completed status.
type CancelledStatus15Choice struct {

	// Reason not specified.
	NoSpecifiedReason *NoReasonCode `xml:"NoSpcfdRsn"`

	// Reason for the cancelled status.
	Reason []*CancelledStatusReason14 `xml:"Rsn"`
}

func (c *CancelledStatus15Choice) SetNoSpecifiedReason(value string) {
	c.NoSpecifiedReason = (*NoReasonCode)(&value)
}

func (c *CancelledStatus15Choice) AddReason() *CancelledStatusReason14 {
	newValue := new(CancelledStatusReason14)
	c.Reason = append(c.Reason, newValue)
	return newValue
}
