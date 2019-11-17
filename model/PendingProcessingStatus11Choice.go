package model

// Specifies whether the status is provided with a reason or not.
type PendingProcessingStatus11Choice struct {

	// Indicates that there is no reason available or to report.
	NoSpecifiedReason *NoReasonCode `xml:"NoSpcfdRsn"`

	// Specifies the reason of the pending processing status.
	Reason []*PendingProcessingReason8 `xml:"Rsn"`
}

func (p *PendingProcessingStatus11Choice) SetNoSpecifiedReason(value string) {
	p.NoSpecifiedReason = (*NoReasonCode)(&value)
}

func (p *PendingProcessingStatus11Choice) AddReason() *PendingProcessingReason8 {
	newValue := new(PendingProcessingReason8)
	p.Reason = append(p.Reason, newValue)
	return newValue
}
