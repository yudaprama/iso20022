package model

// Choice between a reason or no reason for the corporate action instruction processing pending status.
type PendingStatus52Choice struct {

	// Reason not specified.
	NoSpecifiedReason *NoReasonCode `xml:"NoSpcfdRsn"`

	// Reason for the pending status.
	Reason []*PendingStatusReason13 `xml:"Rsn"`
}

func (p *PendingStatus52Choice) SetNoSpecifiedReason(value string) {
	p.NoSpecifiedReason = (*NoReasonCode)(&value)
}

func (p *PendingStatus52Choice) AddReason() *PendingStatusReason13 {
	newValue := new(PendingStatusReason13)
	p.Reason = append(p.Reason, newValue)
	return newValue
}
