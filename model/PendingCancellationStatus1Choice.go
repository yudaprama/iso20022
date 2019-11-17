package model

// Choice between a reason or no reason for the corporate action instruction cancellation pending cancellation status.
type PendingCancellationStatus1Choice struct {

	// Reason not specified.
	NotSpecifiedReason *NoReasonCode `xml:"NotSpcfdRsn"`

	// Reason for the pending cancellation status.
	Reason []*PendingCancellationStatusReason1 `xml:"Rsn"`
}

func (p *PendingCancellationStatus1Choice) SetNotSpecifiedReason(value string) {
	p.NotSpecifiedReason = (*NoReasonCode)(&value)
}

func (p *PendingCancellationStatus1Choice) AddReason() *PendingCancellationStatusReason1 {
	newValue := new(PendingCancellationStatusReason1)
	p.Reason = append(p.Reason, newValue)
	return newValue
}
