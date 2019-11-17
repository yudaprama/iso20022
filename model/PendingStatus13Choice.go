package model

// Specifies whether the status is provided with a reason or not.
type PendingStatus13Choice struct {

	// Indicates that there is no reason available or to report.
	NoSpecifiedReason *NoReasonCode `xml:"NoSpcfdRsn"`

	// Specifies the reason of the PendingStatus.
	Reason []*PendingReason2 `xml:"Rsn"`
}

func (p *PendingStatus13Choice) SetNoSpecifiedReason(value string) {
	p.NoSpecifiedReason = (*NoReasonCode)(&value)
}

func (p *PendingStatus13Choice) AddReason() *PendingReason2 {
	newValue := new(PendingReason2)
	p.Reason = append(p.Reason, newValue)
	return newValue
}
