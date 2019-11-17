package model

// Reference to a related message or transaction.
type AdditionalReference7 struct {

	// Reference issued by a party to identify an instruction, transaction or a message.
	Reference *Max35Text `xml:"Ref"`

	// Party that issued the reference.
	ReferenceIssuer *PartyIdentification97Choice `xml:"RefIssr,omitempty"`

	// Name of the message.
	MessageName *Max35Text `xml:"MsgNm,omitempty"`
}

func (a *AdditionalReference7) SetReference(value string) {
	a.Reference = (*Max35Text)(&value)
}

func (a *AdditionalReference7) AddReferenceIssuer() *PartyIdentification97Choice {
	a.ReferenceIssuer = new(PartyIdentification97Choice)
	return a.ReferenceIssuer
}

func (a *AdditionalReference7) SetMessageName(value string) {
	a.MessageName = (*Max35Text)(&value)
}
