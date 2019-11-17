package model

// Choice of formats for a proforma reason code.
type ProformaStatusReason2Choice struct {

	// Reason for the enabled account status expressed as a code.
	Code *ProformaStatusReason1Code `xml:"Cd"`

	// Reason for the enabled account status expressed as a proprietary code.
	Proprietary *GenericIdentification36 `xml:"Prtry"`
}

func (p *ProformaStatusReason2Choice) SetCode(value string) {
	p.Code = (*ProformaStatusReason1Code)(&value)
}

func (p *ProformaStatusReason2Choice) AddProprietary() *GenericIdentification36 {
	p.Proprietary = new(GenericIdentification36)
	return p.Proprietary
}
