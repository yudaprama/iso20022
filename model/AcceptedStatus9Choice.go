package model

// Choice between a reason or no reason for the corporate action instruction processing accepted status.
type AcceptedStatus9Choice struct {

	// Reason not specified.
	NoSpecifiedReason *NoReasonCode `xml:"NoSpcfdRsn"`

	// Reason for the accepted status.
	Reason []*AcceptedStatusReason10 `xml:"Rsn"`
}

func (a *AcceptedStatus9Choice) SetNoSpecifiedReason(value string) {
	a.NoSpecifiedReason = (*NoReasonCode)(&value)
}

func (a *AcceptedStatus9Choice) AddReason() *AcceptedStatusReason10 {
	newValue := new(AcceptedStatusReason10)
	a.Reason = append(a.Reason, newValue)
	return newValue
}
