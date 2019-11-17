package model

// Specifies whether the status is provided with a reason or not.
type CancelledStatus10Choice struct {

	// Indicates that there is no reason available or to report.
	NoSpecifiedReason *NoReasonCode `xml:"NoSpcfdRsn"`

	// Reason for the cancellation status.
	Reason *CancellationReason11 `xml:"Rsn"`
}

func (c *CancelledStatus10Choice) SetNoSpecifiedReason(value string) {
	c.NoSpecifiedReason = (*NoReasonCode)(&value)
}

func (c *CancelledStatus10Choice) AddReason() *CancellationReason11 {
	c.Reason = new(CancellationReason11)
	return c.Reason
}
