package model

// Identification of a cryptographic asymmetric key.
type Recipient5Choice struct {

	// Certificate issuer name and serial number (see ITU X.509).
	IssuerAndSerialNumber *IssuerAndSerialNumber1 `xml:"IssrAndSrlNb"`

	// Identifier of a cryptographic asymmetric key, previously exchanged between initiator and recipient.
	KeyIdentifier *KEKIdentifier2 `xml:"KeyIdr"`
}

func (r *Recipient5Choice) AddIssuerAndSerialNumber() *IssuerAndSerialNumber1 {
	r.IssuerAndSerialNumber = new(IssuerAndSerialNumber1)
	return r.IssuerAndSerialNumber
}

func (r *Recipient5Choice) AddKeyIdentifier() *KEKIdentifier2 {
	r.KeyIdentifier = new(KEKIdentifier2)
	return r.KeyIdentifier
}
