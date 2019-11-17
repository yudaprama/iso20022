package model

// References a related message or provides another reference, such as a pool reference, linking a set of messages. The party which issued the related reference may be the Sender of the referenced message or a party other than the Sender.
type AdditionalReference9 struct {

	// Reference identifying a set of messages.
	Reference *Max35Text `xml:"Ref"`

	// Issuer of the reference.
	ReferenceIssuer *PartyIdentification113 `xml:"RefIssr,omitempty"`

	// Name of the message.
	MessageName *Max35Text `xml:"MsgNm,omitempty"`
}

func (a *AdditionalReference9) SetReference(value string) {
	a.Reference = (*Max35Text)(&value)
}

func (a *AdditionalReference9) AddReferenceIssuer() *PartyIdentification113 {
	a.ReferenceIssuer = new(PartyIdentification113)
	return a.ReferenceIssuer
}

func (a *AdditionalReference9) SetMessageName(value string) {
	a.MessageName = (*Max35Text)(&value)
}
