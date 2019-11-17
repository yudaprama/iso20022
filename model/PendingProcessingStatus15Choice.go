package model

// Specifies whether the status is provided with a reason or not.
type PendingProcessingStatus15Choice struct {

	// Indicates that there is no reason available or to report.
	NoSpecifiedReason *NoReasonCode `xml:"NoSpcfdRsn"`

	// Specifies the reason of the pending processing status.
	Reason []*PendingProcessingReason12 `xml:"Rsn"`
}

func (p *PendingProcessingStatus15Choice) SetNoSpecifiedReason(value string) {
	p.NoSpecifiedReason = (*NoReasonCode)(&value)
}

func (p *PendingProcessingStatus15Choice) AddReason() *PendingProcessingReason12 {
	newValue := new(PendingProcessingReason12)
	p.Reason = append(p.Reason, newValue)
	return newValue
}
