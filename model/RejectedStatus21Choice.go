package model

// Choice between a reason or no reason for the corporate action instruction processing rejected status.
type RejectedStatus21Choice struct {

	// Reason not specified.
	NoSpecifiedReason *NoReasonCode `xml:"NoSpcfdRsn"`

	// Reason for the rejected status.
	Reason []*RejectedStatusReason20 `xml:"Rsn"`
}

func (r *RejectedStatus21Choice) SetNoSpecifiedReason(value string) {
	r.NoSpecifiedReason = (*NoReasonCode)(&value)
}

func (r *RejectedStatus21Choice) AddReason() *RejectedStatusReason20 {
	newValue := new(RejectedStatusReason20)
	r.Reason = append(r.Reason, newValue)
	return newValue
}
