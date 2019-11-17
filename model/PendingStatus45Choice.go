package model

// Specifies whether the status is provided with a reason or not.
type PendingStatus45Choice struct {

	// Indicates that there is no reason available or to report.
	NoSpecifiedReason *NoReasonCode `xml:"NoSpcfdRsn"`

	// Specifies the reason of the pending status.
	Reason []*PendingReason19 `xml:"Rsn"`
}

func (p *PendingStatus45Choice) SetNoSpecifiedReason(value string) {
	p.NoSpecifiedReason = (*NoReasonCode)(&value)
}

func (p *PendingStatus45Choice) AddReason() *PendingReason19 {
	newValue := new(PendingReason19)
	p.Reason = append(p.Reason, newValue)
	return newValue
}
