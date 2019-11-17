package model

// Choice of formats for the reason for an in repair status.
type InRepairStatusReason4Choice struct {

	// No reason available or to report for the in repair status.
	NoSpecifiedReason *NoReasonCode `xml:"NoSpcfdRsn"`

	// Reason for the in repair status, expressed as a code.
	ReasonDetails []*InRepairStatusReason4 `xml:"RsnDtls"`
}

func (i *InRepairStatusReason4Choice) SetNoSpecifiedReason(value string) {
	i.NoSpecifiedReason = (*NoReasonCode)(&value)
}

func (i *InRepairStatusReason4Choice) AddReasonDetails() *InRepairStatusReason4 {
	newValue := new(InRepairStatusReason4)
	i.ReasonDetails = append(i.ReasonDetails, newValue)
	return newValue
}
