package model

// Reason for a conditionally accepted status.
type ConditionallyAcceptedStatus3Choice struct {

	// No reason available or to report for the conditionally accepted status.
	NoSpecifiedReason *NoReasonCode `xml:"NoSpcfdRsn"`

	// Reason for the conditionally accepted status, expressed as a code.
	ReasonDetails []*ConditionallyAcceptedStatusReason3 `xml:"RsnDtls"`
}

func (c *ConditionallyAcceptedStatus3Choice) SetNoSpecifiedReason(value string) {
	c.NoSpecifiedReason = (*NoReasonCode)(&value)
}

func (c *ConditionallyAcceptedStatus3Choice) AddReasonDetails() *ConditionallyAcceptedStatusReason3 {
	newValue := new(ConditionallyAcceptedStatusReason3)
	c.ReasonDetails = append(c.ReasonDetails, newValue)
	return newValue
}
