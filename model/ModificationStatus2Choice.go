package model

// Choice of modification status.
type ModificationStatus2Choice struct {

	// Indicates that there is no reason available or to report.
	NoSpecifiedReason *NoReasonCode `xml:"NoSpcfdRsn"`

	// Specifies the reason of the modification status.
	Reason []*ModificationReason2 `xml:"Rsn,omitempty"`
}

func (m *ModificationStatus2Choice) SetNoSpecifiedReason(value string) {
	m.NoSpecifiedReason = (*NoReasonCode)(&value)
}

func (m *ModificationStatus2Choice) AddReason() *ModificationReason2 {
	newValue := new(ModificationReason2)
	m.Reason = append(m.Reason, newValue)
	return newValue
}
