package model

// Reference to a related message or transaction.
type AdditionalReference6 struct {

	// Message identification of a message. This reference was assigned by the party issuing the message.
	Reference *Max35Text `xml:"Ref"`

	// Issuer of the reference.
	ReferenceIssuer *PartyIdentification90Choice `xml:"RefIssr,omitempty"`

	// Name of the message.
	MessageName *Max35Text `xml:"MsgNm,omitempty"`
}

func (a *AdditionalReference6) SetReference(value string) {
	a.Reference = (*Max35Text)(&value)
}

func (a *AdditionalReference6) AddReferenceIssuer() *PartyIdentification90Choice {
	a.ReferenceIssuer = new(PartyIdentification90Choice)
	return a.ReferenceIssuer
}

func (a *AdditionalReference6) SetMessageName(value string) {
	a.MessageName = (*Max35Text)(&value)
}
