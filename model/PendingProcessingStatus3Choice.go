package model

// Specifies whether the status is provided with a reason or not.
type PendingProcessingStatus3Choice struct {

	// Indicates that there is no reason available or to report.
	NoSpecifiedReason *NoReasonCode `xml:"NoSpcfdRsn"`

	// Specifies the reason of the PendingProcessingStatus.
	Reason []*PendingProcessingReason3 `xml:"Rsn"`
}

func (p *PendingProcessingStatus3Choice) SetNoSpecifiedReason(value string) {
	p.NoSpecifiedReason = (*NoReasonCode)(&value)
}

func (p *PendingProcessingStatus3Choice) AddReason() *PendingProcessingReason3 {
	newValue := new(PendingProcessingReason3)
	p.Reason = append(p.Reason, newValue)
	return newValue
}
