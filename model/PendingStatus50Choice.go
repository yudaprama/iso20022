package model

// Specifies whether the status is provided with a reason or not.
type PendingStatus50Choice struct {

	// Indicates that there is no reason available or to report.
	NoSpecifiedReason *NoReasonCode `xml:"NoSpcfdRsn"`

	// Specifies the reason of the pending status.
	Reason []*PendingReason24 `xml:"Rsn"`
}

func (p *PendingStatus50Choice) SetNoSpecifiedReason(value string) {
	p.NoSpecifiedReason = (*NoReasonCode)(&value)
}

func (p *PendingStatus50Choice) AddReason() *PendingReason24 {
	newValue := new(PendingReason24)
	p.Reason = append(p.Reason, newValue)
	return newValue
}
