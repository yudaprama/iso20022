package model

// Specifies whether the status is provided with a reason or not.
type RejectionOrRepairStatus15Choice struct {

	// Indicates that there is no reason available or to report.
	NoSpecifiedReason *NoReasonCode `xml:"NoSpcfdRsn"`

	// Specifies the reason of the RejectionOrRepairStatus.
	Reason []*RejectionOrRepairReason14 `xml:"Rsn"`
}

func (r *RejectionOrRepairStatus15Choice) SetNoSpecifiedReason(value string) {
	r.NoSpecifiedReason = (*NoReasonCode)(&value)
}

func (r *RejectionOrRepairStatus15Choice) AddReason() *RejectionOrRepairReason14 {
	newValue := new(RejectionOrRepairReason14)
	r.Reason = append(r.Reason, newValue)
	return newValue
}
