package model

// Specifies the elements that identify a proxy appointed to represent a party authorised to vote at a shareholders meeting.
type Proxy3 struct {

	// Specifies the type of proxy.
	ProxyType []*ProxyType2Code `xml:"PrxyTp"`

	// Identifies an authorized proxy which has been assigned by the issuer of the meeting.
	PreassignedProxy *IndividualPerson16 `xml:"PrssgndPrxy,omitempty"`
}

func (p *Proxy3) AddProxyType(value string) {
	p.ProxyType = append(p.ProxyType, (*ProxyType2Code)(&value))
}

func (p *Proxy3) AddPreassignedProxy() *IndividualPerson16 {
	p.PreassignedProxy = new(IndividualPerson16)
	return p.PreassignedProxy
}
