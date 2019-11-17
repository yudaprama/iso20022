package model

// Choice between a reason or no reason for the corporate action instruction cancellation pending cancellation status.
type PendingCancellationStatus6Choice struct {

	// Reason not specified.
	NotSpecifiedReason *NoReasonCode `xml:"NotSpcfdRsn"`

	// Reason for the pending cancellation status.
	Reason []*PendingCancellationStatusReason6 `xml:"Rsn"`
}

func (p *PendingCancellationStatus6Choice) SetNotSpecifiedReason(value string) {
	p.NotSpecifiedReason = (*NoReasonCode)(&value)
}

func (p *PendingCancellationStatus6Choice) AddReason() *PendingCancellationStatusReason6 {
	newValue := new(PendingCancellationStatusReason6)
	p.Reason = append(p.Reason, newValue)
	return newValue
}
