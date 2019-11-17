package model

// Choice between a reason or no reason for the corporate action instruction processing accepted status.
type AcceptedStatus10Choice struct {

	// Reason not specified.
	NoSpecifiedReason *NoReasonCode `xml:"NoSpcfdRsn"`

	// Reason for the accepted status.
	Reason []*AcceptedStatusReason11 `xml:"Rsn"`
}

func (a *AcceptedStatus10Choice) SetNoSpecifiedReason(value string) {
	a.NoSpecifiedReason = (*NoReasonCode)(&value)
}

func (a *AcceptedStatus10Choice) AddReason() *AcceptedStatusReason11 {
	newValue := new(AcceptedStatusReason11)
	a.Reason = append(a.Reason, newValue)
	return newValue
}
