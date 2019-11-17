package model

// Specifies whether the status is provided with a reason or not.
type PendingProcessingStatus4Choice struct {

	// Indicates that there is no reason available or to report.
	NoSpecifiedReason *NoReasonCode `xml:"NoSpcfdRsn"`

	// Specifies the reason of the PendingProcessingStatus.
	Reason []*PendingProcessingReason1 `xml:"Rsn"`
}

func (p *PendingProcessingStatus4Choice) SetNoSpecifiedReason(value string) {
	p.NoSpecifiedReason = (*NoReasonCode)(&value)
}

func (p *PendingProcessingStatus4Choice) AddReason() *PendingProcessingReason1 {
	newValue := new(PendingProcessingReason1)
	p.Reason = append(p.Reason, newValue)
	return newValue
}
