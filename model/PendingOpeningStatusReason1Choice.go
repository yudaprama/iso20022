package model

// Choice of formats for a pending account opening status reason.
type PendingOpeningStatusReason1Choice struct {

	// There is no reason available or to report for the pending account opening status.
	NoSpecifiedReason *NoReasonCode `xml:"NoSpcfdRsn"`

	// Reason for the pending account opening status.
	Reason []*PendingOpeningStatusReason1 `xml:"Rsn"`
}

func (p *PendingOpeningStatusReason1Choice) SetNoSpecifiedReason(value string) {
	p.NoSpecifiedReason = (*NoReasonCode)(&value)
}

func (p *PendingOpeningStatusReason1Choice) AddReason() *PendingOpeningStatusReason1 {
	newValue := new(PendingOpeningStatusReason1)
	p.Reason = append(p.Reason, newValue)
	return newValue
}
