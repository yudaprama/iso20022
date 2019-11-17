package model

// Choice between a reason or no reason for the corporate action instruction processing rejected status.
type RejectedStatus20Choice struct {

	// Reason not specified.
	NoSpecifiedReason *NoReasonCode `xml:"NoSpcfdRsn"`

	// Reason for the rejected status.
	Reason []*RejectedStatusReason19 `xml:"Rsn"`
}

func (r *RejectedStatus20Choice) SetNoSpecifiedReason(value string) {
	r.NoSpecifiedReason = (*NoReasonCode)(&value)
}

func (r *RejectedStatus20Choice) AddReason() *RejectedStatusReason19 {
	newValue := new(RejectedStatusReason19)
	r.Reason = append(r.Reason, newValue)
	return newValue
}
