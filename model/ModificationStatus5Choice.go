package model

// Choice of modification status.
type ModificationStatus5Choice struct {

	// Indicates that there is no reason available or to report.
	NoSpecifiedReason *NoReasonCode `xml:"NoSpcfdRsn"`

	// Specifies the reason of the modification status.
	Reason []*ModificationReason5 `xml:"Rsn,omitempty"`
}

func (m *ModificationStatus5Choice) SetNoSpecifiedReason(value string) {
	m.NoSpecifiedReason = (*NoReasonCode)(&value)
}

func (m *ModificationStatus5Choice) AddReason() *ModificationReason5 {
	newValue := new(ModificationReason5)
	m.Reason = append(m.Reason, newValue)
	return newValue
}
