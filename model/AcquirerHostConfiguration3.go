package model

// Acquirer configuration parameters for a host.
type AcquirerHostConfiguration3 struct {

	// Identification of a host.
	HostIdentification *Max35Text `xml:"HstId"`

	// Types of message to sent to this host.
	MessageToSend []*MessageFunction5Code `xml:"MsgToSnd,omitempty"`
}

func (a *AcquirerHostConfiguration3) SetHostIdentification(value string) {
	a.HostIdentification = (*Max35Text)(&value)
}

func (a *AcquirerHostConfiguration3) AddMessageToSend(value string) {
	a.MessageToSend = append(a.MessageToSend, (*MessageFunction5Code)(&value))
}
