package model

// Specifies whether the status is provided with a reason or not.
type PendingStatus9Choice struct {

	// Indicates that there is no reason available or to report.
	NoSpecifiedReason *NoReasonCode `xml:"NoSpcfdRsn"`

	// Specifies the reason of the PendingStatus.
	Reason []*PendingReason5 `xml:"Rsn"`
}

func (p *PendingStatus9Choice) SetNoSpecifiedReason(value string) {
	p.NoSpecifiedReason = (*NoReasonCode)(&value)
}

func (p *PendingStatus9Choice) AddReason() *PendingReason5 {
	newValue := new(PendingReason5)
	p.Reason = append(p.Reason, newValue)
	return newValue
}
