package model

// Choice between a reason or no reason for the corporate action instruction processing rejected status.
type RejectedStatus9Choice struct {

	// Reason not specified.
	NoSpecifiedReason *NoReasonCode `xml:"NoSpcfdRsn"`

	// Reason for the rejected status.
	Reason []*RejectedStatusReason10 `xml:"Rsn"`
}

func (r *RejectedStatus9Choice) SetNoSpecifiedReason(value string) {
	r.NoSpecifiedReason = (*NoReasonCode)(&value)
}

func (r *RejectedStatus9Choice) AddReason() *RejectedStatusReason10 {
	newValue := new(RejectedStatusReason10)
	r.Reason = append(r.Reason, newValue)
	return newValue
}
