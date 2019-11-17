package model

// Choice of formats for the reason of an accepted status.
type AcceptedReason8Choice struct {

	// Indicates that there is no reason available or to report.
	NoSpecifiedReason *NoReasonCode `xml:"NoSpcfdRsn"`

	// Reason for the accepted status.
	Reason *AcceptedReason7Choice `xml:"Rsn"`
}

func (a *AcceptedReason8Choice) SetNoSpecifiedReason(value string) {
	a.NoSpecifiedReason = (*NoReasonCode)(&value)
}

func (a *AcceptedReason8Choice) AddReason() *AcceptedReason7Choice {
	a.Reason = new(AcceptedReason7Choice)
	return a.Reason
}
