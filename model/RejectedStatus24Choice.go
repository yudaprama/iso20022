package model

// Choice between a reason or no reason for the corporate action instruction processing rejected status.
type RejectedStatus24Choice struct {

	// Reason not specified.
	NoSpecifiedReason *NoReasonCode `xml:"NoSpcfdRsn"`

	// Reason for the rejected status.
	Reason []*RejectedStatusReason22 `xml:"Rsn"`
}

func (r *RejectedStatus24Choice) SetNoSpecifiedReason(value string) {
	r.NoSpecifiedReason = (*NoReasonCode)(&value)
}

func (r *RejectedStatus24Choice) AddReason() *RejectedStatusReason22 {
	newValue := new(RejectedStatusReason22)
	r.Reason = append(r.Reason, newValue)
	return newValue
}
