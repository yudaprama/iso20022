package model

// Choice between a reason or no reason for the corporate action instruction cancellation pending cancellation status.
type PendingCancellationStatus3Choice struct {

	// Reason not specified.
	NotSpecifiedReason *NoReasonCode `xml:"NotSpcfdRsn"`

	// Reason for the pending cancellation status.
	Reason []*PendingCancellationStatusReason3 `xml:"Rsn"`
}

func (p *PendingCancellationStatus3Choice) SetNotSpecifiedReason(value string) {
	p.NotSpecifiedReason = (*NoReasonCode)(&value)
}

func (p *PendingCancellationStatus3Choice) AddReason() *PendingCancellationStatusReason3 {
	newValue := new(PendingCancellationStatusReason3)
	p.Reason = append(p.Reason, newValue)
	return newValue
}
