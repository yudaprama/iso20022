package model

// Specifies whether the status is provided with a reason or not.
type CancelledStatus5Choice struct {

	// Indicates that there is no reason available or to report.
	NoSpecifiedReason *NoReasonCode `xml:"NoSpcfdRsn"`

	// Reason for the cancellation status.
	Reason *CancellationReason7 `xml:"Rsn"`
}

func (c *CancelledStatus5Choice) SetNoSpecifiedReason(value string) {
	c.NoSpecifiedReason = (*NoReasonCode)(&value)
}

func (c *CancelledStatus5Choice) AddReason() *CancellationReason7 {
	c.Reason = new(CancellationReason7)
	return c.Reason
}
