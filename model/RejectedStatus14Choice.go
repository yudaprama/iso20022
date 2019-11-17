package model

// Choice between a reason or no reason for the corporate action instruction processing rejected status.
type RejectedStatus14Choice struct {

	// Reason not specified.
	NoSpecifiedReason *NoReasonCode `xml:"NoSpcfdRsn"`

	// Reason for the rejected status.
	Reason []*RejectedStatusReason14 `xml:"Rsn"`
}

func (r *RejectedStatus14Choice) SetNoSpecifiedReason(value string) {
	r.NoSpecifiedReason = (*NoReasonCode)(&value)
}

func (r *RejectedStatus14Choice) AddReason() *RejectedStatusReason14 {
	newValue := new(RejectedStatusReason14)
	r.Reason = append(r.Reason, newValue)
	return newValue
}
