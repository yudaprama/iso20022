package model

// Specifies whether the status is provided with a reason or not.
type PendingStatus37Choice struct {

	// Indicates that there is no reason available or to report.
	NoSpecifiedReason *NoReasonCode `xml:"NoSpcfdRsn"`

	// Specifies the reason of the pending status.
	Reason []*PendingReason15 `xml:"Rsn"`
}

func (p *PendingStatus37Choice) SetNoSpecifiedReason(value string) {
	p.NoSpecifiedReason = (*NoReasonCode)(&value)
}

func (p *PendingStatus37Choice) AddReason() *PendingReason15 {
	newValue := new(PendingReason15)
	p.Reason = append(p.Reason, newValue)
	return newValue
}
