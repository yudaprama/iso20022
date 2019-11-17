package model

// Choice of formats for the reason of a rejected status.
type RejectedReason8Choice struct {

	// Indicates that there is no reason available or to report.
	NoSpecifiedReason *NoReasonCode `xml:"NoSpcfdRsn"`

	// Reason for the rejected status.
	Reason []*RejectedReason7Choice `xml:"Rsn"`
}

func (r *RejectedReason8Choice) SetNoSpecifiedReason(value string) {
	r.NoSpecifiedReason = (*NoReasonCode)(&value)
}

func (r *RejectedReason8Choice) AddReason() *RejectedReason7Choice {
	newValue := new(RejectedReason7Choice)
	r.Reason = append(r.Reason, newValue)
	return newValue
}
