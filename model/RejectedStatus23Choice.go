package model

// Choice between a reason or no reason for the corporate action instruction processing rejected status.
type RejectedStatus23Choice struct {

	// Reason not specified.
	NoSpecifiedReason *NoReasonCode `xml:"NoSpcfdRsn"`

	// Reason for the rejected status.
	Reason []*RejectedStatusReason21 `xml:"Rsn"`
}

func (r *RejectedStatus23Choice) SetNoSpecifiedReason(value string) {
	r.NoSpecifiedReason = (*NoReasonCode)(&value)
}

func (r *RejectedStatus23Choice) AddReason() *RejectedStatusReason21 {
	newValue := new(RejectedStatusReason21)
	r.Reason = append(r.Reason, newValue)
	return newValue
}
