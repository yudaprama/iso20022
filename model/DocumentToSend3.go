package model

// Type of document and the type of communication method to be used to notify a party.
type DocumentToSend3 struct {

	// Type of document.
	Type *Max140Text `xml:"Tp"`

	// Party that should receive the document.
	Recipient *PartyIdentification70Choice `xml:"Rcpt"`

	// Communication method to be used.
	MethodOfTransmission *CommunicationMethod3Choice `xml:"MtdOfTrnsmssn"`
}

func (d *DocumentToSend3) SetType(value string) {
	d.Type = (*Max140Text)(&value)
}

func (d *DocumentToSend3) AddRecipient() *PartyIdentification70Choice {
	d.Recipient = new(PartyIdentification70Choice)
	return d.Recipient
}

func (d *DocumentToSend3) AddMethodOfTransmission() *CommunicationMethod3Choice {
	d.MethodOfTransmission = new(CommunicationMethod3Choice)
	return d.MethodOfTransmission
}
