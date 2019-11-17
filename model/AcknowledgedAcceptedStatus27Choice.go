package model

// Specifies whether the status is provided with a reason or not.
type AcknowledgedAcceptedStatus27Choice struct {

	// Indicates that there is no reason available or to report.
	NoSpecifiedReason *NoReasonCode `xml:"NoSpcfdRsn"`

	// Specifies the reason of the acknowledged accepted status.
	Reason []*AcknowledgementReason15 `xml:"Rsn"`
}

func (a *AcknowledgedAcceptedStatus27Choice) SetNoSpecifiedReason(value string) {
	a.NoSpecifiedReason = (*NoReasonCode)(&value)
}

func (a *AcknowledgedAcceptedStatus27Choice) AddReason() *AcknowledgementReason15 {
	newValue := new(AcknowledgementReason15)
	a.Reason = append(a.Reason, newValue)
	return newValue
}
