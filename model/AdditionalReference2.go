package model

// References a related message or provides another reference, such as a pool reference, linking a set of messages. The party which issued the related reference may be the Sender of the referenced message or a party other than the Sender.
type AdditionalReference2 struct {

	// Business reference of a message assigned by the party issuing the message. This reference must be unique amongst all messages of the same name sent by the same party.
	Reference *Max35Text `xml:"Ref"`

	// Issuer of the reference.
	ReferenceIssuer *PartyIdentification1Choice `xml:"RefIssr,omitempty"`

	// Name of a message.
	MessageName *Max35Text `xml:"MsgNm,omitempty"`
}

func (a *AdditionalReference2) SetReference(value string) {
	a.Reference = (*Max35Text)(&value)
}

func (a *AdditionalReference2) AddReferenceIssuer() *PartyIdentification1Choice {
	a.ReferenceIssuer = new(PartyIdentification1Choice)
	return a.ReferenceIssuer
}

func (a *AdditionalReference2) SetMessageName(value string) {
	a.MessageName = (*Max35Text)(&value)
}
