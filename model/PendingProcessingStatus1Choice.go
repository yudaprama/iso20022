package model

// Specifies whether the status is provided with a reason or not.
type PendingProcessingStatus1Choice struct {

	// Indicates that there is no reason available or to report.
	NoSpecifiedReason *NoReasonCode `xml:"NoSpcfdRsn"`

	// Specifies the reason of the PendingProcessingStatus.
	Reason []*PendingProcessingReason1 `xml:"Rsn,omitempty"`
}

func (p *PendingProcessingStatus1Choice) SetNoSpecifiedReason(value string) {
	p.NoSpecifiedReason = (*NoReasonCode)(&value)
}

func (p *PendingProcessingStatus1Choice) AddReason() *PendingProcessingReason1 {
	newValue := new(PendingProcessingReason1)
	p.Reason = append(p.Reason, newValue)
	return newValue
}
