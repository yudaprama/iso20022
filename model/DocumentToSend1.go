package model

// Determine the type of document and the type of communication method to be used to notify a Party.
type DocumentToSend1 struct {

	// Type of document.
	Type *Max140Text `xml:"Tp"`

	// Party that should receive the document.
	Recipient *PartyIdentification2Choice `xml:"Rcpt"`

	// Communication method to be used.
	MethodOfTransmission *CommunicationMethod1Code `xml:"MtdOfTrnsmssn"`

	// Communication means used to send information.
	ExtendedMethodOfTransmission *Extended350Code `xml:"XtndedMtdOfTrnsmssn"`
}

func (d *DocumentToSend1) SetType(value string) {
	d.Type = (*Max140Text)(&value)
}

func (d *DocumentToSend1) AddRecipient() *PartyIdentification2Choice {
	d.Recipient = new(PartyIdentification2Choice)
	return d.Recipient
}

func (d *DocumentToSend1) SetMethodOfTransmission(value string) {
	d.MethodOfTransmission = (*CommunicationMethod1Code)(&value)
}

func (d *DocumentToSend1) SetExtendedMethodOfTransmission(value string) {
	d.ExtendedMethodOfTransmission = (*Extended350Code)(&value)
}
