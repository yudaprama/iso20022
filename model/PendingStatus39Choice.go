package model

// Specifies whether the status is provided with a reason or not.
type PendingStatus39Choice struct {

	// Indicates that there is no reason available or to report.
	NoSpecifiedReason *NoReasonCode `xml:"NoSpcfdRsn"`

	// Specifies the reason of the pending status.
	Reason []*PendingReason17 `xml:"Rsn"`
}

func (p *PendingStatus39Choice) SetNoSpecifiedReason(value string) {
	p.NoSpecifiedReason = (*NoReasonCode)(&value)
}

func (p *PendingStatus39Choice) AddReason() *PendingReason17 {
	newValue := new(PendingReason17)
	p.Reason = append(p.Reason, newValue)
	return newValue
}
