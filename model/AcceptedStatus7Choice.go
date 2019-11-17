package model

// Choice between a reason or no reason for the corporate action instruction processing accepted status.
type AcceptedStatus7Choice struct {

	// Reason not specified.
	NoSpecifiedReason *NoReasonCode `xml:"NoSpcfdRsn"`

	// Reason for the accepted status.
	Reason []*AcceptedStatusReason8 `xml:"Rsn"`
}

func (a *AcceptedStatus7Choice) SetNoSpecifiedReason(value string) {
	a.NoSpecifiedReason = (*NoReasonCode)(&value)
}

func (a *AcceptedStatus7Choice) AddReason() *AcceptedStatusReason8 {
	newValue := new(AcceptedStatusReason8)
	a.Reason = append(a.Reason, newValue)
	return newValue
}
