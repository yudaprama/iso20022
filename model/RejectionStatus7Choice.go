package model

// Choice of  rejection status.
type RejectionStatus7Choice struct {

	// Indicates that there is no reason available or to report.
	NoSpecifiedReason *NoReasonCode `xml:"NoSpcfdRsn"`

	// Specifies the reason of the rejection or repair status.
	Reason *RejectionReason12 `xml:"Rsn"`
}

func (r *RejectionStatus7Choice) SetNoSpecifiedReason(value string) {
	r.NoSpecifiedReason = (*NoReasonCode)(&value)
}

func (r *RejectionStatus7Choice) AddReason() *RejectionReason12 {
	r.Reason = new(RejectionReason12)
	return r.Reason
}
