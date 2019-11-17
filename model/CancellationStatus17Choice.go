package model

// Specifies whether the status is provided with a reason or not.
type CancellationStatus17Choice struct {

	// Indicates that there is no reason available or to report.
	NoSpecifiedReason *NoReasonCode `xml:"NoSpcfdRsn"`

	// Specifies the reason of the cancellation status.
	Reason []*CancellationReason14 `xml:"Rsn"`
}

func (c *CancellationStatus17Choice) SetNoSpecifiedReason(value string) {
	c.NoSpecifiedReason = (*NoReasonCode)(&value)
}

func (c *CancellationStatus17Choice) AddReason() *CancellationReason14 {
	newValue := new(CancellationReason14)
	c.Reason = append(c.Reason, newValue)
	return newValue
}
