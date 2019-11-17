package model

// Specifies whether the status is provided with a reason or not.
type RejectionStatus8Choice struct {

	// Indicates that there is no reason available or to report.
	NoSpecifiedReason *NoReasonCode `xml:"NoSpcfdRsn"`

	// Specifies the reason of the RejectionStatus.
	Reason []*RejectionReason11 `xml:"Rsn"`
}

func (r *RejectionStatus8Choice) SetNoSpecifiedReason(value string) {
	r.NoSpecifiedReason = (*NoReasonCode)(&value)
}

func (r *RejectionStatus8Choice) AddReason() *RejectionReason11 {
	newValue := new(RejectionReason11)
	r.Reason = append(r.Reason, newValue)
	return newValue
}
