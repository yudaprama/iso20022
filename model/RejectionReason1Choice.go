package model

// Allows the sender of the rejection message to indicate only one rejection reason that applies to the entire rejected message.
type RejectionReason1Choice struct {

	// Rejection reason that applies to the whole report.
	GlobalRejectionReason *Reason2 `xml:"GblRjctnRsn"`

	// Specifies a rejection reason for each individual element of a report.
	RejectedElement []*RejectedElement1 `xml:"RjctdElmt"`
}

func (r *RejectionReason1Choice) AddGlobalRejectionReason() *Reason2 {
	r.GlobalRejectionReason = new(Reason2)
	return r.GlobalRejectionReason
}

func (r *RejectionReason1Choice) AddRejectedElement() *RejectedElement1 {
	newValue := new(RejectedElement1)
	r.RejectedElement = append(r.RejectedElement, newValue)
	return newValue
}
