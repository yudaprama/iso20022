package model

// Parameters to communicate with a host.
type NetworkParameters3 struct {

	// Network addresses of the host.
	Address []*NetworkParameters4 `xml:"Adr"`

	// User name identifying the client.
	UserName *Max35Text `xml:"UsrNm,omitempty"`

	// Password authenticating the client.
	AccessCode *Max35Binary `xml:"AccsCd,omitempty"`

	// X.509 Certificate required to authenticate the server.
	ServerCertificate []*Max3000Binary `xml:"SvrCert,omitempty"`

	// Identification of the X.509 Certificate required to authenticate the server, for instance a digest of the certificate.
	ServerCertificateIdentifier []*Max140Binary `xml:"SvrCertIdr,omitempty"`
}

func (n *NetworkParameters3) AddAddress() *NetworkParameters4 {
	newValue := new(NetworkParameters4)
	n.Address = append(n.Address, newValue)
	return newValue
}

func (n *NetworkParameters3) SetUserName(value string) {
	n.UserName = (*Max35Text)(&value)
}

func (n *NetworkParameters3) SetAccessCode(value string) {
	n.AccessCode = (*Max35Binary)(&value)
}

func (n *NetworkParameters3) AddServerCertificate(value string) {
	n.ServerCertificate = append(n.ServerCertificate, (*Max3000Binary)(&value))
}

func (n *NetworkParameters3) AddServerCertificateIdentifier(value string) {
	n.ServerCertificateIdentifier = append(n.ServerCertificateIdentifier, (*Max140Binary)(&value))
}
