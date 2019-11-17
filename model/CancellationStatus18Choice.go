package model

// Specifies whether the status is provided with a reason or not.
type CancellationStatus18Choice struct {

	// Indicates that there is no reason available or to report.
	NoSpecifiedReason *NoReasonCode `xml:"NoSpcfdRsn"`

	// Specifies the reason of the cancellation status.
	Reason []*CancellationReason15 `xml:"Rsn"`
}

func (c *CancellationStatus18Choice) SetNoSpecifiedReason(value string) {
	c.NoSpecifiedReason = (*NoReasonCode)(&value)
}

func (c *CancellationStatus18Choice) AddReason() *CancellationReason15 {
	newValue := new(CancellationReason15)
	c.Reason = append(c.Reason, newValue)
	return newValue
}
