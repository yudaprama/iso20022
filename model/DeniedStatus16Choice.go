package model

// Specifies whether the status is provided with a reason or not.
type DeniedStatus16Choice struct {

	// Indicates that there is no reason available or to report.
	NoSpecifiedReason *NoReasonCode `xml:"NoSpcfdRsn"`

	// Specifies the reason of the denied status.
	Reason []*DeniedReason11 `xml:"Rsn"`
}

func (d *DeniedStatus16Choice) SetNoSpecifiedReason(value string) {
	d.NoSpecifiedReason = (*NoReasonCode)(&value)
}

func (d *DeniedStatus16Choice) AddReason() *DeniedReason11 {
	newValue := new(DeniedReason11)
	d.Reason = append(d.Reason, newValue)
	return newValue
}
