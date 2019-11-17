package model

// Specifies whether the status is provided with a reason or not.
type RejectionOrRepairStatus14Choice struct {

	// Indicates that there is no reason available or to report.
	NoSpecifiedReason *NoReasonCode `xml:"NoSpcfdRsn"`

	// Specifies the reason for the status.
	Reason []*RejectionOrRepairReason13 `xml:"Rsn"`
}

func (r *RejectionOrRepairStatus14Choice) SetNoSpecifiedReason(value string) {
	r.NoSpecifiedReason = (*NoReasonCode)(&value)
}

func (r *RejectionOrRepairStatus14Choice) AddReason() *RejectionOrRepairReason13 {
	newValue := new(RejectionOrRepairReason13)
	r.Reason = append(r.Reason, newValue)
	return newValue
}
