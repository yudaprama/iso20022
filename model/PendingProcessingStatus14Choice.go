package model

// Specifies whether the status is provided with a reason or not.
type PendingProcessingStatus14Choice struct {

	// Indicates that there is no reason available or to report.
	NoSpecifiedReason *NoReasonCode `xml:"NoSpcfdRsn"`

	// Specifies the reason of the pending processing status.
	Reason []*PendingProcessingReason11 `xml:"Rsn"`
}

func (p *PendingProcessingStatus14Choice) SetNoSpecifiedReason(value string) {
	p.NoSpecifiedReason = (*NoReasonCode)(&value)
}

func (p *PendingProcessingStatus14Choice) AddReason() *PendingProcessingReason11 {
	newValue := new(PendingProcessingReason11)
	p.Reason = append(p.Reason, newValue)
	return newValue
}
