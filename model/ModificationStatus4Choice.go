package model

// Choice of modification status.
type ModificationStatus4Choice struct {

	// Indicates that there is no reason available or to report.
	NoSpecifiedReason *NoReasonCode `xml:"NoSpcfdRsn"`

	// Specifies the reason of the modification status.
	Reason []*ModificationReason4 `xml:"Rsn,omitempty"`
}

func (m *ModificationStatus4Choice) SetNoSpecifiedReason(value string) {
	m.NoSpecifiedReason = (*NoReasonCode)(&value)
}

func (m *ModificationStatus4Choice) AddReason() *ModificationReason4 {
	newValue := new(ModificationReason4)
	m.Reason = append(m.Reason, newValue)
	return newValue
}
