package model

// Proprietary identification of the reason related to a status.
type ProprietaryReason1Choice struct {

	// Indicates that there is no reason available or to report.
	NoSpecifiedReason *NoReasonCode `xml:"NoSpcfdRsn"`

	// Proprietary identification of the reason for the status.
	Reason []*GenericIdentification36 `xml:"Rsn"`
}

func (p *ProprietaryReason1Choice) SetNoSpecifiedReason(value string) {
	p.NoSpecifiedReason = (*NoReasonCode)(&value)
}

func (p *ProprietaryReason1Choice) AddReason() *GenericIdentification36 {
	newValue := new(GenericIdentification36)
	p.Reason = append(p.Reason, newValue)
	return newValue
}
