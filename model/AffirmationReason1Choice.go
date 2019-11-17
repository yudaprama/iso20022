package model

// Choice of format for the affirmation status.
type AffirmationReason1Choice struct {

	// Specifies the reason of the UnaffirmedStatus.
	Reason []*AffirmationReason1 `xml:"Rsn"`

	// Indicates that there is no reason available or to report.
	NoSpecifiedReason *NoReasonCode `xml:"NoSpcfdRsn"`
}

func (a *AffirmationReason1Choice) AddReason() *AffirmationReason1 {
	newValue := new(AffirmationReason1)
	a.Reason = append(a.Reason, newValue)
	return newValue
}

func (a *AffirmationReason1Choice) SetNoSpecifiedReason(value string) {
	a.NoSpecifiedReason = (*NoReasonCode)(&value)
}
