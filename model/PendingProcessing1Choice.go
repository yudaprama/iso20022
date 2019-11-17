package model

// Choice of format for the Pending Processing status.
type PendingProcessing1Choice struct {

	// Indicates that there is no reason available or to report.
	NoSpecifiedReason *NoReasonCode `xml:"NoSpcfdRsn"`

	// Specifies the reason of the Pending Processing Status.
	Reason []*AwaitingAffirmationReason1 `xml:"Rsn"`
}

func (p *PendingProcessing1Choice) SetNoSpecifiedReason(value string) {
	p.NoSpecifiedReason = (*NoReasonCode)(&value)
}

func (p *PendingProcessing1Choice) AddReason() *AwaitingAffirmationReason1 {
	newValue := new(AwaitingAffirmationReason1)
	p.Reason = append(p.Reason, newValue)
	return newValue
}
