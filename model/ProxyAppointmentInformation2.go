package model

// Indicates how to register a proxy.
type ProxyAppointmentInformation2 struct {

	// Indicates how to register a proxy.
	RegistrationMethod *Max350Text `xml:"RegnMtd,omitempty"`

	// Date by which the information on proxy assignment must be received by the intermediary.
	Deadline *DateFormat2Choice `xml:"Ddln,omitempty"`

	// Date by which the information on proxy assignment must be received by the intermediary (STP mode).
	STPDeadline *DateFormat2Choice `xml:"STPDdln,omitempty"`

	// Date by which the information on proxy assignment must be received by the issuer.
	MarketDeadline *DateFormat2Choice `xml:"MktDdln,omitempty"`

	// Specifies the proxy persons which are authorised by the issuer.
	AuthorisedProxy []*Proxy3 `xml:"AuthrsdPrxy,omitempty"`
}

func (p *ProxyAppointmentInformation2) SetRegistrationMethod(value string) {
	p.RegistrationMethod = (*Max350Text)(&value)
}

func (p *ProxyAppointmentInformation2) AddDeadline() *DateFormat2Choice {
	p.Deadline = new(DateFormat2Choice)
	return p.Deadline
}

func (p *ProxyAppointmentInformation2) AddSTPDeadline() *DateFormat2Choice {
	p.STPDeadline = new(DateFormat2Choice)
	return p.STPDeadline
}

func (p *ProxyAppointmentInformation2) AddMarketDeadline() *DateFormat2Choice {
	p.MarketDeadline = new(DateFormat2Choice)
	return p.MarketDeadline
}

func (p *ProxyAppointmentInformation2) AddAuthorisedProxy() *Proxy3 {
	newValue := new(Proxy3)
	p.AuthorisedProxy = append(p.AuthorisedProxy, newValue)
	return newValue
}
