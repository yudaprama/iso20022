package model

// Specifies whether the status is provided with a reason or not.
type DeniedStatus17Choice struct {

	// Indicates that there is no reason available or to report.
	NoSpecifiedReason *NoReasonCode `xml:"NoSpcfdRsn"`

	// Specifies the reason of the denied status.
	Reason []*DeniedReason12 `xml:"Rsn"`
}

func (d *DeniedStatus17Choice) SetNoSpecifiedReason(value string) {
	d.NoSpecifiedReason = (*NoReasonCode)(&value)
}

func (d *DeniedStatus17Choice) AddReason() *DeniedReason12 {
	newValue := new(DeniedReason12)
	d.Reason = append(d.Reason, newValue)
	return newValue
}
