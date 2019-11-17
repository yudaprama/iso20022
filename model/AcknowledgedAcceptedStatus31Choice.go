package model

// Specifies whether the status is provided with a reason or not.
type AcknowledgedAcceptedStatus31Choice struct {

	// Indicates that there is no reason available or to report.
	NoSpecifiedReason *NoReasonCode `xml:"NoSpcfdRsn"`

	// Specifies the reason of the acknowledged accepted status.
	Reason []*AcknowledgementReason19 `xml:"Rsn"`
}

func (a *AcknowledgedAcceptedStatus31Choice) SetNoSpecifiedReason(value string) {
	a.NoSpecifiedReason = (*NoReasonCode)(&value)
}

func (a *AcknowledgedAcceptedStatus31Choice) AddReason() *AcknowledgementReason19 {
	newValue := new(AcknowledgementReason19)
	a.Reason = append(a.Reason, newValue)
	return newValue
}
