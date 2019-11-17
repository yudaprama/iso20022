package model

// Specifies whether the status is provided with a reason or not.
type PendingStatus12Choice struct {

	// Indicates that there is no reason available or to report.
	NoSpecifiedReason *NoReasonCode `xml:"NoSpcfdRsn"`

	// Specifies the reason of the PendingStatus.
	Reason []*PendingReason8 `xml:"Rsn"`
}

func (p *PendingStatus12Choice) SetNoSpecifiedReason(value string) {
	p.NoSpecifiedReason = (*NoReasonCode)(&value)
}

func (p *PendingStatus12Choice) AddReason() *PendingReason8 {
	newValue := new(PendingReason8)
	p.Reason = append(p.Reason, newValue)
	return newValue
}
