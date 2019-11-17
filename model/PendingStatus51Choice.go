package model

// Specifies whether the status is provided with a reason or not.
type PendingStatus51Choice struct {

	// Indicates that there is no reason available or to report.
	NoSpecifiedReason *NoReasonCode `xml:"NoSpcfdRsn"`

	// Specifies the reason of the pending status.
	Reason []*PendingReason25 `xml:"Rsn"`
}

func (p *PendingStatus51Choice) SetNoSpecifiedReason(value string) {
	p.NoSpecifiedReason = (*NoReasonCode)(&value)
}

func (p *PendingStatus51Choice) AddReason() *PendingReason25 {
	newValue := new(PendingReason25)
	p.Reason = append(p.Reason, newValue)
	return newValue
}
