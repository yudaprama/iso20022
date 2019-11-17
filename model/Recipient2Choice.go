package model

// Transport key or key encryption key (KEK) for the recipient.
type Recipient2Choice struct {

	// Encryption key using previously distributed asymmetric public key.
	KeyTransport *KeyTransport2 `xml:"KeyTrnsprt,omitempty"`

	// Encryption key using previously distributed symmetric key.
	KEK *KEK2 `xml:"KEK,omitempty"`
}

func (r *Recipient2Choice) AddKeyTransport() *KeyTransport2 {
	r.KeyTransport = new(KeyTransport2)
	return r.KeyTransport
}

func (r *Recipient2Choice) AddKEK() *KEK2 {
	r.KEK = new(KEK2)
	return r.KEK
}
