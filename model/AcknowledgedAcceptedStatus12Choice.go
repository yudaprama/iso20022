package model

// Specifies whether the status is provided with a reason or not.
type AcknowledgedAcceptedStatus12Choice struct {

	// Indicates that there is no reason available or to report.
	NoSpecifiedReason *NoReasonCode `xml:"NoSpcfdRsn"`

	// Specifies the reason of the AcknowledgedAcceptedStatus.
	Reason []*AcknowledgementReason1 `xml:"Rsn"`
}

func (a *AcknowledgedAcceptedStatus12Choice) SetNoSpecifiedReason(value string) {
	a.NoSpecifiedReason = (*NoReasonCode)(&value)
}

func (a *AcknowledgedAcceptedStatus12Choice) AddReason() *AcknowledgementReason1 {
	newValue := new(AcknowledgementReason1)
	a.Reason = append(a.Reason, newValue)
	return newValue
}
