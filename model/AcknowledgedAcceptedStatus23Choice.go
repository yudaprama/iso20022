package model

// Specifies whether the status is provided with a reason or not.
type AcknowledgedAcceptedStatus23Choice struct {

	// Indicates that there is no reason available or to report.
	NoSpecifiedReason *NoReasonCode `xml:"NoSpcfdRsn"`

	// Specifies the reason of the acknowledged accepted status.
	Reason []*AcknowledgementReason11 `xml:"Rsn"`
}

func (a *AcknowledgedAcceptedStatus23Choice) SetNoSpecifiedReason(value string) {
	a.NoSpecifiedReason = (*NoReasonCode)(&value)
}

func (a *AcknowledgedAcceptedStatus23Choice) AddReason() *AcknowledgementReason11 {
	newValue := new(AcknowledgementReason11)
	a.Reason = append(a.Reason, newValue)
	return newValue
}
