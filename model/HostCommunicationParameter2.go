package model

// Configuration parameters to communicate with a host.
type HostCommunicationParameter2 struct {

	// Identification of the host.
	HostIdentification *Max35Text `xml:"HstId"`

	// Network parameters of the host.
	Address *NetworkParameters1 `xml:"Adr,omitempty"`

	// Cryptographic key used to communicate with the host.
	Key []*CryptographicKey3 `xml:"Key,omitempty"`
}

func (h *HostCommunicationParameter2) SetHostIdentification(value string) {
	h.HostIdentification = (*Max35Text)(&value)
}

func (h *HostCommunicationParameter2) AddAddress() *NetworkParameters1 {
	h.Address = new(NetworkParameters1)
	return h.Address
}

func (h *HostCommunicationParameter2) AddKey() *CryptographicKey3 {
	newValue := new(CryptographicKey3)
	h.Key = append(h.Key, newValue)
	return newValue
}
