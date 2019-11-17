package model

// Defines a signing party and a digital signature made by this party.
type QualifiedPartyAndXMLSignature1 struct {

	// Identification of the signing party.
	Party *QualifiedPartyIdentification1 `xml:"Pty,omitempty"`

	// Digital signature in XML-DSIG format and reference to signing party.
	Signature *SignatureEnvelope `xml:"Sgntr"`
}

func (q *QualifiedPartyAndXMLSignature1) AddParty() *QualifiedPartyIdentification1 {
	q.Party = new(QualifiedPartyIdentification1)
	return q.Party
}

func (q *QualifiedPartyAndXMLSignature1) AddSignature() *SignatureEnvelope {
	q.Signature = new(SignatureEnvelope)
	return q.Signature
}
