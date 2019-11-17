package model

// Indicates how to register a proxy.
type ProxyAppointmentInformation1 struct {

	// Indicates how to register a proxy.
	RegistrationMethod *Max350Text `xml:"RegnMtd,omitempty"`

	// Date by which the information on proxy assignment must be received by the intermediary.
	Deadline *DateFormat2Choice `xml:"Ddln,omitempty"`

	// Date by which the information on proxy assignment must be received by the intermediary (STP mode).
	STPDeadline *DateFormat2Choice `xml:"STPDdln,omitempty"`

	// Date by which the information on proxy assignment must be received by the issuer.
	MarketDeadline *DateFormat2Choice `xml:"MktDdln,omitempty"`

	// Specifies the proxy persons which are authorised by the issuer.
	AuthorisedProxy []*Proxy1 `xml:"AuthrsdPrxy,omitempty"`
}

func (p *ProxyAppointmentInformation1) SetRegistrationMethod(value string) {
	p.RegistrationMethod = (*Max350Text)(&value)
}

func (p *ProxyAppointmentInformation1) AddDeadline() *DateFormat2Choice {
	p.Deadline = new(DateFormat2Choice)
	return p.Deadline
}

func (p *ProxyAppointmentInformation1) AddSTPDeadline() *DateFormat2Choice {
	p.STPDeadline = new(DateFormat2Choice)
	return p.STPDeadline
}

func (p *ProxyAppointmentInformation1) AddMarketDeadline() *DateFormat2Choice {
	p.MarketDeadline = new(DateFormat2Choice)
	return p.MarketDeadline
}

func (p *ProxyAppointmentInformation1) AddAuthorisedProxy() *Proxy1 {
	newValue := new(Proxy1)
	p.AuthorisedProxy = append(p.AuthorisedProxy, newValue)
	return newValue
}
