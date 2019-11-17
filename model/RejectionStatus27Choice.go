package model

// Choice of  rejection status.
type RejectionStatus27Choice struct {

	// Indicates that there is no reason available or to report.
	NoSpecifiedReason *NoReasonCode `xml:"NoSpcfdRsn"`

	// Specifies the reason of the rejection or repair status.
	Reason *RejectionReason40 `xml:"Rsn"`
}

func (r *RejectionStatus27Choice) SetNoSpecifiedReason(value string) {
	r.NoSpecifiedReason = (*NoReasonCode)(&value)
}

func (r *RejectionStatus27Choice) AddReason() *RejectionReason40 {
	r.Reason = new(RejectionReason40)
	return r.Reason
}
