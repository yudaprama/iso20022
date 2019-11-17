package model

// Specifies whether the status is provided with a reason or not.
type AcknowledgedAcceptedStatus24Choice struct {

	// Indicates that there is no reason available or to report.
	NoSpecifiedReason *NoReasonCode `xml:"NoSpcfdRsn"`

	// Specifies the reason of the acknowledged accepted status.
	Reason []*AcknowledgementReason12 `xml:"Rsn"`
}

func (a *AcknowledgedAcceptedStatus24Choice) SetNoSpecifiedReason(value string) {
	a.NoSpecifiedReason = (*NoReasonCode)(&value)
}

func (a *AcknowledgedAcceptedStatus24Choice) AddReason() *AcknowledgementReason12 {
	newValue := new(AcknowledgementReason12)
	a.Reason = append(a.Reason, newValue)
	return newValue
}
