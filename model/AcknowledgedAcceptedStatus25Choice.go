package model

// Specifies whether the status is provided with a reason or not.
type AcknowledgedAcceptedStatus25Choice struct {

	// Indicates that there is no reason available or to report.
	NoSpecifiedReason *NoReasonCode `xml:"NoSpcfdRsn"`

	// Specifies the reason of the acknowledged accepted status.
	Reason []*AcknowledgementReason13 `xml:"Rsn"`
}

func (a *AcknowledgedAcceptedStatus25Choice) SetNoSpecifiedReason(value string) {
	a.NoSpecifiedReason = (*NoReasonCode)(&value)
}

func (a *AcknowledgedAcceptedStatus25Choice) AddReason() *AcknowledgementReason13 {
	newValue := new(AcknowledgementReason13)
	a.Reason = append(a.Reason, newValue)
	return newValue
}
