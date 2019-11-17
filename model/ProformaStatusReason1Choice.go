package model

// Choice of formats for a proforma status reason.
type ProformaStatusReason1Choice struct {

	// There is no reason available or to report for the proforma account status.
	NoSpecifiedReason *NoReasonCode `xml:"NoSpcfdRsn"`

	// Reason for the proforma account status.
	Reason []*ProformaStatusReason1 `xml:"Rsn"`
}

func (p *ProformaStatusReason1Choice) SetNoSpecifiedReason(value string) {
	p.NoSpecifiedReason = (*NoReasonCode)(&value)
}

func (p *ProformaStatusReason1Choice) AddReason() *ProformaStatusReason1 {
	newValue := new(ProformaStatusReason1)
	p.Reason = append(p.Reason, newValue)
	return newValue
}
