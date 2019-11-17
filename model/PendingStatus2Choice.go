package model

// Choice between a reason or no reason for the corporate action event processing pending status.
type PendingStatus2Choice struct {

	// Reason not specified.
	NoSpecifiedReason *NoReasonCode `xml:"NoSpcfdRsn"`

	// Reason for the pending status.
	Reason []*PendingStatusReason2 `xml:"Rsn"`
}

func (p *PendingStatus2Choice) SetNoSpecifiedReason(value string) {
	p.NoSpecifiedReason = (*NoReasonCode)(&value)
}

func (p *PendingStatus2Choice) AddReason() *PendingStatusReason2 {
	newValue := new(PendingStatusReason2)
	p.Reason = append(p.Reason, newValue)
	return newValue
}
