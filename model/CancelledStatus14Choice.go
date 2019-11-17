package model

// Choice between a reason or no reason for the corporate action instruction or instruction cancellation cancelled/cancellation completed status.
type CancelledStatus14Choice struct {

	// Reason not specified.
	NoSpecifiedReason *NoReasonCode `xml:"NoSpcfdRsn"`

	// Reason for the cancelled status.
	Reason []*CancelledStatusReason13 `xml:"Rsn"`
}

func (c *CancelledStatus14Choice) SetNoSpecifiedReason(value string) {
	c.NoSpecifiedReason = (*NoReasonCode)(&value)
}

func (c *CancelledStatus14Choice) AddReason() *CancelledStatusReason13 {
	newValue := new(CancelledStatusReason13)
	c.Reason = append(c.Reason, newValue)
	return newValue
}
