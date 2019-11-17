package model

// Specifies whether the status is provided with a reason or not.
type DeniedStatus2Choice struct {

	// Indicates that there is no reason available or to report.
	NoSpecifiedReason *NoReasonCode `xml:"NoSpcfdRsn"`

	// Specifies the reason of the DeniedStatus.
	Reason []*DeniedReason2 `xml:"Rsn,omitempty"`
}

func (d *DeniedStatus2Choice) SetNoSpecifiedReason(value string) {
	d.NoSpecifiedReason = (*NoReasonCode)(&value)
}

func (d *DeniedStatus2Choice) AddReason() *DeniedReason2 {
	newValue := new(DeniedReason2)
	d.Reason = append(d.Reason, newValue)
	return newValue
}
