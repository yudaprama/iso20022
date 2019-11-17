package model

// Acquirer configuration parameters for a host.
type AcquirerHostConfiguration1 struct {

	// Identification of a host.
	HostIdentification *Max35Text `xml:"HstId"`

	// Types of message to sent to this host.
	MessageToSend []*MessageFunction3Code `xml:"MsgToSnd"`
}

func (a *AcquirerHostConfiguration1) SetHostIdentification(value string) {
	a.HostIdentification = (*Max35Text)(&value)
}

func (a *AcquirerHostConfiguration1) AddMessageToSend(value string) {
	a.MessageToSend = append(a.MessageToSend, (*MessageFunction3Code)(&value))
}
