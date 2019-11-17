package model

// Choice between a reason or no reason for the corporate action instruction processing accepted status.
type AcceptedStatus4Choice struct {

	// Reason not specified.
	NoSpecifiedReason *NoReasonCode `xml:"NoSpcfdRsn"`

	// Reason for the accepted status.
	Reason []*AcceptedStatusReason4 `xml:"Rsn"`
}

func (a *AcceptedStatus4Choice) SetNoSpecifiedReason(value string) {
	a.NoSpecifiedReason = (*NoReasonCode)(&value)
}

func (a *AcceptedStatus4Choice) AddReason() *AcceptedStatusReason4 {
	newValue := new(AcceptedStatusReason4)
	a.Reason = append(a.Reason, newValue)
	return newValue
}
