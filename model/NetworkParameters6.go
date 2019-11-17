package model

// Parameters to communicate with a host.
type NetworkParameters6 struct {

	// Type of proxy.
	Type *NetworkType2Code `xml:"Tp"`

	// Access information to the proxy.
	Access *NetworkParameters5 `xml:"Accs"`
}

func (n *NetworkParameters6) SetType(value string) {
	n.Type = (*NetworkType2Code)(&value)
}

func (n *NetworkParameters6) AddAccess() *NetworkParameters5 {
	n.Access = new(NetworkParameters5)
	return n.Access
}
