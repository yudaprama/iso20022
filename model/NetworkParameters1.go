package model

// Configuration parameters to communicate with a host.
type NetworkParameters1 struct {

	// IP address or host name of the primary host.
	PrimaryAddress *Max35Text `xml:"PmryAdr"`

	// Port number of the primary host.
	PrimaryPortNumber *Number `xml:"PmryPortNb"`

	// IP address or host name of the secondary host.
	SecondaryAddress *Max35Text `xml:"ScndryAdr,omitempty"`

	// Port number of the secondary host.
	SecondaryPortNumber *Number `xml:"ScndryPortNb,omitempty"`

	// User name identifying the client.
	UserName *Max35Text `xml:"UsrNm,omitempty"`

	// Password authenticating the client.
	AccessCode *Max35Text `xml:"AccsCd,omitempty"`

	// Client certificate chain.
	ClientCertificate *Max3000Binary `xml:"ClntCert,omitempty"`
}

func (n *NetworkParameters1) SetPrimaryAddress(value string) {
	n.PrimaryAddress = (*Max35Text)(&value)
}

func (n *NetworkParameters1) SetPrimaryPortNumber(value string) {
	n.PrimaryPortNumber = (*Number)(&value)
}

func (n *NetworkParameters1) SetSecondaryAddress(value string) {
	n.SecondaryAddress = (*Max35Text)(&value)
}

func (n *NetworkParameters1) SetSecondaryPortNumber(value string) {
	n.SecondaryPortNumber = (*Number)(&value)
}

func (n *NetworkParameters1) SetUserName(value string) {
	n.UserName = (*Max35Text)(&value)
}

func (n *NetworkParameters1) SetAccessCode(value string) {
	n.AccessCode = (*Max35Text)(&value)
}

func (n *NetworkParameters1) SetClientCertificate(value string) {
	n.ClientCertificate = (*Max3000Binary)(&value)
}
