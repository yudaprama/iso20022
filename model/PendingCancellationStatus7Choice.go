package model

// Choice between a reason or no reason for the corporate action instruction cancellation pending cancellation status.
type PendingCancellationStatus7Choice struct {

	// Reason not specified.
	NotSpecifiedReason *NoReasonCode `xml:"NotSpcfdRsn"`

	// Reason for the pending cancellation status.
	Reason []*PendingCancellationStatusReason7 `xml:"Rsn"`
}

func (p *PendingCancellationStatus7Choice) SetNotSpecifiedReason(value string) {
	p.NotSpecifiedReason = (*NoReasonCode)(&value)
}

func (p *PendingCancellationStatus7Choice) AddReason() *PendingCancellationStatusReason7 {
	newValue := new(PendingCancellationStatusReason7)
	p.Reason = append(p.Reason, newValue)
	return newValue
}
