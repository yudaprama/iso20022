package model

// Status is rejected.
type RejectedStatus5 struct {

	// Reason for a rejected status.
	Reason *RejectedStatusReason6Code `xml:"Rsn"`

	// Reason for a rejected status.
	ExtendedReason *Extended350Code `xml:"XtndedRsn"`
}

func (r *RejectedStatus5) SetReason(value string) {
	r.Reason = (*RejectedStatusReason6Code)(&value)
}

func (r *RejectedStatus5) SetExtendedReason(value string) {
	r.ExtendedReason = (*Extended350Code)(&value)
}
