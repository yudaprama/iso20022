package model

// Specifies whether the status is provided with a reason or not.
type PendingStatus38Choice struct {

	// Indicates that there is no reason available or to report.
	NoSpecifiedReason *NoReasonCode `xml:"NoSpcfdRsn"`

	// Specifies the reason of the pending status.
	Reason []*PendingReason16 `xml:"Rsn"`
}

func (p *PendingStatus38Choice) SetNoSpecifiedReason(value string) {
	p.NoSpecifiedReason = (*NoReasonCode)(&value)
}

func (p *PendingStatus38Choice) AddReason() *PendingReason16 {
	newValue := new(PendingReason16)
	p.Reason = append(p.Reason, newValue)
	return newValue
}
