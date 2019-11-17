package model

// No specified reason available for the status.
type NoSpecifiedReason1 struct {

	// Indicates that there is no reason available or to report.
	NoSpecifiedReason *NoReasonCode `xml:"NoSpcfdRsn"`
}

func (n *NoSpecifiedReason1) SetNoSpecifiedReason(value string) {
	n.NoSpecifiedReason = (*NoReasonCode)(&value)
}
