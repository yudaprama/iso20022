package model

// Choice of proxy allowance.
type Proxy1Choice struct {

	// Specifies the elements required to assign a proxy.
	Proxy *ProxyAppointmentInformation2 `xml:"Prxy"`

	// Indicates that no proxy is allowed for a meeting.
	ProxyNotAllowed *ProxyNotAllowedCode `xml:"PrxyNotAllwd"`
}

func (p *Proxy1Choice) AddProxy() *ProxyAppointmentInformation2 {
	p.Proxy = new(ProxyAppointmentInformation2)
	return p.Proxy
}

func (p *Proxy1Choice) SetProxyNotAllowed(value string) {
	p.ProxyNotAllowed = (*ProxyNotAllowedCode)(&value)
}
