package model

// Specifies whether the status is provided with a reason or not.
type RejectionOrRepairStatus32Choice struct {

	// Indicates that there is no reason available or to report.
	NoSpecifiedReason *NoReasonCode `xml:"NoSpcfdRsn"`

	// Specifies the reason of the rejection or repair status.
	Reason []*RejectionOrRepairReason26 `xml:"Rsn"`
}

func (r *RejectionOrRepairStatus32Choice) SetNoSpecifiedReason(value string) {
	r.NoSpecifiedReason = (*NoReasonCode)(&value)
}

func (r *RejectionOrRepairStatus32Choice) AddReason() *RejectionOrRepairReason26 {
	newValue := new(RejectionOrRepairReason26)
	r.Reason = append(r.Reason, newValue)
	return newValue
}
