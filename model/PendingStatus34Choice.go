package model

// Choice between a reason or no reason for the corporate action instruction processing pending status.
type PendingStatus34Choice struct {

	// Reason not specified.
	NoSpecifiedReason *NoReasonCode `xml:"NoSpcfdRsn"`

	// Reason for the pending status.
	Reason []*PendingStatusReason7 `xml:"Rsn"`
}

func (p *PendingStatus34Choice) SetNoSpecifiedReason(value string) {
	p.NoSpecifiedReason = (*NoReasonCode)(&value)
}

func (p *PendingStatus34Choice) AddReason() *PendingStatusReason7 {
	newValue := new(PendingStatusReason7)
	p.Reason = append(p.Reason, newValue)
	return newValue
}
