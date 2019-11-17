package model

// Configuration parameters to communicate with a host.
type HostCommunicationParameter3 struct {

	// Identification of the host.
	HostIdentification *Max35Text `xml:"HstId"`

	// Network parameters of the host.
	Address *NetworkParameters3 `xml:"Adr,omitempty"`

	// Cryptographic key used to communicate with the host.
	Key []*KEKIdentifier2 `xml:"Key,omitempty"`
}

func (h *HostCommunicationParameter3) SetHostIdentification(value string) {
	h.HostIdentification = (*Max35Text)(&value)
}

func (h *HostCommunicationParameter3) AddAddress() *NetworkParameters3 {
	h.Address = new(NetworkParameters3)
	return h.Address
}

func (h *HostCommunicationParameter3) AddKey() *KEKIdentifier2 {
	newValue := new(KEKIdentifier2)
	h.Key = append(h.Key, newValue)
	return newValue
}
