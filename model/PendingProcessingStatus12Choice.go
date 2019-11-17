package model

// Specifies whether the status is provided with a reason or not.
type PendingProcessingStatus12Choice struct {

	// Indicates that there is no reason available or to report.
	NoSpecifiedReason *NoReasonCode `xml:"NoSpcfdRsn"`

	// Specifies the reason of the pending processing status.
	Reason []*PendingProcessingReason9 `xml:"Rsn"`
}

func (p *PendingProcessingStatus12Choice) SetNoSpecifiedReason(value string) {
	p.NoSpecifiedReason = (*NoReasonCode)(&value)
}

func (p *PendingProcessingStatus12Choice) AddReason() *PendingProcessingReason9 {
	newValue := new(PendingProcessingReason9)
	p.Reason = append(p.Reason, newValue)
	return newValue
}
