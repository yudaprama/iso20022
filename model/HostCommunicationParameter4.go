package model

// Configuration parameters to communicate with a host.
type HostCommunicationParameter4 struct {

	// Type of action for the configuration parameters.
	ActionType *TerminalManagementAction3Code `xml:"ActnTp"`

	// Identification of the host.
	HostIdentification *Max35Text `xml:"HstId"`

	// Network parameters of the host.
	Address *NetworkParameters3 `xml:"Adr,omitempty"`

	// Cryptographic key used to communicate with the host.
	Key []*KEKIdentifier5 `xml:"Key,omitempty"`

	// Access information to reach an intermediate network service provider.
	NetworkServiceProvider *NetworkParameters5 `xml:"NtwkSvcPrvdr,omitempty"`
}

func (h *HostCommunicationParameter4) SetActionType(value string) {
	h.ActionType = (*TerminalManagementAction3Code)(&value)
}

func (h *HostCommunicationParameter4) SetHostIdentification(value string) {
	h.HostIdentification = (*Max35Text)(&value)
}

func (h *HostCommunicationParameter4) AddAddress() *NetworkParameters3 {
	h.Address = new(NetworkParameters3)
	return h.Address
}

func (h *HostCommunicationParameter4) AddKey() *KEKIdentifier5 {
	newValue := new(KEKIdentifier5)
	h.Key = append(h.Key, newValue)
	return newValue
}

func (h *HostCommunicationParameter4) AddNetworkServiceProvider() *NetworkParameters5 {
	h.NetworkServiceProvider = new(NetworkParameters5)
	return h.NetworkServiceProvider
}
