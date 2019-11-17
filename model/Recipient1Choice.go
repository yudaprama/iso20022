package model

// Transport key or key encryption key (KEK) for the recipient.
type Recipient1Choice struct {

	// Encryption key using previously distributed asymmetric key.
	KeyTransport *KeyTransport1 `xml:"KeyTrnsprt,omitempty"`

	// Encryption key using previously distributed symmetric key.
	KEK *KEK1 `xml:"KEK,omitempty"`
}

func (r *Recipient1Choice) AddKeyTransport() *KeyTransport1 {
	r.KeyTransport = new(KeyTransport1)
	return r.KeyTransport
}

func (r *Recipient1Choice) AddKEK() *KEK1 {
	r.KEK = new(KEK1)
	return r.KEK
}
