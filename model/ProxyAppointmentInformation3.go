package model

// Indicates how a proxy is registered.
type ProxyAppointmentInformation3 struct {

	// Specifies how to register the proxy.
	RegistrationMethod *Max350Text `xml:"RegnMtd,omitempty"`

	// Date by which the information on the proxy assignment must be received by the intermediary.
	Deadline *DateFormat29Choice `xml:"Ddln,omitempty"`

	// Date by which the information on the proxy assignment must be received by the intermediary (STP mode).
	STPDeadline *DateFormat29Choice `xml:"STPDdln,omitempty"`

	// Date by which the information on the proxy assignment must be received by the issuer.
	MarketDeadline *DateFormat29Choice `xml:"MktDdln,omitempty"`

	// Specifies the proxy person that is authorised by the issuer.
	AuthorisedProxy []*Proxy5 `xml:"AuthrsdPrxy,omitempty"`
}

func (p *ProxyAppointmentInformation3) SetRegistrationMethod(value string) {
	p.RegistrationMethod = (*Max350Text)(&value)
}

func (p *ProxyAppointmentInformation3) AddDeadline() *DateFormat29Choice {
	p.Deadline = new(DateFormat29Choice)
	return p.Deadline
}

func (p *ProxyAppointmentInformation3) AddSTPDeadline() *DateFormat29Choice {
	p.STPDeadline = new(DateFormat29Choice)
	return p.STPDeadline
}

func (p *ProxyAppointmentInformation3) AddMarketDeadline() *DateFormat29Choice {
	p.MarketDeadline = new(DateFormat29Choice)
	return p.MarketDeadline
}

func (p *ProxyAppointmentInformation3) AddAuthorisedProxy() *Proxy5 {
	newValue := new(Proxy5)
	p.AuthorisedProxy = append(p.AuthorisedProxy, newValue)
	return newValue
}
