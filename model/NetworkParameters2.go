package model

// Configuration parameters to communicate with a host.
type NetworkParameters2 struct {

	// IP address or hostname.
	Address *Max35Text `xml:"Adr"`

	// Port number of the server, if the default port number is not used.
	PortNumber *Number `xml:"PortNb,omitempty"`

	// Delay between two contacts of the server..
	Delay *ISOTime `xml:"Dely,omitempty"`
}

func (n *NetworkParameters2) SetAddress(value string) {
	n.Address = (*Max35Text)(&value)
}

func (n *NetworkParameters2) SetPortNumber(value string) {
	n.PortNumber = (*Number)(&value)
}

func (n *NetworkParameters2) SetDelay(value string) {
	n.Delay = (*ISOTime)(&value)
}
