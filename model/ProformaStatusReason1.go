package model

// Reason for an proforma status.
type ProformaStatusReason1 struct {

	// Reason for the proforma account status.
	Code *ProformaStatusReason2Choice `xml:"Cd"`

	// Additional information about the reason for the proforma account status.
	AdditionalInformation *Max350Text `xml:"AddtlInf,omitempty"`
}

func (p *ProformaStatusReason1) AddCode() *ProformaStatusReason2Choice {
	p.Code = new(ProformaStatusReason2Choice)
	return p.Code
}

func (p *ProformaStatusReason1) SetAdditionalInformation(value string) {
	p.AdditionalInformation = (*Max350Text)(&value)
}
