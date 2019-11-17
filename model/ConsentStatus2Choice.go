package model

// Choice of consent status.
type ConsentStatus2Choice struct {

	// Indicates that there is no reason available or to report.
	NoSpecifiedReason *NoReasonCode `xml:"NoSpcfdRsn"`

	// Reason provided for the status.
	Reason []*ConsentReason2 `xml:"Rsn"`
}

func (c *ConsentStatus2Choice) SetNoSpecifiedReason(value string) {
	c.NoSpecifiedReason = (*NoReasonCode)(&value)
}

func (c *ConsentStatus2Choice) AddReason() *ConsentReason2 {
	newValue := new(ConsentReason2)
	c.Reason = append(c.Reason, newValue)
	return newValue
}
