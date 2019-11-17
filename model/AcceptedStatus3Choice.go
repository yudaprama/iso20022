package model

// Choice between a reason or no reason for the corporate action instruction processing accepted status.
type AcceptedStatus3Choice struct {

	// Reason not specified.
	NoSpecifiedReason *NoReasonCode `xml:"NoSpcfdRsn"`

	// Reason for the accepted status.
	Reason []*AcceptedStatusReason3 `xml:"Rsn"`
}

func (a *AcceptedStatus3Choice) SetNoSpecifiedReason(value string) {
	a.NoSpecifiedReason = (*NoReasonCode)(&value)
}

func (a *AcceptedStatus3Choice) AddReason() *AcceptedStatusReason3 {
	newValue := new(AcceptedStatusReason3)
	a.Reason = append(a.Reason, newValue)
	return newValue
}
