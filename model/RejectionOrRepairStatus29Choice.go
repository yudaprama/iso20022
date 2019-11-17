package model

// Specifies whether the status is provided with a reason or not.
type RejectionOrRepairStatus29Choice struct {

	// Indicates that there is no reason available or to report.
	NoSpecifiedReason *NoReasonCode `xml:"NoSpcfdRsn"`

	// Specifies the reason of the rejection or repair status.
	Reason []*RejectionOrRepairReason23 `xml:"Rsn"`
}

func (r *RejectionOrRepairStatus29Choice) SetNoSpecifiedReason(value string) {
	r.NoSpecifiedReason = (*NoReasonCode)(&value)
}

func (r *RejectionOrRepairStatus29Choice) AddReason() *RejectionOrRepairReason23 {
	newValue := new(RejectionOrRepairReason23)
	r.Reason = append(r.Reason, newValue)
	return newValue
}
