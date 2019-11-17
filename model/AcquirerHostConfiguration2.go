package model

// Acquirer configuration parameters for a host.
type AcquirerHostConfiguration2 struct {

	// Identification of a host.
	HostIdentification *Max35Text `xml:"HstId"`

	// Types of message to sent to this host.
	MessageToSend []*MessageFunction3Code `xml:"MsgToSnd,omitempty"`
}

func (a *AcquirerHostConfiguration2) SetHostIdentification(value string) {
	a.HostIdentification = (*Max35Text)(&value)
}

func (a *AcquirerHostConfiguration2) AddMessageToSend(value string) {
	a.MessageToSend = append(a.MessageToSend, (*MessageFunction3Code)(&value))
}
