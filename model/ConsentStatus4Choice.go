package model

// Choice of consent status.
type ConsentStatus4Choice struct {

	// Indicates that there is no reason available or to report.
	NoSpecifiedReason *NoReasonCode `xml:"NoSpcfdRsn"`

	// Reason provided for the status.
	Reason []*ConsentReason4 `xml:"Rsn"`
}

func (c *ConsentStatus4Choice) SetNoSpecifiedReason(value string) {
	c.NoSpecifiedReason = (*NoReasonCode)(&value)
}

func (c *ConsentStatus4Choice) AddReason() *ConsentReason4 {
	newValue := new(ConsentReason4)
	c.Reason = append(c.Reason, newValue)
	return newValue
}
