package model

// Choice between a reason or no reason for the corporate action instruction processing pending status.
type PendingStatus32Choice struct {

	// Reason not specified.
	NoSpecifiedReason *NoReasonCode `xml:"NoSpcfdRsn"`

	// Reason for the pending status.
	Reason []*PendingStatusReason5 `xml:"Rsn"`
}

func (p *PendingStatus32Choice) SetNoSpecifiedReason(value string) {
	p.NoSpecifiedReason = (*NoReasonCode)(&value)
}

func (p *PendingStatus32Choice) AddReason() *PendingStatusReason5 {
	newValue := new(PendingStatusReason5)
	p.Reason = append(p.Reason, newValue)
	return newValue
}
