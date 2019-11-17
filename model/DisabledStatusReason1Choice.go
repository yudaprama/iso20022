package model

// Choice of formats for a disabled status reason.
type DisabledStatusReason1Choice struct {

	// There is no reason available or to report for the disabled account status.
	NoSpecifiedReason *NoReasonCode `xml:"NoSpcfdRsn"`

	// Reason for the disabled account status.
	Reason []*DisabledStatusReason1 `xml:"Rsn"`
}

func (d *DisabledStatusReason1Choice) SetNoSpecifiedReason(value string) {
	d.NoSpecifiedReason = (*NoReasonCode)(&value)
}

func (d *DisabledStatusReason1Choice) AddReason() *DisabledStatusReason1 {
	newValue := new(DisabledStatusReason1)
	d.Reason = append(d.Reason, newValue)
	return newValue
}
