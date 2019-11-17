package model

// Choice of formats for the reason of a pending status.
type PendingProcessingReason9Choice struct {

	// Indicates that there is no reason available or to report.
	NoSpecifiedReason *NoReasonCode `xml:"NoSpcfdRsn"`

	// Reason for the pending status.
	Reason []*PendingProcessingReason8Choice `xml:"Rsn"`
}

func (p *PendingProcessingReason9Choice) SetNoSpecifiedReason(value string) {
	p.NoSpecifiedReason = (*NoReasonCode)(&value)
}

func (p *PendingProcessingReason9Choice) AddReason() *PendingProcessingReason8Choice {
	newValue := new(PendingProcessingReason8Choice)
	p.Reason = append(p.Reason, newValue)
	return newValue
}
