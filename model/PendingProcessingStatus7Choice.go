package model

// Specifies whether the status is provided with a reason or not.
type PendingProcessingStatus7Choice struct {

	// Indicates that there is no reason available or to report.
	NoSpecifiedReason *NoReasonCode `xml:"NoSpcfdRsn"`

	// Specifies the reason of the PendingProcessingStatus.
	Reason []*PendingProcessingReason5 `xml:"Rsn"`
}

func (p *PendingProcessingStatus7Choice) SetNoSpecifiedReason(value string) {
	p.NoSpecifiedReason = (*NoReasonCode)(&value)
}

func (p *PendingProcessingStatus7Choice) AddReason() *PendingProcessingReason5 {
	newValue := new(PendingProcessingReason5)
	p.Reason = append(p.Reason, newValue)
	return newValue
}
