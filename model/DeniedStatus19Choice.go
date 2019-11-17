package model

// Specifies whether the status is provided with a reason or not.
type DeniedStatus19Choice struct {

	// Indicates that there is no reason available or to report.
	NoSpecifiedReason *NoReasonCode `xml:"NoSpcfdRsn"`

	// Specifies the reason of the denied status.
	Reason []*DeniedReason17 `xml:"Rsn"`
}

func (d *DeniedStatus19Choice) SetNoSpecifiedReason(value string) {
	d.NoSpecifiedReason = (*NoReasonCode)(&value)
}

func (d *DeniedStatus19Choice) AddReason() *DeniedReason17 {
	newValue := new(DeniedReason17)
	d.Reason = append(d.Reason, newValue)
	return newValue
}
