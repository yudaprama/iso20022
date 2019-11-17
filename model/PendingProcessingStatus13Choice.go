package model

// Specifies whether the status is provided with a reason or not.
type PendingProcessingStatus13Choice struct {

	// Indicates that there is no reason available or to report.
	NoSpecifiedReason *NoReasonCode `xml:"NoSpcfdRsn"`

	// Specifies the reason of the pending processing status.
	Reason []*PendingProcessingReason10 `xml:"Rsn"`
}

func (p *PendingProcessingStatus13Choice) SetNoSpecifiedReason(value string) {
	p.NoSpecifiedReason = (*NoReasonCode)(&value)
}

func (p *PendingProcessingStatus13Choice) AddReason() *PendingProcessingReason10 {
	newValue := new(PendingProcessingReason10)
	p.Reason = append(p.Reason, newValue)
	return newValue
}
