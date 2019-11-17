package model

// Specifies whether the status is provided with a reason or not.
type AcknowledgedAcceptedStatus30Choice struct {

	// Indicates that there is no reason available or to report.
	NoSpecifiedReason *NoReasonCode `xml:"NoSpcfdRsn"`

	// Specifies the reason of the acknowledged accepted status.
	Reason []*AcknowledgementReason18 `xml:"Rsn"`
}

func (a *AcknowledgedAcceptedStatus30Choice) SetNoSpecifiedReason(value string) {
	a.NoSpecifiedReason = (*NoReasonCode)(&value)
}

func (a *AcknowledgedAcceptedStatus30Choice) AddReason() *AcknowledgementReason18 {
	newValue := new(AcknowledgementReason18)
	a.Reason = append(a.Reason, newValue)
	return newValue
}
