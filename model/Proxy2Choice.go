package model

// Choice of proxy allowance.
type Proxy2Choice struct {

	// Specifies the elements required to assign a proxy.
	Proxy *ProxyAppointmentInformation3 `xml:"Prxy"`

	// Indicates that no proxy is allowed for a meeting.
	ProxyNotAllowed *ProxyNotAllowedCode `xml:"PrxyNotAllwd"`
}

func (p *Proxy2Choice) AddProxy() *ProxyAppointmentInformation3 {
	p.Proxy = new(ProxyAppointmentInformation3)
	return p.Proxy
}

func (p *Proxy2Choice) SetProxyNotAllowed(value string) {
	p.ProxyNotAllowed = (*ProxyNotAllowedCode)(&value)
}
