package model

// Specifies whether the status is provided with a reason or not.
type CancellationStatus8Choice struct {

	// Indicates that there is no reason available or to report.
	NoSpecifiedReason *NoReasonCode `xml:"NoSpcfdRsn"`

	// Specifies the reason of the CancellationStatus.
	Reason []*CancellationReason1 `xml:"Rsn"`
}

func (c *CancellationStatus8Choice) SetNoSpecifiedReason(value string) {
	c.NoSpecifiedReason = (*NoReasonCode)(&value)
}

func (c *CancellationStatus8Choice) AddReason() *CancellationReason1 {
	newValue := new(CancellationReason1)
	c.Reason = append(c.Reason, newValue)
	return newValue
}
