package model

// Choice of formats for a pending status reason.
type PendingStatusReason1Choice struct {

	// There is no reason available or to report for the pending account status.
	NoSpecifiedReason *NoReasonCode `xml:"NoSpcfdRsn"`

	// Reason for the pending account status.
	Reason []*PendingStatusReason14 `xml:"Rsn"`
}

func (p *PendingStatusReason1Choice) SetNoSpecifiedReason(value string) {
	p.NoSpecifiedReason = (*NoReasonCode)(&value)
}

func (p *PendingStatusReason1Choice) AddReason() *PendingStatusReason14 {
	newValue := new(PendingStatusReason14)
	p.Reason = append(p.Reason, newValue)
	return newValue
}
