package model

// Choice between a reason or no reason for the corporate action instruction cancellation pending cancellation status.
type PendingCancellationStatus5Choice struct {

	// Reason not specified.
	NotSpecifiedReason *NoReasonCode `xml:"NotSpcfdRsn"`

	// Reason for the pending cancellation status.
	Reason []*PendingCancellationStatusReason5 `xml:"Rsn"`
}

func (p *PendingCancellationStatus5Choice) SetNotSpecifiedReason(value string) {
	p.NotSpecifiedReason = (*NoReasonCode)(&value)
}

func (p *PendingCancellationStatus5Choice) AddReason() *PendingCancellationStatusReason5 {
	newValue := new(PendingCancellationStatusReason5)
	p.Reason = append(p.Reason, newValue)
	return newValue
}
