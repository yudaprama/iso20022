package model

// Specifies whether the status is provided with a reason or not.
type DeniedStatus5Choice struct {

	// Indicates that there is no reason available or to report.
	NoSpecifiedReason *NoReasonCode `xml:"NoSpcfdRsn"`

	// Specifies the reason of the DeniedStatus.
	Reason []*DeniedReason1 `xml:"Rsn"`
}

func (d *DeniedStatus5Choice) SetNoSpecifiedReason(value string) {
	d.NoSpecifiedReason = (*NoReasonCode)(&value)
}

func (d *DeniedStatus5Choice) AddReason() *DeniedReason1 {
	newValue := new(DeniedReason1)
	d.Reason = append(d.Reason, newValue)
	return newValue
}
