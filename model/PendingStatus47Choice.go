package model

// Specifies whether the status is provided with a reason or not.
type PendingStatus47Choice struct {

	// Indicates that there is no reason available or to report.
	NoSpecifiedReason *NoReasonCode `xml:"NoSpcfdRsn"`

	// Specifies the reason of the pending status.
	Reason []*PendingReason21 `xml:"Rsn"`
}

func (p *PendingStatus47Choice) SetNoSpecifiedReason(value string) {
	p.NoSpecifiedReason = (*NoReasonCode)(&value)
}

func (p *PendingStatus47Choice) AddReason() *PendingReason21 {
	newValue := new(PendingReason21)
	p.Reason = append(p.Reason, newValue)
	return newValue
}
