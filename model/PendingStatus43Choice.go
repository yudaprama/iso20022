package model

// Choice between a reason or no reason for the corporate action event processing pending status.
type PendingStatus43Choice struct {

	// Reason not specified.
	NoSpecifiedReason *NoReasonCode `xml:"NoSpcfdRsn"`

	// Reason for the pending status.
	Reason []*PendingStatusReason11 `xml:"Rsn"`
}

func (p *PendingStatus43Choice) SetNoSpecifiedReason(value string) {
	p.NoSpecifiedReason = (*NoReasonCode)(&value)
}

func (p *PendingStatus43Choice) AddReason() *PendingStatusReason11 {
	newValue := new(PendingStatusReason11)
	p.Reason = append(p.Reason, newValue)
	return newValue
}
