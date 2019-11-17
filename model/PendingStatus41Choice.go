package model

// Choice between a reason or no reason for the corporate action event processing pending status.
type PendingStatus41Choice struct {

	// Reason not specified.
	NoSpecifiedReason *NoReasonCode `xml:"NoSpcfdRsn"`

	// Reason for the pending status.
	Reason []*PendingStatusReason9 `xml:"Rsn"`
}

func (p *PendingStatus41Choice) SetNoSpecifiedReason(value string) {
	p.NoSpecifiedReason = (*NoReasonCode)(&value)
}

func (p *PendingStatus41Choice) AddReason() *PendingStatusReason9 {
	newValue := new(PendingStatusReason9)
	p.Reason = append(p.Reason, newValue)
	return newValue
}
