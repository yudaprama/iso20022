package model

// Specifies whether the status is provided with a reason or not.
type PendingStatus46Choice struct {

	// Indicates that there is no reason available or to report.
	NoSpecifiedReason *NoReasonCode `xml:"NoSpcfdRsn"`

	// Specifies the reason of the pending status.
	Reason []*PendingReason20 `xml:"Rsn"`
}

func (p *PendingStatus46Choice) SetNoSpecifiedReason(value string) {
	p.NoSpecifiedReason = (*NoReasonCode)(&value)
}

func (p *PendingStatus46Choice) AddReason() *PendingReason20 {
	newValue := new(PendingReason20)
	p.Reason = append(p.Reason, newValue)
	return newValue
}
