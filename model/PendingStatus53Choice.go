package model

// Choice between a reason or no reason for the corporate action instruction processing pending status.
type PendingStatus53Choice struct {

	// Reason not specified.
	NoSpecifiedReason *NoReasonCode `xml:"NoSpcfdRsn"`

	// Reason for the pending status.
	Reason []*PendingStatusReason15 `xml:"Rsn"`
}

func (p *PendingStatus53Choice) SetNoSpecifiedReason(value string) {
	p.NoSpecifiedReason = (*NoReasonCode)(&value)
}

func (p *PendingStatus53Choice) AddReason() *PendingStatusReason15 {
	newValue := new(PendingStatusReason15)
	p.Reason = append(p.Reason, newValue)
	return newValue
}
