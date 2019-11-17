package model

// Specifies whether the status is provided with a reason or not.
type CancelledStatus16Choice struct {

	// Indicates that there is no reason available or to report.
	NoSpecifiedReason *NoReasonCode `xml:"NoSpcfdRsn"`

	// Reason for the cancellation status.
	Reason *CancellationReason16 `xml:"Rsn"`
}

func (c *CancelledStatus16Choice) SetNoSpecifiedReason(value string) {
	c.NoSpecifiedReason = (*NoReasonCode)(&value)
}

func (c *CancelledStatus16Choice) AddReason() *CancellationReason16 {
	c.Reason = new(CancellationReason16)
	return c.Reason
}
