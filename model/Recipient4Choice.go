package model

// Transport key or key encryption key (KEK) for the recipient.
type Recipient4Choice struct {

	// Encryption key using previously distributed asymmetric public key.
	KeyTransport *KeyTransport4 `xml:"KeyTrnsprt"`

	// Key encryption key using previously distributed symmetric key.
	KEK *KEK4 `xml:"KEK"`

	// Identification of a protection key without a session key, shared and previously exchanged between the initiator and the recipient.
	KeyIdentifier *KEKIdentifier2 `xml:"KeyIdr"`
}

func (r *Recipient4Choice) AddKeyTransport() *KeyTransport4 {
	r.KeyTransport = new(KeyTransport4)
	return r.KeyTransport
}

func (r *Recipient4Choice) AddKEK() *KEK4 {
	r.KEK = new(KEK4)
	return r.KEK
}

func (r *Recipient4Choice) AddKeyIdentifier() *KEKIdentifier2 {
	r.KeyIdentifier = new(KEKIdentifier2)
	return r.KeyIdentifier
}
