package model

// Specifies whether the status is provided with a reason or not.
type RejectionStatus18Choice struct {

	// Indicates that there is no reason available or to report.
	NoSpecifiedReason *NoReasonCode `xml:"NoSpcfdRsn"`

	// Specifies the reason of the rejection status.
	Reason []*RejectionReason26 `xml:"Rsn"`
}

func (r *RejectionStatus18Choice) SetNoSpecifiedReason(value string) {
	r.NoSpecifiedReason = (*NoReasonCode)(&value)
}

func (r *RejectionStatus18Choice) AddReason() *RejectionReason26 {
	newValue := new(RejectionReason26)
	r.Reason = append(r.Reason, newValue)
	return newValue
}
