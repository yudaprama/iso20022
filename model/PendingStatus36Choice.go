package model

// Specifies whether the status is provided with a reason or not.
type PendingStatus36Choice struct {

	// Indicates that there is no reason available or to report.
	NoSpecifiedReason *NoReasonCode `xml:"NoSpcfdRsn"`

	// Specifies the reason of the pending status.
	Reason []*PendingReason14 `xml:"Rsn"`
}

func (p *PendingStatus36Choice) SetNoSpecifiedReason(value string) {
	p.NoSpecifiedReason = (*NoReasonCode)(&value)
}

func (p *PendingStatus36Choice) AddReason() *PendingReason14 {
	newValue := new(PendingReason14)
	p.Reason = append(p.Reason, newValue)
	return newValue
}
