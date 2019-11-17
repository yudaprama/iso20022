package model

// Choice between a reason or no reason for the corporate action instruction processing accepted status.
type AcceptedStatus8Choice struct {

	// Reason not specified.
	NoSpecifiedReason *NoReasonCode `xml:"NoSpcfdRsn"`

	// Reason for the accepted status.
	Reason []*AcceptedStatusReason9 `xml:"Rsn"`
}

func (a *AcceptedStatus8Choice) SetNoSpecifiedReason(value string) {
	a.NoSpecifiedReason = (*NoReasonCode)(&value)
}

func (a *AcceptedStatus8Choice) AddReason() *AcceptedStatusReason9 {
	newValue := new(AcceptedStatusReason9)
	a.Reason = append(a.Reason, newValue)
	return newValue
}
