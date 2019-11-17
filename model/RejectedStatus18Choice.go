package model

// Choice between a reason or no reason for the corporate action instruction processing rejected status.
type RejectedStatus18Choice struct {

	// Reason not specified.
	NoSpecifiedReason *NoReasonCode `xml:"NoSpcfdRsn"`

	// Reason for the rejected status.
	Reason []*RejectedStatusReason18 `xml:"Rsn"`
}

func (r *RejectedStatus18Choice) SetNoSpecifiedReason(value string) {
	r.NoSpecifiedReason = (*NoReasonCode)(&value)
}

func (r *RejectedStatus18Choice) AddReason() *RejectedStatusReason18 {
	newValue := new(RejectedStatusReason18)
	r.Reason = append(r.Reason, newValue)
	return newValue
}
