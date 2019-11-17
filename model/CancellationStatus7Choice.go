package model

// Specifies whether the status is provided with a reason or not.
type CancellationStatus7Choice struct {

	// Indicates that there is no reason available or to report.
	NoSpecifiedReason *NoReasonCode `xml:"NoSpcfdRsn"`

	// Specifies the reason of the CancellationStatus.
	Reason []*CancellationReason5 `xml:"Rsn"`
}

func (c *CancellationStatus7Choice) SetNoSpecifiedReason(value string) {
	c.NoSpecifiedReason = (*NoReasonCode)(&value)
}

func (c *CancellationStatus7Choice) AddReason() *CancellationReason5 {
	newValue := new(CancellationReason5)
	c.Reason = append(c.Reason, newValue)
	return newValue
}
