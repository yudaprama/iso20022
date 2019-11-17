package model

// Specifies whether the status is provided with a reason or not.
type RejectionStatus24Choice struct {

	// Indicates that there is no reason available or to report.
	NoSpecifiedReason *NoReasonCode `xml:"NoSpcfdRsn"`

	// Specifies the reason of the rejection status.
	Reason []*RejectionReason37 `xml:"Rsn"`
}

func (r *RejectionStatus24Choice) SetNoSpecifiedReason(value string) {
	r.NoSpecifiedReason = (*NoReasonCode)(&value)
}

func (r *RejectionStatus24Choice) AddReason() *RejectionReason37 {
	newValue := new(RejectionReason37)
	r.Reason = append(r.Reason, newValue)
	return newValue
}
