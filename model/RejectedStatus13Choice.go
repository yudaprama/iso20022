package model

// Choice between a reason or no reason for the corporate action instruction processing rejected status.
type RejectedStatus13Choice struct {

	// Reason not specified.
	NoSpecifiedReason *NoReasonCode `xml:"NoSpcfdRsn"`

	// Reason for the rejected status.
	Reason []*RejectedStatusReason13 `xml:"Rsn"`
}

func (r *RejectedStatus13Choice) SetNoSpecifiedReason(value string) {
	r.NoSpecifiedReason = (*NoReasonCode)(&value)
}

func (r *RejectedStatus13Choice) AddReason() *RejectedStatusReason13 {
	newValue := new(RejectedStatusReason13)
	r.Reason = append(r.Reason, newValue)
	return newValue
}
