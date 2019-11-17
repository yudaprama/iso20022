package model

// Specifies whether the status is provided with a reason or not.
type PendingProcessingStatus16Choice struct {

	// Indicates that there is no reason available or to report.
	NoSpecifiedReason *NoReasonCode `xml:"NoSpcfdRsn"`

	// Specifies the reason of the pending processing status.
	Reason []*PendingProcessingReason13 `xml:"Rsn"`
}

func (p *PendingProcessingStatus16Choice) SetNoSpecifiedReason(value string) {
	p.NoSpecifiedReason = (*NoReasonCode)(&value)
}

func (p *PendingProcessingStatus16Choice) AddReason() *PendingProcessingReason13 {
	newValue := new(PendingProcessingReason13)
	p.Reason = append(p.Reason, newValue)
	return newValue
}
