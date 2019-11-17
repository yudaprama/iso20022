package model

// Choice between a reason or no reason for the corporate action instruction processing pending status.
type PendingStatus42Choice struct {

	// Reason not specified.
	NoSpecifiedReason *NoReasonCode `xml:"NoSpcfdRsn"`

	// Reason for the pending status.
	Reason []*PendingStatusReason10 `xml:"Rsn"`
}

func (p *PendingStatus42Choice) SetNoSpecifiedReason(value string) {
	p.NoSpecifiedReason = (*NoReasonCode)(&value)
}

func (p *PendingStatus42Choice) AddReason() *PendingStatusReason10 {
	newValue := new(PendingStatusReason10)
	p.Reason = append(p.Reason, newValue)
	return newValue
}
