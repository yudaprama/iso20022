package model

// Specifies whether the status is provided with a reason or not.
type CancellationStatus3Choice struct {

	// Indicates that there is no reason available or to report.
	NoSpecifiedReason *NoReasonCode `xml:"NoSpcfdRsn"`

	// Specifies the reason of the CancellationStatus.
	Reason []*CancellationReason2 `xml:"Rsn,omitempty"`
}

func (c *CancellationStatus3Choice) SetNoSpecifiedReason(value string) {
	c.NoSpecifiedReason = (*NoReasonCode)(&value)
}

func (c *CancellationStatus3Choice) AddReason() *CancellationReason2 {
	newValue := new(CancellationReason2)
	c.Reason = append(c.Reason, newValue)
	return newValue
}
