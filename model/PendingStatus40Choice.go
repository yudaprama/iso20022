package model

// Specifies whether the status is provided with a reason or not.
type PendingStatus40Choice struct {

	// Indicates that there is no reason available or to report.
	NoSpecifiedReason *NoReasonCode `xml:"NoSpcfdRsn"`

	// Specifies the reason of the pending status.
	Reason []*PendingReason18 `xml:"Rsn"`
}

func (p *PendingStatus40Choice) SetNoSpecifiedReason(value string) {
	p.NoSpecifiedReason = (*NoReasonCode)(&value)
}

func (p *PendingStatus40Choice) AddReason() *PendingReason18 {
	newValue := new(PendingReason18)
	p.Reason = append(p.Reason, newValue)
	return newValue
}
