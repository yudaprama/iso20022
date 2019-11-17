package model

// Choice between a reason or no reason for the corporate action instruction processing rejected status.
type RejectedStatus19Choice struct {

	// Reason not specified.
	NoSpecifiedReason *NoReasonCode `xml:"NoSpcfdRsn"`

	// Reason for the rejected status.
	Reason []*RejectedStatusReason17 `xml:"Rsn"`
}

func (r *RejectedStatus19Choice) SetNoSpecifiedReason(value string) {
	r.NoSpecifiedReason = (*NoReasonCode)(&value)
}

func (r *RejectedStatus19Choice) AddReason() *RejectedStatusReason17 {
	newValue := new(RejectedStatusReason17)
	r.Reason = append(r.Reason, newValue)
	return newValue
}
