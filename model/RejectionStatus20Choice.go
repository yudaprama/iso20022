package model

// Choice of  rejection status.
type RejectionStatus20Choice struct {

	// Indicates that there is no reason available or to report.
	NoSpecifiedReason *NoReasonCode `xml:"NoSpcfdRsn"`

	// Specifies the reason of the rejection or repair status.
	Reason *RejectionReason29 `xml:"Rsn"`
}

func (r *RejectionStatus20Choice) SetNoSpecifiedReason(value string) {
	r.NoSpecifiedReason = (*NoReasonCode)(&value)
}

func (r *RejectionStatus20Choice) AddReason() *RejectionReason29 {
	r.Reason = new(RejectionReason29)
	return r.Reason
}
