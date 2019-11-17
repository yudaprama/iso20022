package model

// Specifies whether the status is provided with a reason or not.
type AcknowledgedAcceptedStatus14Choice struct {

	// Indicates that there is no reason available or to report.
	NoSpecifiedReason *NoReasonCode `xml:"NoSpcfdRsn"`

	// Specifies the reason of the AcknowledgedAcceptedStatus.
	Reason []*AcknowledgementReason7 `xml:"Rsn"`
}

func (a *AcknowledgedAcceptedStatus14Choice) SetNoSpecifiedReason(value string) {
	a.NoSpecifiedReason = (*NoReasonCode)(&value)
}

func (a *AcknowledgedAcceptedStatus14Choice) AddReason() *AcknowledgementReason7 {
	newValue := new(AcknowledgementReason7)
	a.Reason = append(a.Reason, newValue)
	return newValue
}
