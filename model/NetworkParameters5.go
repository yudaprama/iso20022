package model

// Parameters to communicate with a host.
type NetworkParameters5 struct {

	// Network addresses of the host.
	Address []*NetworkParameters4 `xml:"Adr"`

	// User name identifying the client.
	UserName *Max35Text `xml:"UsrNm,omitempty"`

	// Password authenticating the client.
	AccessCode *Max35Binary `xml:"AccsCd,omitempty"`

	// X.509 Certificate required to authenticate the server.
	ServerCertificate []*Max10KBinary `xml:"SvrCert,omitempty"`

	// Identification of the X.509 Certificates required to authenticate the server, for instance a digest of the certificate.
	ServerCertificateIdentifier []*Max140Binary `xml:"SvrCertIdr,omitempty"`

	// X.509 Certificate required to authenticate the client.
	ClientCertificate []*Max10KBinary `xml:"ClntCert,omitempty"`

	// Identification of the set of security elements to access the host.
	SecurityProfile *Max35Text `xml:"SctyPrfl,omitempty"`
}

func (n *NetworkParameters5) AddAddress() *NetworkParameters4 {
	newValue := new(NetworkParameters4)
	n.Address = append(n.Address, newValue)
	return newValue
}

func (n *NetworkParameters5) SetUserName(value string) {
	n.UserName = (*Max35Text)(&value)
}

func (n *NetworkParameters5) SetAccessCode(value string) {
	n.AccessCode = (*Max35Binary)(&value)
}

func (n *NetworkParameters5) AddServerCertificate(value string) {
	n.ServerCertificate = append(n.ServerCertificate, (*Max10KBinary)(&value))
}

func (n *NetworkParameters5) AddServerCertificateIdentifier(value string) {
	n.ServerCertificateIdentifier = append(n.ServerCertificateIdentifier, (*Max140Binary)(&value))
}

func (n *NetworkParameters5) AddClientCertificate(value string) {
	n.ClientCertificate = append(n.ClientCertificate, (*Max10KBinary)(&value))
}

func (n *NetworkParameters5) SetSecurityProfile(value string) {
	n.SecurityProfile = (*Max35Text)(&value)
}
