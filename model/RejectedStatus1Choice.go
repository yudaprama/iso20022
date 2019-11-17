package model

// Choice between a reason or no reason for the corporate action instruction processing rejected status.
type RejectedStatus1Choice struct {

	// Reason not specified.
	NoSpecifiedReason *NoReasonCode `xml:"NoSpcfdRsn"`

	// Reason for the rejected status.
	Reason []*RejectedStatusReason8 `xml:"Rsn"`
}

func (r *RejectedStatus1Choice) SetNoSpecifiedReason(value string) {
	r.NoSpecifiedReason = (*NoReasonCode)(&value)
}

func (r *RejectedStatus1Choice) AddReason() *RejectedStatusReason8 {
	newValue := new(RejectedStatusReason8)
	r.Reason = append(r.Reason, newValue)
	return newValue
}
