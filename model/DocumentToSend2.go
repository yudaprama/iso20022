package model

// Determine the type of document and the type of communication method to be used to notify a Party.
type DocumentToSend2 struct {

	// Type of document.
	Type *Max140Text `xml:"Tp"`

	// Party that should receive the document.
	Recipient *PartyIdentification2Choice `xml:"Rcpt"`

	// Communication method to be used.
	MethodOfTransmission *CommunicationMethod3Choice `xml:"MtdOfTrnsmssn"`
}

func (d *DocumentToSend2) SetType(value string) {
	d.Type = (*Max140Text)(&value)
}

func (d *DocumentToSend2) AddRecipient() *PartyIdentification2Choice {
	d.Recipient = new(PartyIdentification2Choice)
	return d.Recipient
}

func (d *DocumentToSend2) AddMethodOfTransmission() *CommunicationMethod3Choice {
	d.MethodOfTransmission = new(CommunicationMethod3Choice)
	return d.MethodOfTransmission
}
