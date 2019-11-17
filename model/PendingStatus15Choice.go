package model

// Specifies whether the status is provided with a reason or not.
type PendingStatus15Choice struct {

	// Indicates that there is no reason available or to report.
	NoSpecifiedReason *NoReasonCode `xml:"NoSpcfdRsn"`

	// Specifies the reason of the PendingStatus.
	Reason []*PendingReason1 `xml:"Rsn"`
}

func (p *PendingStatus15Choice) SetNoSpecifiedReason(value string) {
	p.NoSpecifiedReason = (*NoReasonCode)(&value)
}

func (p *PendingStatus15Choice) AddReason() *PendingReason1 {
	newValue := new(PendingReason1)
	p.Reason = append(p.Reason, newValue)
	return newValue
}
