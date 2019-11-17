package model

// Specifies whether the status is provided with a reason or not.
type RejectionStatus22Choice struct {

	// Indicates that there is no reason available or to report.
	NoSpecifiedReason *NoReasonCode `xml:"NoSpcfdRsn"`

	// Specifies the reason of the rejection status.
	Reason []*RejectionReason34 `xml:"Rsn"`
}

func (r *RejectionStatus22Choice) SetNoSpecifiedReason(value string) {
	r.NoSpecifiedReason = (*NoReasonCode)(&value)
}

func (r *RejectionStatus22Choice) AddReason() *RejectionReason34 {
	newValue := new(RejectionReason34)
	r.Reason = append(r.Reason, newValue)
	return newValue
}
