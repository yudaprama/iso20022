package model

// Status is in repair.
type InRepairStatus2 struct {

	// Reason for the in-repair status.
	ReasonDetails *InRepairStatusReason3 `xml:"RsnDtls"`

	// Indicates that there is no reason available or to report.
	NoSpecifiedReason *NoReasonCode `xml:"NoSpcfdRsn"`
}

func (i *InRepairStatus2) AddReasonDetails() *InRepairStatusReason3 {
	i.ReasonDetails = new(InRepairStatusReason3)
	return i.ReasonDetails
}

func (i *InRepairStatus2) SetNoSpecifiedReason(value string) {
	i.NoSpecifiedReason = (*NoReasonCode)(&value)
}
