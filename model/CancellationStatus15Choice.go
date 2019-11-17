package model

// Specifies whether the status is provided with a reason or not.
type CancellationStatus15Choice struct {

	// Indicates that there is no reason available or to report.
	NoSpecifiedReason *NoReasonCode `xml:"NoSpcfdRsn"`

	// Specifies the reason of the cancellation status.
	Reason []*CancellationReason10 `xml:"Rsn"`
}

func (c *CancellationStatus15Choice) SetNoSpecifiedReason(value string) {
	c.NoSpecifiedReason = (*NoReasonCode)(&value)
}

func (c *CancellationStatus15Choice) AddReason() *CancellationReason10 {
	newValue := new(CancellationReason10)
	c.Reason = append(c.Reason, newValue)
	return newValue
}
