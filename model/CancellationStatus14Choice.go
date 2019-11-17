package model

// Specifies whether the status is provided with a reason or not.
type CancellationStatus14Choice struct {

	// Indicates that there is no reason available or to report.
	NoSpecifiedReason *NoReasonCode `xml:"NoSpcfdRsn"`

	// Specifies the reason of the cancellation status.
	Reason []*CancellationReason9 `xml:"Rsn"`
}

func (c *CancellationStatus14Choice) SetNoSpecifiedReason(value string) {
	c.NoSpecifiedReason = (*NoReasonCode)(&value)
}

func (c *CancellationStatus14Choice) AddReason() *CancellationReason9 {
	newValue := new(CancellationReason9)
	c.Reason = append(c.Reason, newValue)
	return newValue
}
