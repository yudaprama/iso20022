package model

// Specifies whether the status is provided with a reason or not.
type RejectionOrRepairStatus37Choice struct {

	// Indicates that there is no reason available or to report.
	NoSpecifiedReason *NoReasonCode `xml:"NoSpcfdRsn"`

	// Specifies the reason of the rejection or repair status.
	Reason []*RejectionOrRepairReason31 `xml:"Rsn"`
}

func (r *RejectionOrRepairStatus37Choice) SetNoSpecifiedReason(value string) {
	r.NoSpecifiedReason = (*NoReasonCode)(&value)
}

func (r *RejectionOrRepairStatus37Choice) AddReason() *RejectionOrRepairReason31 {
	newValue := new(RejectionOrRepairReason31)
	r.Reason = append(r.Reason, newValue)
	return newValue
}
