package model

// Choice of formats for a closure pending status reason.
type ClosurePendingStatusReason1Choice struct {

	// There is no reason available or to report for the closure pending status.
	NoSpecifiedReason *NoReasonCode `xml:"NoSpcfdRsn"`

	// Reason for the closure pending status.
	Reason []*ClosurePendingStatusReason1 `xml:"Rsn"`
}

func (c *ClosurePendingStatusReason1Choice) SetNoSpecifiedReason(value string) {
	c.NoSpecifiedReason = (*NoReasonCode)(&value)
}

func (c *ClosurePendingStatusReason1Choice) AddReason() *ClosurePendingStatusReason1 {
	newValue := new(ClosurePendingStatusReason1)
	c.Reason = append(c.Reason, newValue)
	return newValue
}
