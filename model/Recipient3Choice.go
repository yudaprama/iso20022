package model

// Transport key or key encryption key (KEK) for the recipient.
type Recipient3Choice struct {

	// Encryption key using previously distributed asymmetric public key.
	KeyTransport *KeyTransport3 `xml:"KeyTrnsprt"`

	// Encryption key using previously distributed symmetric key.
	KEK *KEK3 `xml:"KEK"`

	// Identification of a cryptographic symetric key, previously exchanged between the initiator and the recipient.
	KeyIdentifier *KEKIdentifier1 `xml:"KeyIdr"`
}

func (r *Recipient3Choice) AddKeyTransport() *KeyTransport3 {
	r.KeyTransport = new(KeyTransport3)
	return r.KeyTransport
}

func (r *Recipient3Choice) AddKEK() *KEK3 {
	r.KEK = new(KEK3)
	return r.KEK
}

func (r *Recipient3Choice) AddKeyIdentifier() *KEKIdentifier1 {
	r.KeyIdentifier = new(KEKIdentifier1)
	return r.KeyIdentifier
}
