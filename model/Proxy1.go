package model

// Specifies the elements that identify a proxy appointed to represent a party authorised to vote at a shareholders meeting.
type Proxy1 struct {

	// Specifies the type of proxy.
	ProxyType []*ProxyType2Code `xml:"PrxyTp"`

	// Identifies an authorized proxy which has been assigned by the issuer of the meeting.
	PreassignedProxy *IndividualPerson14 `xml:"PrssgndPrxy,omitempty"`
}

func (p *Proxy1) AddProxyType(value string) {
	p.ProxyType = append(p.ProxyType, (*ProxyType2Code)(&value))
}

func (p *Proxy1) AddPreassignedProxy() *IndividualPerson14 {
	p.PreassignedProxy = new(IndividualPerson14)
	return p.PreassignedProxy
}
