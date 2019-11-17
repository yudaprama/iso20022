package model

// Choice between a reason or no reason for the corporate action instruction cancellation pending cancellation status.
type PendingCancellationStatus8Choice struct {

	// Reason not specified.
	NotSpecifiedReason *NoReasonCode `xml:"NotSpcfdRsn"`

	// Reason for the pending cancellation status.
	Reason []*PendingCancellationStatusReason8 `xml:"Rsn"`
}

func (p *PendingCancellationStatus8Choice) SetNotSpecifiedReason(value string) {
	p.NotSpecifiedReason = (*NoReasonCode)(&value)
}

func (p *PendingCancellationStatus8Choice) AddReason() *PendingCancellationStatusReason8 {
	newValue := new(PendingCancellationStatusReason8)
	p.Reason = append(p.Reason, newValue)
	return newValue
}
