package model

// Choice of consent status.
type ConsentStatus5Choice struct {

	// Indicates that there is no reason available or to report.
	NoSpecifiedReason *NoReasonCode `xml:"NoSpcfdRsn"`

	// Reason provided for the status.
	Reason []*ConsentReason5 `xml:"Rsn"`
}

func (c *ConsentStatus5Choice) SetNoSpecifiedReason(value string) {
	c.NoSpecifiedReason = (*NoReasonCode)(&value)
}

func (c *ConsentStatus5Choice) AddReason() *ConsentReason5 {
	newValue := new(ConsentReason5)
	c.Reason = append(c.Reason, newValue)
	return newValue
}
