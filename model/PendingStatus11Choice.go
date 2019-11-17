package model

// Specifies whether the status is provided with a reason or not.
type PendingStatus11Choice struct {

	// Indicates that there is no reason available or to report.
	NoSpecifiedReason *NoReasonCode `xml:"NoSpcfdRsn"`

	// Specifies the reason of the PendingStatus.
	Reason []*PendingReason7 `xml:"Rsn"`
}

func (p *PendingStatus11Choice) SetNoSpecifiedReason(value string) {
	p.NoSpecifiedReason = (*NoReasonCode)(&value)
}

func (p *PendingStatus11Choice) AddReason() *PendingReason7 {
	newValue := new(PendingReason7)
	p.Reason = append(p.Reason, newValue)
	return newValue
}
