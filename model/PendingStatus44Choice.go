package model

// Choice between a reason or no reason for the corporate action instruction processing pending status.
type PendingStatus44Choice struct {

	// Reason not specified.
	NoSpecifiedReason *NoReasonCode `xml:"NoSpcfdRsn"`

	// Reason for the pending status.
	Reason []*PendingStatusReason12 `xml:"Rsn"`
}

func (p *PendingStatus44Choice) SetNoSpecifiedReason(value string) {
	p.NoSpecifiedReason = (*NoReasonCode)(&value)
}

func (p *PendingStatus44Choice) AddReason() *PendingStatusReason12 {
	newValue := new(PendingStatusReason12)
	p.Reason = append(p.Reason, newValue)
	return newValue
}
