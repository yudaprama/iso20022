package model

// Specifies the elements that identify a proxy appointed to represent a party authorised to vote at a shareholders meeting.
type Proxy5 struct {

	// Specifies the type of proxy.
	ProxyType []*ProxyType2Code `xml:"PrxyTp"`

	// Identifies an authorized proxy that has been assigned by the issuer of the meeting.
	PreassignedProxy *IndividualPerson25 `xml:"PrssgndPrxy,omitempty"`
}

func (p *Proxy5) AddProxyType(value string) {
	p.ProxyType = append(p.ProxyType, (*ProxyType2Code)(&value))
}

func (p *Proxy5) AddPreassignedProxy() *IndividualPerson25 {
	p.PreassignedProxy = new(IndividualPerson25)
	return p.PreassignedProxy
}
