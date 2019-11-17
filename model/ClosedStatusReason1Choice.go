package model

// Choice of formats for a closed status reason.
type ClosedStatusReason1Choice struct {

	// There is no reason available or to report for the closed account status.
	NoSpecifiedReason *NoReasonCode `xml:"NoSpcfdRsn"`

	// Reason for the closed account status.
	Reason []*ClosedStatusReason1 `xml:"Rsn"`
}

func (c *ClosedStatusReason1Choice) SetNoSpecifiedReason(value string) {
	c.NoSpecifiedReason = (*NoReasonCode)(&value)
}

func (c *ClosedStatusReason1Choice) AddReason() *ClosedStatusReason1 {
	newValue := new(ClosedStatusReason1)
	c.Reason = append(c.Reason, newValue)
	return newValue
}
