package model

// Specifies whether the status is provided with a reason or not.
type AcknowledgedAcceptedStatus22Choice struct {

	// Indicates that there is no reason available or to report.
	NoSpecifiedReason *NoReasonCode `xml:"NoSpcfdRsn"`

	// Specifies the reason of the acknowledged accepted status.
	Reason []*AcknowledgementReason10 `xml:"Rsn"`
}

func (a *AcknowledgedAcceptedStatus22Choice) SetNoSpecifiedReason(value string) {
	a.NoSpecifiedReason = (*NoReasonCode)(&value)
}

func (a *AcknowledgedAcceptedStatus22Choice) AddReason() *AcknowledgementReason10 {
	newValue := new(AcknowledgementReason10)
	a.Reason = append(a.Reason, newValue)
	return newValue
}
