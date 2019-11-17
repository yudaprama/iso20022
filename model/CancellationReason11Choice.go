package model

// Choice of format for the cancellation processing status.
type CancellationReason11Choice struct {

	// Indicates that there is no reason available or to report.
	NoSpecifiedReason *NoReasonCode `xml:"NoSpcfdRsn"`

	// Specifies the reason of the CancellationPending Status.
	Reason []*AwaitingCancellationReason1 `xml:"Rsn"`
}

func (c *CancellationReason11Choice) SetNoSpecifiedReason(value string) {
	c.NoSpecifiedReason = (*NoReasonCode)(&value)
}

func (c *CancellationReason11Choice) AddReason() *AwaitingCancellationReason1 {
	newValue := new(AwaitingCancellationReason1)
	c.Reason = append(c.Reason, newValue)
	return newValue
}
