package model

// Specifies whether the status is provided with a reason or not.
type DeniedStatus21Choice struct {

	// Indicates that there is no reason available or to report.
	NoSpecifiedReason *NoReasonCode `xml:"NoSpcfdRsn"`

	// Specifies the reason of the denied status.
	Reason []*DeniedReason16 `xml:"Rsn"`
}

func (d *DeniedStatus21Choice) SetNoSpecifiedReason(value string) {
	d.NoSpecifiedReason = (*NoReasonCode)(&value)
}

func (d *DeniedStatus21Choice) AddReason() *DeniedReason16 {
	newValue := new(DeniedReason16)
	d.Reason = append(d.Reason, newValue)
	return newValue
}
