package model

// Choice of formats for an enabled status reason.
type EnabledStatusReason1Choice struct {

	// There is no reason available or to report for the enabled account status.
	NoSpecifiedReason *NoReasonCode `xml:"NoSpcfdRsn"`

	// Reason for the enabled account status.
	Reason []*EnabledStatusReason1 `xml:"Rsn"`
}

func (e *EnabledStatusReason1Choice) SetNoSpecifiedReason(value string) {
	e.NoSpecifiedReason = (*NoReasonCode)(&value)
}

func (e *EnabledStatusReason1Choice) AddReason() *EnabledStatusReason1 {
	newValue := new(EnabledStatusReason1)
	e.Reason = append(e.Reason, newValue)
	return newValue
}
