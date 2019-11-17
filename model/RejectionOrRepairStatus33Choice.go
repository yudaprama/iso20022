package model

// Specifies whether the status is provided with a reason or not.
type RejectionOrRepairStatus33Choice struct {

	// Indicates that there is no reason available or to report.
	NoSpecifiedReason *NoReasonCode `xml:"NoSpcfdRsn"`

	// Specifies the reason of the rejection or repair status.
	Reason []*RejectionOrRepairReason27 `xml:"Rsn"`
}

func (r *RejectionOrRepairStatus33Choice) SetNoSpecifiedReason(value string) {
	r.NoSpecifiedReason = (*NoReasonCode)(&value)
}

func (r *RejectionOrRepairStatus33Choice) AddReason() *RejectionOrRepairReason27 {
	newValue := new(RejectionOrRepairReason27)
	r.Reason = append(r.Reason, newValue)
	return newValue
}
