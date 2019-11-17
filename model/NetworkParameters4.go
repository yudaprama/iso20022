package model

// Parameters to communicate with a host.
type NetworkParameters4 struct {

	// Type of communication network.
	NetworkType *NetworkType1Code `xml:"NtwkTp"`

	// Value of the address. The value of an internet protocol address contains the IP address or the DNS (Domain Name Server) address, followed by the character ':' and the port number if the default port is not used. The value of a public telephone address contains the phone number with possible prefix and extensions.
	AddressValue *Max70Text `xml:"AdrVal"`
}

func (n *NetworkParameters4) SetNetworkType(value string) {
	n.NetworkType = (*NetworkType1Code)(&value)
}

func (n *NetworkParameters4) SetAddressValue(value string) {
	n.AddressValue = (*Max70Text)(&value)
}
