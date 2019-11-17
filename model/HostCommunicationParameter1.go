package model

// Configuration parameters to communicate with a host.
type HostCommunicationParameter1 struct {

	// Identification of the host.
	HostIdentification *Max35Text `xml:"HstId"`

	// Network parameters of the host.
	Address *NetworkParameters1 `xml:"Adr,omitempty"`

	// Cryptographic key used to communicate with the host.
	Key []*CryptographicKey1 `xml:"Key,omitempty"`
}

func (h *HostCommunicationParameter1) SetHostIdentification(value string) {
	h.HostIdentification = (*Max35Text)(&value)
}

func (h *HostCommunicationParameter1) AddAddress() *NetworkParameters1 {
	h.Address = new(NetworkParameters1)
	return h.Address
}

func (h *HostCommunicationParameter1) AddKey() *CryptographicKey1 {
	newValue := new(CryptographicKey1)
	h.Key = append(h.Key, newValue)
	return newValue
}
