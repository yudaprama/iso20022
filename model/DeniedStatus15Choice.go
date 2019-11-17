package model

// Specifies whether the status is provided with a reason or not.
type DeniedStatus15Choice struct {

	// Indicates that there is no reason available or to report.
	NoSpecifiedReason *NoReasonCode `xml:"NoSpcfdRsn"`

	// Specifies the reason of the denied status.
	Reason []*DeniedReason10 `xml:"Rsn"`
}

func (d *DeniedStatus15Choice) SetNoSpecifiedReason(value string) {
	d.NoSpecifiedReason = (*NoReasonCode)(&value)
}

func (d *DeniedStatus15Choice) AddReason() *DeniedReason10 {
	newValue := new(DeniedReason10)
	d.Reason = append(d.Reason, newValue)
	return newValue
}
