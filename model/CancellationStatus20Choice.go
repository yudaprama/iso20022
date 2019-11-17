package model

// Specifies whether the status is provided with a reason or not.
type CancellationStatus20Choice struct {

	// Indicates that there is no reason available or to report.
	NoSpecifiedReason *NoReasonCode `xml:"NoSpcfdRsn"`

	// Specifies the reason of the cancellation status.
	Reason []*CancellationReason18 `xml:"Rsn"`
}

func (c *CancellationStatus20Choice) SetNoSpecifiedReason(value string) {
	c.NoSpecifiedReason = (*NoReasonCode)(&value)
}

func (c *CancellationStatus20Choice) AddReason() *CancellationReason18 {
	newValue := new(CancellationReason18)
	c.Reason = append(c.Reason, newValue)
	return newValue
}
