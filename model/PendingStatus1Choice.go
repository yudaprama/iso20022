package model

// Choice between a reason or no reason for the corporate action instruction processing pending status.
type PendingStatus1Choice struct {

	// Reason not specified.
	NoSpecifiedReason *NoReasonCode `xml:"NoSpcfdRsn"`

	// Reason for the pending status.
	Reason []*PendingStatusReason1 `xml:"Rsn"`
}

func (p *PendingStatus1Choice) SetNoSpecifiedReason(value string) {
	p.NoSpecifiedReason = (*NoReasonCode)(&value)
}

func (p *PendingStatus1Choice) AddReason() *PendingStatusReason1 {
	newValue := new(PendingStatusReason1)
	p.Reason = append(p.Reason, newValue)
	return newValue
}
