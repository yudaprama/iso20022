package model

// Status is accepted under certain conditions.
type ConditionallyAcceptedStatus2 struct {

	// Indicates that there is no reason available or to report.
	NoSpecifiedReason *NoReasonCode `xml:"NoSpcfdRsn"`

	// Reason for the conditionally accepted status.
	ReasonDetails []*ConditionallyAcceptedStatusReason2 `xml:"RsnDtls"`
}

func (c *ConditionallyAcceptedStatus2) SetNoSpecifiedReason(value string) {
	c.NoSpecifiedReason = (*NoReasonCode)(&value)
}

func (c *ConditionallyAcceptedStatus2) AddReasonDetails() *ConditionallyAcceptedStatusReason2 {
	newValue := new(ConditionallyAcceptedStatusReason2)
	c.ReasonDetails = append(c.ReasonDetails, newValue)
	return newValue
}
