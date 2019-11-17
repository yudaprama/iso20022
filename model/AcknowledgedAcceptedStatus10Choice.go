package model

// Specifies whether the status is provided with a reason or not.
type AcknowledgedAcceptedStatus10Choice struct {

	// Indicates that there is no reason available or to report.
	NoSpecifiedReason *NoReasonCode `xml:"NoSpcfdRsn"`

	// Specifies the reason of the AcknowledgedAcceptedStatus.
	Reason []*AcknowledgementReason2 `xml:"Rsn"`
}

func (a *AcknowledgedAcceptedStatus10Choice) SetNoSpecifiedReason(value string) {
	a.NoSpecifiedReason = (*NoReasonCode)(&value)
}

func (a *AcknowledgedAcceptedStatus10Choice) AddReason() *AcknowledgementReason2 {
	newValue := new(AcknowledgementReason2)
	a.Reason = append(a.Reason, newValue)
	return newValue
}
