package model

// Specifies whether the status is provided with a reason or not.
type AcknowledgedAcceptedStatus21Choice struct {

	// Indicates that there is no reason available or to report.
	NoSpecifiedReason *NoReasonCode `xml:"NoSpcfdRsn"`

	// Specifies the reason of the acknowledged accepted status.
	Reason []*AcknowledgementReason9 `xml:"Rsn"`
}

func (a *AcknowledgedAcceptedStatus21Choice) SetNoSpecifiedReason(value string) {
	a.NoSpecifiedReason = (*NoReasonCode)(&value)
}

func (a *AcknowledgedAcceptedStatus21Choice) AddReason() *AcknowledgementReason9 {
	newValue := new(AcknowledgementReason9)
	a.Reason = append(a.Reason, newValue)
	return newValue
}
